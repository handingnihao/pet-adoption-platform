package dao

import (
	"context"
	"errors"
	"pet-adoption-platform/internal/model"
	"time"

	"gorm.io/gorm"
)

// DonationDAO 捐赠数据访问对象
type DonationDAO struct {
	db *gorm.DB
}

// NewDonationDAO 创建DonationDAO实例
func NewDonationDAO(db *gorm.DB) *DonationDAO {
	return &DonationDAO{db: db}
}

// Create 创建捐赠
func (d *DonationDAO) Create(ctx context.Context, donation *model.Donation) error {
	return d.db.WithContext(ctx).Create(donation).Error
}

// GetByID 根据ID获取捐赠
func (d *DonationDAO) GetByID(ctx context.Context, id uint64) (*model.Donation, error) {
	var donation model.Donation
	err := d.db.WithContext(ctx).
		Preload("User").
		Preload("Organization").
		First(&donation, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &donation, nil
}

// List 获取捐赠列表
func (d *DonationDAO) List(ctx context.Context, page, pageSize int, donationType *model.DonationType, status *model.DonationStatus) ([]*model.Donation, int64, error) {
	var donations []*model.Donation
	var total int64

	query := d.db.WithContext(ctx).Model(&model.Donation{})

	if donationType != nil {
		query = query.Where("type = ?", *donationType)
	}
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := query.
		Preload("User").
		Preload("Organization").
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&donations).Error

	return donations, total, err
}

// ListByUserID 获取用户的捐赠记录
func (d *DonationDAO) ListByUserID(ctx context.Context, userID uint64, page, pageSize int) ([]*model.Donation, int64, error) {
	var donations []*model.Donation
	var total int64

	query := d.db.WithContext(ctx).Model(&model.Donation{}).Where("user_id = ?", userID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := query.
		Preload("Organization").
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&donations).Error

	return donations, total, err
}

// ListByOrganizationID 获取机构收到的捐赠
func (d *DonationDAO) ListByOrganizationID(ctx context.Context, orgID uint64, page, pageSize int) ([]*model.Donation, int64, error) {
	var donations []*model.Donation
	var total int64

	query := d.db.WithContext(ctx).Model(&model.Donation{}).Where("organization_id = ?", orgID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := query.
		Preload("User").
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&donations).Error

	return donations, total, err
}

// UpdateStatus 更新捐赠状态
func (d *DonationDAO) UpdateStatus(ctx context.Context, id uint64, status model.DonationStatus, confirmedBy *uint64) error {
	updates := map[string]interface{}{
		"status": status,
	}
	if status == model.DonationStatusConfirmed || status == model.DonationStatusCompleted {
		now := time.Now()
		updates["confirmed_at"] = now
		updates["confirmed_by"] = confirmedBy
	}
	return d.db.WithContext(ctx).Model(&model.Donation{}).Where("id = ?", id).Updates(updates).Error
}

// UpdateTransactionID 更新交易ID
func (d *DonationDAO) UpdateTransactionID(ctx context.Context, id uint64, transactionID string) error {
	return d.db.WithContext(ctx).Model(&model.Donation{}).
		Where("id = ?", id).
		Update("transaction_id", transactionID).Error
}

// GetStatistics 获取捐赠统计
func (d *DonationDAO) GetStatistics(ctx context.Context) (*model.DonationStatistics, error) {
	stats := &model.DonationStatistics{}

	// 总捐赠次数（已完成的）
	d.db.WithContext(ctx).Model(&model.Donation{}).
		Where("status = ?", model.DonationStatusCompleted).
		Count(&stats.TotalDonations)

	// 总捐赠金额
	d.db.WithContext(ctx).Model(&model.Donation{}).
		Where("status = ? AND type = ?", model.DonationStatusCompleted, model.DonationTypeMoney).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&stats.TotalAmount)

	// 捐赠人数
	d.db.WithContext(ctx).Model(&model.Donation{}).
		Where("status = ?", model.DonationStatusCompleted).
		Distinct("user_id").
		Count(&stats.TotalDonors)

	// 本月捐赠
	now := time.Now()
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	d.db.WithContext(ctx).Model(&model.Donation{}).
		Where("status = ? AND created_at >= ?", model.DonationStatusCompleted, monthStart).
		Count(&stats.MonthlyDonations)

	d.db.WithContext(ctx).Model(&model.Donation{}).
		Where("status = ? AND type = ? AND created_at >= ?", model.DonationStatusCompleted, model.DonationTypeMoney, monthStart).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&stats.MonthlyAmount)

	return stats, nil
}

// GetPublicList 获取公开捐赠列表（用于展示墙）
func (d *DonationDAO) GetPublicList(ctx context.Context, limit int) ([]*model.Donation, error) {
	var donations []*model.Donation
	err := d.db.WithContext(ctx).
		Preload("User").
		Preload("Organization").
		Where("status = ?", model.DonationStatusCompleted).
		Order("created_at DESC").
		Limit(limit).
		Find(&donations).Error
	return donations, err
}
