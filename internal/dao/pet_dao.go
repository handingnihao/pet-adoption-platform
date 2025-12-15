package dao

import (
"context"
"strings"
"time"
"gorm.io/gorm"
"pet-adoption-platform/internal/model"
)

type PetDAO struct {
db *gorm.DB
}

func NewPetDAO(db *gorm.DB) *PetDAO {
return &PetDAO{db: db}
}

func (d *PetDAO) Create(ctx context.Context, pet *model.Pet) error {
return d.db.WithContext(ctx).Create(pet).Error
}

func (d *PetDAO) GetByID(ctx context.Context, id uint) (*model.Pet, error) {
var pet model.Pet
err := d.db.WithContext(ctx).Preload("User").First(&pet, id).Error
if err != nil {
return nil, err
}
return &pet, nil
}

func (d *PetDAO) Update(ctx context.Context, pet *model.Pet) error {
return d.db.WithContext(ctx).Save(pet).Error
}

func (d *PetDAO) UpdateByID(ctx context.Context, id uint64, updates map[string]interface{}) error {
return d.db.WithContext(ctx).Model(&model.Pet{}).Where("id = ?", id).Updates(updates).Error
}

func (d *PetDAO) Delete(ctx context.Context, id uint) error {
return d.db.WithContext(ctx).Delete(&model.Pet{}, id).Error
}

func (d *PetDAO) List(ctx context.Context, page, pageSize int) ([]*model.Pet, int64, error) {
var pets []*model.Pet
var total int64
offset := (page - 1) * pageSize
// 只查询已发布状态的宠物
query := d.db.WithContext(ctx).Model(&model.Pet{}).Where("status = ?", model.PetStatusAvailable)
if err := query.Count(&total).Error; err != nil {
return nil, 0, err
}
err := query.Preload("User").Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&pets).Error
return pets, total, err
}

func (d *PetDAO) QueryByConditions(ctx context.Context, req *model.PetQueryRequest) ([]*model.Pet, int64, error) {
var pets []*model.Pet
var total int64
query := d.db.WithContext(ctx).Model(&model.Pet{})
query = d.buildQuery(query, req)
if err := query.Count(&total).Error; err != nil {
return nil, 0, err
}
offset := (req.Page - 1) * req.PageSize
err := query.Preload("User").Offset(offset).Limit(req.PageSize).Order("created_at DESC").Find(&pets).Error
return pets, total, err
}

func (d *PetDAO) buildQuery(query *gorm.DB, req *model.PetQueryRequest) *gorm.DB {
if req.Type != "" {
query = query.Where("type = ?", req.Type)
}
if req.Gender != "" {
query = query.Where("gender = ?", req.Gender)
}
if req.Size != "" {
query = query.Where("size = ?", req.Size)
}
// 如果没有指定状态，默认只查询已发布的宠物
if req.Status != nil {
query = query.Where("status = ?", *req.Status)
} else {
query = query.Where("status = ?", model.PetStatusAvailable)
}
if req.Province != "" {
query = query.Where("province = ?", req.Province)
}
if req.City != "" {
query = query.Where("city = ?", req.City)
}
if req.MinAge != nil {
query = query.Where("age >= ?", *req.MinAge)
}
if req.MaxAge != nil {
query = query.Where("age <= ?", *req.MaxAge)
}
if req.Keyword != "" {
keyword := "%" + req.Keyword + "%"
query = query.Where("name LIKE ? OR breed LIKE ? OR description LIKE ?", keyword, keyword, keyword)
}
return query
}

func (d *PetDAO) SearchByKeyword(ctx context.Context, keyword string, page, pageSize int) ([]*model.Pet, int64, error) {
var pets []*model.Pet
var total int64
// 只搜索已发布状态的宠物
query := d.db.WithContext(ctx).Model(&model.Pet{}).Where("status = ?", model.PetStatusAvailable)
if keyword != "" {
query = query.Where("MATCH(name, breed, description) AGAINST(? IN NATURAL LANGUAGE MODE)", keyword)
}
if err := query.Count(&total).Error; err != nil {
return nil, 0, err
}
offset := (page - 1) * pageSize
err := query.Preload("User").Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&pets).Error
return pets, total, err
}

func (d *PetDAO) GetByUserID(ctx context.Context, userID int64, page, pageSize int) ([]*model.Pet, int64, error) {
var pets []*model.Pet
var total int64
offset := (page - 1) * pageSize
query := d.db.WithContext(ctx).Model(&model.Pet{}).Where("user_id = ?", userID)
if err := query.Count(&total).Error; err != nil {
return nil, 0, err
}
err := query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&pets).Error
return pets, total, err
}

func (d *PetDAO) GetByStatus(ctx context.Context, status model.PetStatus, page, pageSize int) ([]*model.Pet, int64, error) {
var pets []*model.Pet
var total int64
offset := (page - 1) * pageSize
query := d.db.WithContext(ctx).Model(&model.Pet{}).Where("status = ?", status)
if err := query.Count(&total).Error; err != nil {
return nil, 0, err
}
err := query.Preload("User").Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&pets).Error
return pets, total, err
}

func (d *PetDAO) UpdateStatus(ctx context.Context, id uint, status model.PetStatus) error {
return d.db.WithContext(ctx).Model(&model.Pet{}).Where("id = ?", id).Update("status", status).Error
}

func (d *PetDAO) IncrementViewCount(ctx context.Context, id uint) error {
return d.db.WithContext(ctx).Model(&model.Pet{}).Where("id = ?", id).UpdateColumn("view_count", gorm.Expr("view_count + 1")).Error
}

func (d *PetDAO) IncrementFavoriteCount(ctx context.Context, id uint) error {
return d.db.WithContext(ctx).Model(&model.Pet{}).Where("id = ?", id).UpdateColumn("favorite_count", gorm.Expr("favorite_count + 1")).Error
}

func (d *PetDAO) DecrementFavoriteCount(ctx context.Context, id uint) error {
return d.db.WithContext(ctx).Model(&model.Pet{}).Where("id = ?", id).Where("favorite_count > 0").UpdateColumn("favorite_count", gorm.Expr("favorite_count - 1")).Error
}

func (d *PetDAO) SetAdopted(ctx context.Context, petID, userID int64) error {
now := time.Now()
return d.db.WithContext(ctx).Model(&model.Pet{}).Where("id = ?", petID).Updates(map[string]interface{}{"status": model.PetStatusAdopted, "adopted_by": userID, "adopted_at": now}).Error
}

func (d *PetDAO) SetReviewed(ctx context.Context, petID, reviewerID int64, status model.PetStatus, reason string) error {
now := time.Now()
updates := map[string]interface{}{"status": status, "reviewed_by": reviewerID, "reviewed_at": now, "reject_reason": reason}
return d.db.WithContext(ctx).Model(&model.Pet{}).Where("id = ?", petID).Updates(updates).Error
}

// PetStatistics 宠物统计结果
type PetStatistics struct {
	Total     int64            `json:"total"`
	Available int64            `json:"available"`
	Pending   int64            `json:"pending"`
	Adopted   int64            `json:"adopted"`
	Offline   int64            `json:"offline"`
	ByType    map[string]int64 `json:"by_type"`
}

func (d *PetDAO) GetStatistics(ctx context.Context) (*PetStatistics, error) {
	stats := &PetStatistics{
		ByType: make(map[string]int64),
	}

	// 总数
	if err := d.db.WithContext(ctx).Model(&model.Pet{}).Count(&stats.Total).Error; err != nil {
		return nil, err
	}

	// 按状态统计
	var statusCounts []struct {
		Status model.PetStatus
		Count  int64
	}
	if err := d.db.WithContext(ctx).Model(&model.Pet{}).
		Select("status, COUNT(*) as count").
		Group("status").
		Scan(&statusCounts).Error; err != nil {
		return nil, err
	}
	for _, sc := range statusCounts {
		switch sc.Status {
		case model.PetStatusPending:
			stats.Pending = sc.Count
		case model.PetStatusAvailable:
			stats.Available = sc.Count
		case model.PetStatusAdopted:
			stats.Adopted = sc.Count
		case model.PetStatusOffline:
			stats.Offline = sc.Count
		}
	}

	// 按类型统计（只统计可领养状态的宠物）
	var typeCounts []struct {
		Type  string
		Count int64
	}
	if err := d.db.WithContext(ctx).Model(&model.Pet{}).
		Select("type, COUNT(*) as count").
		Where("status = ?", model.PetStatusAvailable).
		Group("type").
		Scan(&typeCounts).Error; err != nil {
		return nil, err
	}
	for _, tc := range typeCounts {
		stats.ByType[tc.Type] = tc.Count
	}

	return stats, nil
}

func (d *PetDAO) GetRecommended(ctx context.Context, limit int) ([]*model.Pet, error) {
var pets []*model.Pet
err := d.db.WithContext(ctx).Preload("User").Where("status = ?", model.PetStatusAvailable).Order("view_count DESC, favorite_count DESC, created_at DESC").Limit(limit).Find(&pets).Error
return pets, err
}

func ConvertPhotosToSlice(photos string) []string {
if photos == "" {
return []string{}
}
return strings.Split(photos, ",")
}

func ConvertPhotosToString(photos []string) string {
if len(photos) == 0 {
return ""
}
return strings.Join(photos, ",")
}
