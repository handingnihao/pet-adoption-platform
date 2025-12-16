package service

import (
	"context"
	"errors"
	"pet-adoption-platform/internal/dao"
	"pet-adoption-platform/internal/model"
	"pet-adoption-platform/pkg/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// OrganizationService 机构服务
type OrganizationService struct {
	orgDAO *dao.OrganizationDAO
}

// NewOrganizationService 创建机构服务实例
func NewOrganizationService(db *gorm.DB) *OrganizationService {
	return &OrganizationService{
		orgDAO: dao.NewOrganizationDAO(db),
	}
}

// CreateOrganization 创建机构
func (s *OrganizationService) CreateOrganization(ctx context.Context, req *model.OrganizationCreateRequest, userID uint64) (*model.Organization, error) {
	// 检查机构名称是否已存在
	exists, err := s.orgDAO.CheckNameExists(ctx, req.Name)
	if err != nil {
		logger.Error("检查机构名称是否存在失败", zap.Error(err))
		return nil, err
	}
	if exists {
		return nil, errors.New("机构名称已存在")
	}

	// 创建机构
	org := &model.Organization{
		Name:        req.Name,
		Logo:        req.Logo,
		Description: req.Description,
		Address:     req.Address,
		Phone:       req.Phone,
		Email:       req.Email,
		Status:      model.OrganizationStatusPending, // 默认待审核
		CreatedBy:   userID,
	}

	if err := s.orgDAO.Create(ctx, org); err != nil {
		logger.Error("创建机构失败", zap.Error(err))
		return nil, err
	}

	return org, nil
}

// UpdateOrganization 更新机构信息
func (s *OrganizationService) UpdateOrganization(ctx context.Context, id uint64, req *model.OrganizationUpdateRequest, userID uint64) error {
	// 检查机构是否存在
	org, err := s.orgDAO.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if org == nil {
		return errors.New("机构不存在")
	}

	// 如果更新了名称，检查新名称是否已存在
	if req.Name != nil && *req.Name != org.Name {
		exists, err := s.orgDAO.CheckNameExists(ctx, *req.Name, id)
		if err != nil {
			logger.Error("检查机构名称是否存在失败", zap.Error(err))
			return err
		}
		if exists {
			return errors.New("机构名称已存在")
		}
	}

	// 构建更新字段
	updates := make(map[string]interface{})
	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.Logo != nil {
		updates["logo"] = *req.Logo
	}
	if req.Description != nil {
		updates["description"] = *req.Description
	}
	if req.Address != nil {
		updates["address"] = *req.Address
	}
	if req.Phone != nil {
		updates["phone"] = *req.Phone
	}
	if req.Email != nil {
		updates["email"] = *req.Email
	}
	updates["updated_by"] = userID

	// 更新机构信息
	return s.orgDAO.Update(ctx, id, updates)
}

// GetOrganization 获取机构详情
func (s *OrganizationService) GetOrganization(ctx context.Context, id uint64) (*model.Organization, error) {
	return s.orgDAO.GetByID(ctx, id)
}

// ListOrganizations 获取机构列表
func (s *OrganizationService) ListOrganizations(ctx context.Context, page, pageSize int, status *model.OrganizationStatus) ([]*model.Organization, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	return s.orgDAO.List(ctx, page, pageSize, status)
}

// DeleteOrganization 删除机构
func (s *OrganizationService) DeleteOrganization(ctx context.Context, id uint64) error {
	// 检查机构是否存在
	org, err := s.orgDAO.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if org == nil {
		return errors.New("机构不存在")
	}

	return s.orgDAO.Delete(ctx, id)
}

// UpdateOrganizationStatus 更新机构状态
func (s *OrganizationService) UpdateOrganizationStatus(ctx context.Context, id uint64, req *model.OrganizationStatusUpdateRequest, userID uint64) error {
	// 检查机构是否存在
	org, err := s.orgDAO.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if org == nil {
		return errors.New("机构不存在")
	}

	// 如果状态没有变化，直接返回
	if org.Status == req.Status {
		return nil
	}

	// 更新状态
	rejectReason := ""
	if req.Status == model.OrganizationStatusRejected {
		rejectReason = req.RejectReason
	}

	return s.orgDAO.UpdateStatus(ctx, id, req.Status, rejectReason)
}

// GetUserOrganizations 获取用户创建的机构列表
func (s *OrganizationService) GetUserOrganizations(ctx context.Context, userID uint64, page, pageSize int) ([]*model.Organization, int64, error) {
	if userID <= 0 {
		return nil, 0, errors.New("invalid user id")
	}

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	return s.orgDAO.ListByUserID(ctx, userID, page, pageSize)
}

// CreateOrganizationForRegistration 组织注册（创建待审核的组织）
func (s *OrganizationService) CreateOrganizationForRegistration(ctx context.Context, org *model.Organization) error {
	if org == nil {
		return errors.New("organization cannot be nil")
	}

	// 检查机构名称是否已存在
	exists, err := s.orgDAO.CheckNameExists(ctx, org.Name)
	if err != nil {
		logger.Error("检查机构名称是否存在失败", zap.Error(err))
		return err
	}
	if exists {
		return errors.New("机构名称已存在")
	}

	// 确保状态为待审核
	org.Status = 0

	// 创建机构
	if err := s.orgDAO.Create(ctx, org); err != nil {
		logger.Error("创建组织失败", zap.Error(err))
		return err
	}

	logger.Info("组织注册申请已创建",
		zap.String("name", org.Name),
		zap.String("contact_name", org.ContactName))

	return nil
}
