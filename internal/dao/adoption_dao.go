package dao

import (
	"context"
	"pet-adoption-platform/internal/model"
	"time"

	"gorm.io/gorm"
)

type AdoptionDAO struct {
	db *gorm.DB
}

func NewAdoptionDAO(db *gorm.DB) *AdoptionDAO {
	return &AdoptionDAO{db: db}
}

// ========== AdoptionApplication CRUD ==========

// CreateApplication 创建领养申请
func (d *AdoptionDAO) CreateApplication(ctx context.Context, app *model.AdoptionApplication) error {
	return d.db.WithContext(ctx).Create(app).Error
}

// GetApplicationByID 根据ID获取申请
func (d *AdoptionDAO) GetApplicationByID(ctx context.Context, id uint64) (*model.AdoptionApplication, error) {
	var app model.AdoptionApplication
	err := d.db.WithContext(ctx).
		Preload("User").
		Preload("Pet").
		Preload("Organization").
		Preload("Reviewer").
		First(&app, id).Error
	if err != nil {
		return nil, err
	}
	return &app, nil
}

// GetApplicationByNo 根据申请编号获取
func (d *AdoptionDAO) GetApplicationByNo(ctx context.Context, applicationNo string) (*model.AdoptionApplication, error) {
	var app model.AdoptionApplication
	err := d.db.WithContext(ctx).
		Preload("User").
		Preload("Pet").
		Preload("Organization").
		Preload("Reviewer").
		Where("application_no = ?", applicationNo).
		First(&app).Error
	if err != nil {
		return nil, err
	}
	return &app, nil
}

// UpdateApplication 更新申请
func (d *AdoptionDAO) UpdateApplication(ctx context.Context, id uint64, updates map[string]interface{}) error {
	return d.db.WithContext(ctx).
		Model(&model.AdoptionApplication{}).
		Where("id = ?", id).
		Updates(updates).Error
}

// DeleteApplication 删除申请（软删除）
func (d *AdoptionDAO) DeleteApplication(ctx context.Context, id uint64) error {
	return d.db.WithContext(ctx).Delete(&model.AdoptionApplication{}, id).Error
}

// ListApplications 分页列表
func (d *AdoptionDAO) ListApplications(ctx context.Context, page, pageSize int) ([]*model.AdoptionApplication, int64, error) {
	var apps []*model.AdoptionApplication
	var total int64

	db := d.db.WithContext(ctx).Model(&model.AdoptionApplication{})

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := db.Preload("User").
		Preload("Pet").
		Preload("Organization").
		Preload("Reviewer").
		Order("created_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&apps).Error

	return apps, total, err
}

// GetApplicationsByUserID 获取用户的所有申请
func (d *AdoptionDAO) GetApplicationsByUserID(ctx context.Context, userID int64, page, pageSize int) ([]*model.AdoptionApplication, int64, error) {
	var apps []*model.AdoptionApplication
	var total int64

	db := d.db.WithContext(ctx).Model(&model.AdoptionApplication{}).Where("user_id = ?", userID)

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := db.Preload("Pet").
		Preload("Organization").
		Preload("Reviewer").
		Order("created_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&apps).Error

	return apps, total, err
}

// GetApplicationsByPetID 获取宠物的所有申请
func (d *AdoptionDAO) GetApplicationsByPetID(ctx context.Context, petID uint64, page, pageSize int) ([]*model.AdoptionApplication, int64, error) {
	var apps []*model.AdoptionApplication
	var total int64

	db := d.db.WithContext(ctx).Model(&model.AdoptionApplication{}).Where("pet_id = ?", petID)

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := db.Preload("User").
		Preload("Organization").
		Preload("Reviewer").
		Order("created_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&apps).Error

	return apps, total, err
}

// GetApplicationsByStatus 根据状态查询
func (d *AdoptionDAO) GetApplicationsByStatus(ctx context.Context, status model.ApplicationStatus, page, pageSize int) ([]*model.AdoptionApplication, int64, error) {
	var apps []*model.AdoptionApplication
	var total int64

	db := d.db.WithContext(ctx).Model(&model.AdoptionApplication{}).Where("status = ?", status)

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := db.Preload("User").
		Preload("Pet").
		Preload("Organization").
		Preload("Reviewer").
		Order("created_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&apps).Error

	return apps, total, err
}

// QueryApplications 条件查询
func (d *AdoptionDAO) QueryApplications(ctx context.Context, query *model.ApplicationQueryRequest) ([]*model.AdoptionApplication, int64, error) {
	var apps []*model.AdoptionApplication
	var total int64

	db := d.db.WithContext(ctx).Model(&model.AdoptionApplication{})

	// 状态过滤
	if query.Status >= 0 {
		db = db.Where("status = ?", query.Status)
	}

	// 宠物ID过滤
	if query.PetID > 0 {
		db = db.Where("pet_id = ?", query.PetID)
	}

	// 机构ID过滤
	if query.OrganizationID > 0 {
		db = db.Where("organization_id = ?", query.OrganizationID)
	}

	// 日期范围过滤
	if query.StartDate != "" {
		db = db.Where("created_at >= ?", query.StartDate)
	}
	if query.EndDate != "" {
		db = db.Where("created_at <= ?", query.EndDate+" 23:59:59")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (query.Page - 1) * query.PageSize
	err := db.Preload("User").
		Preload("Pet").
		Preload("Organization").
		Preload("Reviewer").
		Order("created_at DESC").
		Limit(query.PageSize).
		Offset(offset).
		Find(&apps).Error

	return apps, total, err
}

// CheckDuplicateApplication 检查重复申请
func (d *AdoptionDAO) CheckDuplicateApplication(ctx context.Context, userID int64, petID uint64) (bool, error) {
	var count int64
	// 检查是否有未完成的申请（pending, reviewing, interview, home_visit）
	err := d.db.WithContext(ctx).Model(&model.AdoptionApplication{}).
		Where("user_id = ? AND pet_id = ?", userID, petID).
		Where("status IN ?", []model.ApplicationStatus{
			model.ApplicationStatusPending,
			model.ApplicationStatusReviewing,
			model.ApplicationStatusInterview,
			model.ApplicationStatusHomeVisit,
		}).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// UpdateApplicationStatus 更新申请状态
func (d *AdoptionDAO) UpdateApplicationStatus(ctx context.Context, id uint64, status model.ApplicationStatus, reviewerID int64, comment string) error {
	updates := map[string]interface{}{
		"status":      status,
		"reviewer_id": reviewerID,
	}
	if comment != "" {
		updates["review_comment"] = comment
	}
	
	// 根据状态更新相应的时间字段
	now := time.Now()
	switch status {
	case model.ApplicationStatusApproved:
		updates["approved_at"] = now
	case model.ApplicationStatusRejected:
		updates["rejected_at"] = now
	case model.ApplicationStatusInterview:
		updates["interview_time"] = now.Add(24 * time.Hour) // 默认24小时后面试
	case model.ApplicationStatusHomeVisit:
		updates["home_visit_time"] = now.Add(48 * time.Hour) // 默认48小时后家访
	}

	return d.db.WithContext(ctx).
		Model(&model.AdoptionApplication{}).
		Where("id = ?", id).
		Updates(updates).Error
}

// ========== Adoption CRUD ==========

// CreateAdoption 创建领养记录
func (d *AdoptionDAO) CreateAdoption(ctx context.Context, adoption *model.Adoption) error {
	return d.db.WithContext(ctx).Create(adoption).Error
}

// GetAdoptionByID 根据ID获取领养记录
func (d *AdoptionDAO) GetAdoptionByID(ctx context.Context, id uint64) (*model.Adoption, error) {
	var adoption model.Adoption
	err := d.db.WithContext(ctx).
		Preload("Application").
		Preload("User").
		Preload("Pet").
		Preload("Organization").
		First(&adoption, id).Error
	if err != nil {
		return nil, err
	}
	return &adoption, nil
}

// GetAdoptionByApplicationID 根据申请ID获取领养记录
func (d *AdoptionDAO) GetAdoptionByApplicationID(ctx context.Context, applicationID uint64) (*model.Adoption, error) {
	var adoption model.Adoption
	err := d.db.WithContext(ctx).
		Preload("Application").
		Preload("User").
		Preload("Pet").
		Preload("Organization").
		Where("application_id = ?", applicationID).
		First(&adoption).Error
	if err != nil {
		return nil, err
	}
	return &adoption, nil
}

// UpdateAdoption 更新领养记录
func (d *AdoptionDAO) UpdateAdoption(ctx context.Context, id uint64, updates map[string]interface{}) error {
	return d.db.WithContext(ctx).
		Model(&model.Adoption{}).
		Where("id = ?", id).
		Updates(updates).Error
}

// ListAdoptions 分页列表
func (d *AdoptionDAO) ListAdoptions(ctx context.Context, page, pageSize int) ([]*model.Adoption, int64, error) {
	var adoptions []*model.Adoption
	var total int64

	db := d.db.WithContext(ctx).Model(&model.Adoption{})

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := db.Preload("Application").
		Preload("User").
		Preload("Pet").
		Preload("Organization").
		Order("adoption_date DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&adoptions).Error

	return adoptions, total, err
}

// GetAdoptionsByUserID 获取用户的领养记录
func (d *AdoptionDAO) GetAdoptionsByUserID(ctx context.Context, userID int64, page, pageSize int) ([]*model.Adoption, int64, error) {
	var adoptions []*model.Adoption
	var total int64

	db := d.db.WithContext(ctx).Model(&model.Adoption{}).Where("user_id = ?", userID)

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := db.Preload("Application").
		Preload("Pet").
		Preload("Organization").
		Order("adoption_date DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&adoptions).Error

	return adoptions, total, err
}

// GetAdoptionsByStatus 根据状态查询
func (d *AdoptionDAO) GetAdoptionsByStatus(ctx context.Context, status model.AdoptionStatus, page, pageSize int) ([]*model.Adoption, int64, error) {
	var adoptions []*model.Adoption
	var total int64

	db := d.db.WithContext(ctx).Model(&model.Adoption{}).Where("status = ?", status)

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := db.Preload("Application").
		Preload("User").
		Preload("Pet").
		Preload("Organization").
		Order("adoption_date DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&adoptions).Error

	return adoptions, total, err
}

// GetStatistics 获取统计信息
func (d *AdoptionDAO) GetStatistics(ctx context.Context) (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// 申请统计
	var totalApplications int64
	var pendingApplications int64
	var approvedApplications int64
	var rejectedApplications int64

	d.db.WithContext(ctx).Model(&model.AdoptionApplication{}).Count(&totalApplications)
	d.db.WithContext(ctx).Model(&model.AdoptionApplication{}).Where("status = ?", model.ApplicationStatusPending).Count(&pendingApplications)
	d.db.WithContext(ctx).Model(&model.AdoptionApplication{}).Where("status = ?", model.ApplicationStatusApproved).Count(&approvedApplications)
	d.db.WithContext(ctx).Model(&model.AdoptionApplication{}).Where("status = ?", model.ApplicationStatusRejected).Count(&rejectedApplications)

	stats["total_applications"] = totalApplications
	stats["pending_applications"] = pendingApplications
	stats["approved_applications"] = approvedApplications
	stats["rejected_applications"] = rejectedApplications

	// 领养记录统计
	var totalAdoptions int64
	var activeAdoptions int64
	var returnedAdoptions int64

	d.db.WithContext(ctx).Model(&model.Adoption{}).Count(&totalAdoptions)
	d.db.WithContext(ctx).Model(&model.Adoption{}).Where("status = ?", model.AdoptionStatusActive).Count(&activeAdoptions)
	d.db.WithContext(ctx).Model(&model.Adoption{}).Where("status = ?", model.AdoptionStatusReturned).Count(&returnedAdoptions)

	stats["total_adoptions"] = totalAdoptions
	stats["active_adoptions"] = activeAdoptions
	stats["returned_adoptions"] = returnedAdoptions

	return stats, nil
}
