package service

import (
	"context"
	"errors"
	"fmt"
	"pet-adoption-platform/internal/dao"
	"pet-adoption-platform/internal/model"
	"time"

	"gorm.io/gorm"
)

type AdoptionService struct {
	adoptionDAO *dao.AdoptionDAO
	petDAO      *dao.PetDAO
	userDAO     *dao.UserDAO
}

func NewAdoptionService(db *gorm.DB) *AdoptionService {
	return &AdoptionService{
		adoptionDAO: dao.NewAdoptionDAO(db),
		petDAO:      dao.NewPetDAO(db),
		userDAO:     dao.NewUserDAO(db),
	}
}

// ========== Application相关 ==========

// CreateApplication 创建领养申请
func (s *AdoptionService) CreateApplication(ctx context.Context, userID int64, req *model.ApplicationCreateRequest) (*model.AdoptionApplication, error) {
	// 1. 检查宠物是否存在且可领养
	pet, err := s.petDAO.GetByID(ctx, uint(req.PetID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("宠物不存在")
		}
		return nil, err
	}

	if pet.Status != model.PetStatusAvailable {
		return nil, errors.New("该宠物当前不可领养")
	}

	// 2. 检查是否已有进行中的申请
	hasDuplicate, err := s.adoptionDAO.CheckDuplicateApplication(ctx, userID, req.PetID)
	if err != nil {
		return nil, err
	}
	if hasDuplicate {
		return nil, errors.New("您已经申请过该宠物，请等待审核")
	}

	// 3. 生成申请编号
	applicationNo := fmt.Sprintf("APP%d%d", time.Now().Unix(), userID%1000)

	// 4. 创建申请
	app := &model.AdoptionApplication{
		ApplicationNo:  applicationNo,
		UserID:         userID,
		PetID:          req.PetID,
		OrganizationID: req.OrganizationID,
		ApplicantName:  req.ApplicantName,
		ApplicantPhone: req.ApplicantPhone,
		ApplicantAddr:  req.ApplicantAddr,
		HousingType:    req.HousingType,
		HousingArea:    req.HousingArea,
		HasYard:        req.HasYard,
		FamilyMembers:  req.FamilyMembers,
		HasChildren:    req.HasChildren,
		ChildrenAge:    req.ChildrenAge,
		FamilyAgree:    req.FamilyAgree,
		AdoptionReason: req.AdoptionReason,
		HowToCare:      req.HowToCare,
		EmergencyPlan:  req.EmergencyPlan,
		Status:         model.ApplicationStatusPending,
	}

	if err := s.adoptionDAO.CreateApplication(ctx, app); err != nil {
		return nil, err
	}

	return app, nil
}

// GetApplicationByID 获取申请详情
func (s *AdoptionService) GetApplicationByID(ctx context.Context, id uint64, userID int64, isAdmin bool) (*model.ApplicationResponse, error) {
	app, err := s.adoptionDAO.GetApplicationByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("申请不存在")
		}
		return nil, err
	}

	// 权限检查：只有申请人或管理员可以查看
	if !isAdmin && app.UserID != userID {
		return nil, errors.New("无权查看该申请")
	}

	return s.convertToApplicationResponse(app), nil
}

// UpdateApplication 更新申请
func (s *AdoptionService) UpdateApplication(ctx context.Context, id uint64, userID int64, req *model.ApplicationUpdateRequest) error {
	// 1. 获取申请
	app, err := s.adoptionDAO.GetApplicationByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("申请不存在")
		}
		return err
	}

	// 2. 权限检查
	if app.UserID != userID {
		return errors.New("无权修改该申请")
	}

	// 3. 状态检查：只有待审核状态才能修改
	if app.Status != model.ApplicationStatusPending {
		return errors.New("该申请已进入审核流程，无法修改")
	}

	// 4. 准备更新字段
	updates := make(map[string]interface{})
	if req.ApplicantName != nil {
		updates["applicant_name"] = *req.ApplicantName
	}
	if req.ApplicantPhone != nil {
		updates["applicant_phone"] = *req.ApplicantPhone
	}
	if req.ApplicantAddr != nil {
		updates["applicant_address"] = *req.ApplicantAddr
	}
	if req.HousingType != nil {
		updates["housing_type"] = *req.HousingType
	}
	if req.HousingArea != nil {
		updates["housing_area"] = *req.HousingArea
	}
	if req.HasYard != nil {
		updates["has_yard"] = *req.HasYard
	}
	if req.FamilyMembers != nil {
		updates["family_members"] = *req.FamilyMembers
	}
	if req.HasChildren != nil {
		updates["has_children"] = *req.HasChildren
	}
	if req.ChildrenAge != nil {
		updates["children_age"] = *req.ChildrenAge
	}
	if req.FamilyAgree != nil {
		updates["family_agree"] = *req.FamilyAgree
	}
	if req.HasPetExperience != nil {
		updates["has_pet_experience"] = *req.HasPetExperience
	}
	if req.PetExperience != nil {
		updates["pet_experience"] = *req.PetExperience
	}
	if req.CurrentPets != nil {
		updates["current_pets"] = *req.CurrentPets
	}
	if req.AdoptionReason != nil {
		updates["adoption_reason"] = *req.AdoptionReason
	}
	if req.HowToCare != nil {
		updates["how_to_care"] = *req.HowToCare
	}
	if req.EmergencyPlan != nil {
		updates["emergency_plan"] = *req.EmergencyPlan
	}

	if len(updates) == 0 {
		return errors.New("没有需要更新的字段")
	}

	return s.adoptionDAO.UpdateApplication(ctx, id, updates)
}

// CancelApplication 取消申请
func (s *AdoptionService) CancelApplication(ctx context.Context, id uint64, userID int64) error {
	app, err := s.adoptionDAO.GetApplicationByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("申请不存在")
		}
		return err
	}

	if app.UserID != userID {
		return errors.New("无权取消该申请")
	}

	// 已通过或已拒绝的申请不能取消
	if app.Status == model.ApplicationStatusApproved || app.Status == model.ApplicationStatusRejected {
		return errors.New("该申请已完成审核，无法取消")
	}

	updates := map[string]interface{}{
		"status": model.ApplicationStatusCancelled,
	}
	return s.adoptionDAO.UpdateApplication(ctx, id, updates)
}

// ReviewApplication 审核申请
func (s *AdoptionService) ReviewApplication(ctx context.Context, id uint64, reviewerID int64, req *model.ApplicationReviewRequest) error {
	// 1. 获取申请
	app, err := s.adoptionDAO.GetApplicationByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("申请不存在")
		}
		return err
	}

	// 2. 状态检查
	if app.Status == model.ApplicationStatusApproved || app.Status == model.ApplicationStatusRejected || app.Status == model.ApplicationStatusCancelled {
		return errors.New("该申请已完成，无法再次审核")
	}

	// 3. 根据操作更新状态
	var newStatus model.ApplicationStatus
	updates := make(map[string]interface{})

	switch req.Action {
	case "approve":
		newStatus = model.ApplicationStatusApproved
		now := time.Now()
		updates["approved_at"] = now
	case "reject":
		newStatus = model.ApplicationStatusRejected
		now := time.Now()
		updates["rejected_at"] = now
		updates["rejection_reason"] = req.Reason
	case "interview":
		newStatus = model.ApplicationStatusInterview
		updates["interview_time"] = time.Now().Add(24 * time.Hour)
	case "home_visit":
		newStatus = model.ApplicationStatusHomeVisit
		updates["home_visit_time"] = time.Now().Add(48 * time.Hour)
	default:
		return errors.New("无效的审核操作")
	}

	updates["status"] = newStatus
	updates["reviewer_id"] = reviewerID
	if req.Comment != "" {
		updates["review_comment"] = req.Comment
	}

	return s.adoptionDAO.UpdateApplication(ctx, id, updates)
}

// GetMyApplications 获取我的申请列表
func (s *AdoptionService) GetMyApplications(ctx context.Context, userID int64, page, pageSize int) ([]*model.ApplicationResponse, int64, error) {
	apps, total, err := s.adoptionDAO.GetApplicationsByUserID(ctx, userID, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]*model.ApplicationResponse, len(apps))
	for i, app := range apps {
		responses[i] = s.convertToApplicationResponse(app)
	}

	return responses, total, nil
}

// ListApplications 管理员查看所有申请
func (s *AdoptionService) ListApplications(ctx context.Context, page, pageSize int) ([]*model.ApplicationResponse, int64, error) {
	apps, total, err := s.adoptionDAO.ListApplications(ctx, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]*model.ApplicationResponse, len(apps))
	for i, app := range apps {
		responses[i] = s.convertToApplicationResponse(app)
	}

	return responses, total, nil
}

// GetPendingApplications 获取待审核申请
func (s *AdoptionService) GetPendingApplications(ctx context.Context, page, pageSize int) ([]*model.ApplicationResponse, int64, error) {
	apps, total, err := s.adoptionDAO.GetApplicationsByStatus(ctx, model.ApplicationStatusPending, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]*model.ApplicationResponse, len(apps))
	for i, app := range apps {
		responses[i] = s.convertToApplicationResponse(app)
	}

	return responses, total, nil
}

// QueryApplications 条件查询申请
func (s *AdoptionService) QueryApplications(ctx context.Context, query *model.ApplicationQueryRequest) ([]*model.ApplicationResponse, int64, error) {
	apps, total, err := s.adoptionDAO.QueryApplications(ctx, query)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]*model.ApplicationResponse, len(apps))
	for i, app := range apps {
		responses[i] = s.convertToApplicationResponse(app)
	}

	return responses, total, nil
}

// ========== Adoption相关 ==========

// CreateAdoption 创建领养记录（通过申请后）
func (s *AdoptionService) CreateAdoption(ctx context.Context, req *model.AdoptionCreateRequest) (*model.Adoption, error) {
	// 1. 检查申请是否已通过
	app, err := s.adoptionDAO.GetApplicationByID(ctx, req.ApplicationID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("申请不存在")
		}
		return nil, err
	}

	if app.Status != model.ApplicationStatusApproved {
		return nil, errors.New("该申请未通过审核")
	}

	// 2. 检查是否已创建领养记录
	_, err = s.adoptionDAO.GetAdoptionByApplicationID(ctx, req.ApplicationID)
	if err == nil {
		return nil, errors.New("该申请已创建领养记录")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// 3. 创建领养记录
	adoption := &model.Adoption{
		ApplicationID:    req.ApplicationID,
		UserID:           app.UserID,
		PetID:            app.PetID,
		OrganizationID:   app.OrganizationID,
		AdoptionDate:     req.AdoptionDate,
		HandoverLocation: req.HandoverLocation,
		AgreementURL:     req.AgreementURL,
		Notes:            req.Notes,
		Status:           model.AdoptionStatusActive,
	}

	if err := s.adoptionDAO.CreateAdoption(ctx, adoption); err != nil {
		return nil, err
	}

	// 4. 更新宠物状态为已领养
	petUpdates := map[string]interface{}{
		"status":     model.PetStatusAdopted,
		"adopted_by": app.UserID,
		"adopted_at": time.Now(),
	}
	if err := s.petDAO.UpdateByID(ctx, app.PetID, petUpdates); err != nil {
		return nil, err
	}

	return adoption, nil
}

// GetAdoptionByID 获取领养记录
func (s *AdoptionService) GetAdoptionByID(ctx context.Context, id uint64, userID int64, isAdmin bool) (*model.AdoptionResponse, error) {
	adoption, err := s.adoptionDAO.GetAdoptionByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("领养记录不存在")
		}
		return nil, err
	}

	// 权限检查
	if !isAdmin && adoption.UserID != userID {
		return nil, errors.New("无权查看该记录")
	}

	return s.convertToAdoptionResponse(adoption), nil
}

// GetMyAdoptions 获取我的领养记录
func (s *AdoptionService) GetMyAdoptions(ctx context.Context, userID int64, page, pageSize int) ([]*model.AdoptionResponse, int64, error) {
	adoptions, total, err := s.adoptionDAO.GetAdoptionsByUserID(ctx, userID, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]*model.AdoptionResponse, len(adoptions))
	for i, adoption := range adoptions {
		responses[i] = s.convertToAdoptionResponse(adoption)
	}

	return responses, total, nil
}

// ListAdoptions 管理员查看所有领养记录
func (s *AdoptionService) ListAdoptions(ctx context.Context, page, pageSize int) ([]*model.AdoptionResponse, int64, error) {
	adoptions, total, err := s.adoptionDAO.ListAdoptions(ctx, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]*model.AdoptionResponse, len(adoptions))
	for i, adoption := range adoptions {
		responses[i] = s.convertToAdoptionResponse(adoption)
	}

	return responses, total, nil
}

// UpdateAdoptionStatus 更新领养状态
func (s *AdoptionService) UpdateAdoptionStatus(ctx context.Context, id uint64, status model.AdoptionStatus, notes string) error {
	adoption, err := s.adoptionDAO.GetAdoptionByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("领养记录不存在")
		}
		return err
	}

	updates := map[string]interface{}{
		"status": status,
	}
	if notes != "" {
		updates["notes"] = notes
	}

	// 如果退回，更新宠物状态
	if status == model.AdoptionStatusReturned {
		petUpdates := map[string]interface{}{
			"status":     model.PetStatusAvailable,
			"adopted_by": nil,
			"adopted_at": nil,
		}
		if err := s.petDAO.UpdateByID(ctx, adoption.PetID, petUpdates); err != nil {
			return err
		}
	}

	return s.adoptionDAO.UpdateAdoption(ctx, id, updates)
}

// GetStatistics 获取统计信息
func (s *AdoptionService) GetStatistics(ctx context.Context) (map[string]interface{}, error) {
	return s.adoptionDAO.GetStatistics(ctx)
}

// ========== 辅助方法 ==========

func (s *AdoptionService) convertToApplicationResponse(app *model.AdoptionApplication) *model.ApplicationResponse {
	resp := &model.ApplicationResponse{
		ID:              app.ID,
		ApplicationNo:   app.ApplicationNo,
		UserID:          app.UserID,
		PetID:           app.PetID,
		OrganizationID:  app.OrganizationID,
		ApplicantName:   app.ApplicantName,
		ApplicantPhone:  app.ApplicantPhone,
		ApplicantAddr:   app.ApplicantAddr,
		HousingType:     app.HousingType,
		HousingArea:     app.HousingArea,
		HasYard:         app.HasYard,
		FamilyMembers:   app.FamilyMembers,
		HasChildren:     app.HasChildren,
		FamilyAgree:     app.FamilyAgree,
		AdoptionReason:  app.AdoptionReason,
		HowToCare:       app.HowToCare,
		EmergencyPlan:   app.EmergencyPlan,
		Status:          app.Status,
		ReviewComment:   app.ReviewComment,
		RejectionReason: app.RejectionReason,
		InterviewTime:   app.InterviewTime,
		HomeVisitTime:   app.HomeVisitTime,
		ApprovedAt:      app.ApprovedAt,
		RejectedAt:      app.RejectedAt,
		CreatedAt:       app.CreatedAt,
		UpdatedAt:       app.UpdatedAt,
	}

	// Pet信息
	if app.Pet != nil {
		resp.PetInfo = &model.PetBasicInfo{
			ID:         app.Pet.ID,
			Name:       app.Pet.Name,
			Type:       app.Pet.Type,
			Breed:      app.Pet.Breed,
			Age:        app.Pet.Age,
			Gender:     app.Pet.Gender,
			CoverPhoto: app.Pet.CoverPhoto,
			Status:     app.Pet.Status,
		}
	}

	// User信息
	if app.User != nil {
		resp.UserInfo = &model.UserBasicInfo{
			ID:       uint(app.User.ID),
			Username: app.User.Username,
			Nickname: app.User.RealName,
			Phone:    app.User.Phone,
			Role:     string(app.User.Role),
		}
	}

	// Reviewer信息
	if app.Reviewer != nil {
		resp.ReviewerInfo = &model.UserBasicInfo{
			ID:       uint(app.Reviewer.ID),
			Username: app.Reviewer.Username,
			Nickname: app.Reviewer.RealName,
			Role:     string(app.Reviewer.Role),
		}
	}

	return resp
}

func (s *AdoptionService) convertToAdoptionResponse(adoption *model.Adoption) *model.AdoptionResponse {
	resp := &model.AdoptionResponse{
		ID:               adoption.ID,
		ApplicationID:    adoption.ApplicationID,
		UserID:           adoption.UserID,
		PetID:            adoption.PetID,
		OrganizationID:   adoption.OrganizationID,
		AdoptionDate:     adoption.AdoptionDate,
		HandoverLocation: adoption.HandoverLocation,
		AgreementURL:     adoption.AgreementURL,
		AgreementSigned:  adoption.AgreementSigned,
		Status:           adoption.Status,
		Notes:            adoption.Notes,
		CreatedAt:        adoption.CreatedAt,
		UpdatedAt:        adoption.UpdatedAt,
	}

	// Pet信息
	if adoption.Pet != nil {
		resp.PetInfo = &model.PetBasicInfo{
			ID:         adoption.Pet.ID,
			Name:       adoption.Pet.Name,
			Type:       adoption.Pet.Type,
			Breed:      adoption.Pet.Breed,
			Age:        adoption.Pet.Age,
			Gender:     adoption.Pet.Gender,
			CoverPhoto: adoption.Pet.CoverPhoto,
			Status:     adoption.Pet.Status,
		}
	}

	// User信息
	if adoption.User != nil {
		resp.UserInfo = &model.UserBasicInfo{
			ID:       uint(adoption.User.ID),
			Username: adoption.User.Username,
			Nickname: adoption.User.RealName,
			Phone:    adoption.User.Phone,
			Role:     string(adoption.User.Role),
		}
	}

	return resp
}
