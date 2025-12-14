package dao

import (
	"context"
	"errors"
	"pet-adoption-platform/internal/model"
	"pet-adoption-platform/pkg/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// OrganizationDAO 机构数据访问对象
type OrganizationDAO struct {
	db *gorm.DB
}

// NewOrganizationDAO 创建机构DAO实例
func NewOrganizationDAO(db *gorm.DB) *OrganizationDAO {
	return &OrganizationDAO{db: db}
}

// Create 创建机构
func (dao *OrganizationDAO) Create(ctx context.Context, org *model.Organization) error {
	if org == nil {
		return errors.New("organization is nil")
	}

	if err := dao.db.WithContext(ctx).Create(org).Error; err != nil {
		logger.Error("创建机构失败", zap.Error(err))
		return err
	}

	return nil
}

// Update 更新机构信息
func (dao *OrganizationDAO) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	if id <= 0 {
		return errors.New("invalid organization id")
	}

	updates["updated_at"] = gorm.Expr("CURRENT_TIMESTAMP")

	result := dao.db.WithContext(ctx).Model(&model.Organization{}).
		Where("id = ?", id).
		Updates(updates)

	if result.Error != nil {
		logger.Error("更新机构信息失败", zap.Error(result.Error))
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// GetByID 根据ID获取机构
func (dao *OrganizationDAO) GetByID(ctx context.Context, id int64) (*model.Organization, error) {
	if id <= 0 {
		return nil, errors.New("invalid organization id")
	}

	var org model.Organization
	if err := dao.db.WithContext(ctx).First(&org, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.Error("查询机构失败", zap.Error(err), zap.Int64("id", id))
		return nil, err
	}

	return &org, nil
}

// List 获取机构列表
func (dao *OrganizationDAO) List(ctx context.Context, page, pageSize int, status *model.OrganizationStatus) ([]*model.Organization, int64, error) {
	var (
		orgs  []*model.Organization
		total int64
	)

	db := dao.db.WithContext(ctx).Model(&model.Organization{})

	// 状态筛选
	if status != nil {
		db = db.Where("status = ?", *status)
	}

	// 计算总数
	if err := db.Count(&total).Error; err != nil {
		logger.Error("获取机构总数失败", zap.Error(err))
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := db.Offset(offset).Limit(pageSize).
		Order("created_at DESC").
		Find(&orgs).Error; err != nil {
		logger.Error("获取机构列表失败", zap.Error(err))
		return nil, 0, err
	}

	return orgs, total, nil
}

// Delete 删除机构（软删除）
func (dao *OrganizationDAO) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid organization id")
	}

	result := dao.db.WithContext(ctx).Delete(&model.Organization{}, id)
	if result.Error != nil {
		logger.Error("删除机构失败", zap.Error(result.Error), zap.Int64("id", id))
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// UpdateStatus 更新机构状态
func (dao *OrganizationDAO) UpdateStatus(ctx context.Context, id int64, status model.OrganizationStatus, rejectReason string) error {
	if id <= 0 {
		return errors.New("invalid organization id")
	}

	updates := map[string]interface{}{
		"status":        status,
		"reject_reason": rejectReason,
		"updated_at":    gorm.Expr("CURRENT_TIMESTAMP"),
	}

	result := dao.db.WithContext(ctx).Model(&model.Organization{}).
		Where("id = ?", id).
		Updates(updates)

	if result.Error != nil {
		logger.Error("更新机构状态失败", zap.Error(result.Error), zap.Int64("id", id))
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// CheckNameExists 检查机构名称是否已存在
func (dao *OrganizationDAO) CheckNameExists(ctx context.Context, name string, excludeID ...int64) (bool, error) {
	if name == "" {
		return false, errors.New("organization name is empty")
	}

	db := dao.db.WithContext(ctx).Model(&model.Organization{}).
		Where("name = ?", name)

	if len(excludeID) > 0 && excludeID[0] > 0 {
		db = db.Where("id != ?", excludeID[0])
	}

	var count int64
	if err := db.Count(&count).Error; err != nil {
		logger.Error("检查机构名称是否存在失败", zap.Error(err), zap.String("name", name))
		return false, err
	}

	return count > 0, nil
}
