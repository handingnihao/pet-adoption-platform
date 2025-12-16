package service

import (
	"context"
	"encoding/json"
	"errors"
	"pet-adoption-platform/internal/dao"
	"pet-adoption-platform/internal/model"
	"pet-adoption-platform/pkg/logger"

	"go.uber.org/zap"
)

var (
	ErrDonationNotFound    = errors.New("捐赠记录不存在")
	ErrInvalidDonationType = errors.New("无效的捐赠类型")
	ErrInvalidAmount       = errors.New("资金捐赠金额必须大于0")
)

// DonationService 捐赠服务
type DonationService struct {
	donationDAO *dao.DonationDAO
	orgDAO      *dao.OrganizationDAO
}

// NewDonationService 创建DonationService实例
func NewDonationService(donationDAO *dao.DonationDAO, orgDAO *dao.OrganizationDAO) *DonationService {
	return &DonationService{
		donationDAO: donationDAO,
		orgDAO:      orgDAO,
	}
}

// CreateDonation 创建捐赠
func (s *DonationService) CreateDonation(ctx context.Context, req *model.DonationCreateRequest, userID uint64) (*model.Donation, error) {
	// 验证捐赠类型
	switch req.Type {
	case model.DonationTypeMoney:
		if req.Amount <= 0 {
			return nil, ErrInvalidAmount
		}
	case model.DonationTypeSupply:
		if len(req.SupplyItems) == 0 {
			return nil, errors.New("物资捐赠必须填写物资清单")
		}
	case model.DonationTypeService:
		if req.ServiceDesc == "" {
			return nil, errors.New("服务捐赠必须填写服务描述")
		}
	default:
		return nil, ErrInvalidDonationType
	}

	// 验证机构是否存在
	if req.OrganizationID != nil {
		org, err := s.orgDAO.GetByID(ctx, *req.OrganizationID)
		if err != nil {
			return nil, err
		}
		if org == nil {
			return nil, errors.New("指定的机构不存在")
		}
	}

	// 物资清单转JSON
	var supplyItemsJSON string
	if len(req.SupplyItems) > 0 {
		data, _ := json.Marshal(req.SupplyItems)
		supplyItemsJSON = string(data)
	}

	donation := &model.Donation{
		UserID:         userID,
		OrganizationID: req.OrganizationID,
		Type:           req.Type,
		Amount:         req.Amount,
		SupplyItems:    supplyItemsJSON,
		ServiceDesc:    req.ServiceDesc,
		Message:        req.Message,
		IsAnonymous:    req.IsAnonymous,
		PaymentMethod:  req.PaymentMethod,
		Status:         model.DonationStatusPending,
	}

	if err := s.donationDAO.Create(ctx, donation); err != nil {
		logger.Error("创建捐赠失败", zap.Error(err))
		return nil, err
	}

	logger.Info("创建捐赠成功",
		zap.Uint64("donation_id", donation.ID),
		zap.Uint64("user_id", userID),
		zap.String("type", string(req.Type)))

	return donation, nil
}

// GetDonation 获取捐赠详情
func (s *DonationService) GetDonation(ctx context.Context, id uint64) (*model.DonationInfo, error) {
	donation, err := s.donationDAO.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if donation == nil {
		return nil, ErrDonationNotFound
	}
	return donation.ToInfo(), nil
}

// ListDonations 获取捐赠列表
func (s *DonationService) ListDonations(ctx context.Context, page, pageSize int, donationType *model.DonationType, status *model.DonationStatus) ([]*model.DonationInfo, int64, error) {
	donations, total, err := s.donationDAO.List(ctx, page, pageSize, donationType, status)
	if err != nil {
		return nil, 0, err
	}

	infos := make([]*model.DonationInfo, len(donations))
	for i, d := range donations {
		infos[i] = d.ToInfo()
	}

	return infos, total, nil
}

// GetMyDonations 获取我的捐赠记录
func (s *DonationService) GetMyDonations(ctx context.Context, userID uint64, page, pageSize int) ([]*model.DonationInfo, int64, error) {
	donations, total, err := s.donationDAO.ListByUserID(ctx, userID, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	infos := make([]*model.DonationInfo, len(donations))
	for i, d := range donations {
		infos[i] = d.ToInfo()
	}

	return infos, total, nil
}

// GetOrganizationDonations 获取机构收到的捐赠
func (s *DonationService) GetOrganizationDonations(ctx context.Context, orgID uint64, page, pageSize int) ([]*model.DonationInfo, int64, error) {
	donations, total, err := s.donationDAO.ListByOrganizationID(ctx, orgID, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	infos := make([]*model.DonationInfo, len(donations))
	for i, d := range donations {
		infos[i] = d.ToInfo()
	}

	return infos, total, nil
}

// ConfirmDonation 确认捐赠（用户确认支付）
func (s *DonationService) ConfirmDonation(ctx context.Context, id uint64, userID uint64, transactionID string) error {
	donation, err := s.donationDAO.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if donation == nil {
		return ErrDonationNotFound
	}
	if donation.UserID != userID {
		return errors.New("无权操作此捐赠")
	}
	if donation.Status != model.DonationStatusPending {
		return errors.New("捐赠状态不允许确认")
	}

	// 更新交易ID
	if transactionID != "" {
		if err := s.donationDAO.UpdateTransactionID(ctx, id, transactionID); err != nil {
			return err
		}
	}

	// 更新状态为已确认
	return s.donationDAO.UpdateStatus(ctx, id, model.DonationStatusConfirmed, nil)
}

// CompleteDonation 完成捐赠（管理员操作）
func (s *DonationService) CompleteDonation(ctx context.Context, id uint64, adminID uint64) error {
	donation, err := s.donationDAO.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if donation == nil {
		return ErrDonationNotFound
	}
	if donation.Status != model.DonationStatusConfirmed {
		return errors.New("只能完成已确认的捐赠")
	}

	return s.donationDAO.UpdateStatus(ctx, id, model.DonationStatusCompleted, &adminID)
}

// CancelDonation 取消捐赠
func (s *DonationService) CancelDonation(ctx context.Context, id uint64, userID uint64) error {
	donation, err := s.donationDAO.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if donation == nil {
		return ErrDonationNotFound
	}
	if donation.UserID != userID {
		return errors.New("无权操作此捐赠")
	}
	if donation.Status != model.DonationStatusPending {
		return errors.New("只能取消待确认的捐赠")
	}

	return s.donationDAO.UpdateStatus(ctx, id, model.DonationStatusCancelled, nil)
}

// GetStatistics 获取捐赠统计
func (s *DonationService) GetStatistics(ctx context.Context) (*model.DonationStatistics, error) {
	return s.donationDAO.GetStatistics(ctx)
}

// GetPublicDonations 获取公开捐赠列表（爱心墙）
func (s *DonationService) GetPublicDonations(ctx context.Context, limit int) ([]*model.DonationInfo, error) {
	donations, err := s.donationDAO.GetPublicList(ctx, limit)
	if err != nil {
		return nil, err
	}

	infos := make([]*model.DonationInfo, len(donations))
	for i, d := range donations {
		infos[i] = d.ToInfo()
	}

	return infos, nil
}
