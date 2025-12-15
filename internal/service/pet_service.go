package service

import (
"context"
"errors"
"fmt"
"github.com/redis/go-redis/v9"
"gorm.io/gorm"
"pet-adoption-platform/internal/dao"
"pet-adoption-platform/internal/model"
)

type PetService struct {
dao   *dao.PetDAO
cache *redis.Client
}

func NewPetService(db *gorm.DB, cache *redis.Client) *PetService {
return &PetService{dao: dao.NewPetDAO(db), cache: cache}
}

func (s *PetService) CreatePet(ctx context.Context, userID int64, req *model.PetCreateRequest) (*model.Pet, error) {
photosStr := dao.ConvertPhotosToString(req.Photos)
pet := &model.Pet{
Name: req.Name, Type: req.Type, Breed: req.Breed, Gender: req.Gender,
Age: req.Age, Size: req.Size, Color: req.Color, Weight: req.Weight,
IsVaccinated: req.IsVaccinated, IsSterilized: req.IsSterilized,
HealthStatus: req.HealthStatus, Description: req.Description,
Character: req.Character, Photos: photosStr, CoverPhoto: req.CoverPhoto,
Province: req.Province, City: req.City, District: req.District,
Address: req.Address, UserID: userID, Status: model.PetStatusPending,
}
if err := s.dao.Create(ctx, pet); err != nil {
return nil, fmt.Errorf("创建宠物失败: %w", err)
}
return pet, nil
}

func (s *PetService) GetPetByID(ctx context.Context, id uint, incrementView bool) (*model.PetResponse, error) {
pet, err := s.dao.GetByID(ctx, id)
if err != nil {
if errors.Is(err, gorm.ErrRecordNotFound) {
return nil, errors.New("宠物不存在")
}
return nil, fmt.Errorf("获取宠物失败: %w", err)
}
if incrementView {
_ = s.dao.IncrementViewCount(ctx, id)
}
return s.convertToResponse(pet), nil
}

func (s *PetService) UpdatePet(ctx context.Context, id, userID int64, req *model.PetUpdateRequest) error {
	return s.updatePetInternal(ctx, id, userID, req, false)
}

// AdminUpdatePet 管理员更新宠物信息（不检查所有者）
func (s *PetService) AdminUpdatePet(ctx context.Context, id int64, req *model.PetUpdateRequest) error {
	return s.updatePetInternal(ctx, id, 0, req, true)
}

func (s *PetService) updatePetInternal(ctx context.Context, id, userID int64, req *model.PetUpdateRequest, isAdmin bool) error {
pet, err := s.dao.GetByID(ctx, uint(id))
if err != nil {
if errors.Is(err, gorm.ErrRecordNotFound) {
return errors.New("宠物不存在")
}
return fmt.Errorf("获取宠物失败: %w", err)
}
if !isAdmin && pet.UserID != userID {
return errors.New("无权编辑此宠物")
}
if req.Name != nil { pet.Name = *req.Name }
if req.Type != nil { pet.Type = *req.Type }
if req.Breed != nil { pet.Breed = *req.Breed }
if req.Gender != nil { pet.Gender = *req.Gender }
if req.Age != nil { pet.Age = *req.Age }
if req.Size != nil { pet.Size = *req.Size }
if req.Color != nil { pet.Color = *req.Color }
if req.Weight != nil { pet.Weight = *req.Weight }
if req.IsVaccinated != nil { pet.IsVaccinated = *req.IsVaccinated }
if req.IsSterilized != nil { pet.IsSterilized = *req.IsSterilized }
if req.HealthStatus != nil { pet.HealthStatus = *req.HealthStatus }
if req.Description != nil { pet.Description = *req.Description }
if req.Character != nil { pet.Character = *req.Character }
if req.Photos != nil { pet.Photos = dao.ConvertPhotosToString(req.Photos) }
if req.CoverPhoto != nil { pet.CoverPhoto = *req.CoverPhoto }
if req.Province != nil { pet.Province = *req.Province }
if req.City != nil { pet.City = *req.City }
if req.District != nil { pet.District = *req.District }
if req.Address != nil { pet.Address = *req.Address }
if err := s.dao.Update(ctx, pet); err != nil {
return fmt.Errorf("更新宠物失败: %w", err)
}
if s.cache != nil {
cacheKey := fmt.Sprintf("pet:%d", id)
_ = s.cache.Del(ctx, cacheKey).Err()
}
return nil
}

func (s *PetService) DeletePet(ctx context.Context, id, userID int64, isAdmin bool) error {
pet, err := s.dao.GetByID(ctx, uint(id))
if err != nil {
if errors.Is(err, gorm.ErrRecordNotFound) {
return errors.New("宠物不存在")
}
return fmt.Errorf("获取宠物失败: %w", err)
}
if !isAdmin && pet.UserID != userID {
return errors.New("无权删除此宠物")
}
if err := s.dao.Delete(ctx, uint(id)); err != nil {
return fmt.Errorf("删除宠物失败: %w", err)
}
if s.cache != nil {
cacheKey := fmt.Sprintf("pet:%d", id)
_ = s.cache.Del(ctx, cacheKey).Err()
}
return nil
}

func (s *PetService) ListPets(ctx context.Context, page, pageSize int) ([]*model.PetResponse, int64, error) {
if page < 1 { page = 1 }
if pageSize < 1 || pageSize > 100 { pageSize = 20 }
pets, total, err := s.dao.List(ctx, page, pageSize)
if err != nil {
return nil, 0, fmt.Errorf("获取宠物列表失败: %w", err)
}
responses := make([]*model.PetResponse, len(pets))
for i, pet := range pets {
responses[i] = s.convertToResponse(pet)
}
return responses, total, nil
}

func (s *PetService) QueryPets(ctx context.Context, req *model.PetQueryRequest) ([]*model.PetResponse, int64, error) {
if req.Page < 1 { req.Page = 1 }
if req.PageSize < 1 || req.PageSize > 100 { req.PageSize = 20 }
pets, total, err := s.dao.QueryByConditions(ctx, req)
if err != nil {
return nil, 0, fmt.Errorf("查询宠物失败: %w", err)
}
responses := make([]*model.PetResponse, len(pets))
for i, pet := range pets {
responses[i] = s.convertToResponse(pet)
}
return responses, total, nil
}

func (s *PetService) SearchPets(ctx context.Context, keyword string, page, pageSize int) ([]*model.PetResponse, int64, error) {
if page < 1 { page = 1 }
if pageSize < 1 || pageSize > 100 { pageSize = 20 }
pets, total, err := s.dao.SearchByKeyword(ctx, keyword, page, pageSize)
if err != nil {
return nil, 0, fmt.Errorf("搜索宠物失败: %w", err)
}
responses := make([]*model.PetResponse, len(pets))
for i, pet := range pets {
responses[i] = s.convertToResponse(pet)
}
return responses, total, nil
}

func (s *PetService) GetMyPets(ctx context.Context, userID int64, page, pageSize int) ([]*model.PetResponse, int64, error) {
if page < 1 { page = 1 }
if pageSize < 1 || pageSize > 100 { pageSize = 20 }
pets, total, err := s.dao.GetByUserID(ctx, userID, page, pageSize)
if err != nil {
return nil, 0, fmt.Errorf("获取我的宠物失败: %w", err)
}
responses := make([]*model.PetResponse, len(pets))
for i, pet := range pets {
responses[i] = s.convertToResponse(pet)
}
return responses, total, nil
}

func (s *PetService) GetPendingPets(ctx context.Context, page, pageSize int) ([]*model.PetResponse, int64, error) {
if page < 1 { page = 1 }
if pageSize < 1 || pageSize > 100 { pageSize = 20 }
pets, total, err := s.dao.GetByStatus(ctx, model.PetStatusPending, page, pageSize)
if err != nil {
return nil, 0, fmt.Errorf("获取待审核宠物失败: %w", err)
}
responses := make([]*model.PetResponse, len(pets))
for i, pet := range pets {
responses[i] = s.convertToResponse(pet)
}
return responses, total, nil
}

func (s *PetService) ApprovePet(ctx context.Context, petID, reviewerID int64) error {
pet, err := s.dao.GetByID(ctx, uint(petID))
if err != nil {
if errors.Is(err, gorm.ErrRecordNotFound) {
return errors.New("宠物不存在")
}
return fmt.Errorf("获取宠物失败: %w", err)
}
if pet.Status != model.PetStatusPending {
return errors.New("该宠物不是待审核状态")
}
if err := s.dao.SetReviewed(ctx, petID, reviewerID, model.PetStatusAvailable, ""); err != nil {
return fmt.Errorf("审核失败: %w", err)
}
return nil
}

func (s *PetService) RejectPet(ctx context.Context, petID, reviewerID int64, reason string) error {
pet, err := s.dao.GetByID(ctx, uint(petID))
if err != nil {
if errors.Is(err, gorm.ErrRecordNotFound) {
return errors.New("宠物不存在")
}
return fmt.Errorf("获取宠物失败: %w", err)
}
if pet.Status != model.PetStatusPending {
return errors.New("该宠物不是待审核状态")
}
if err := s.dao.SetReviewed(ctx, petID, reviewerID, model.PetStatusOffline, reason); err != nil {
return fmt.Errorf("拒绝失败: %w", err)
}
return nil
}

func (s *PetService) OfflinePet(ctx context.Context, petID, userID int64, isAdmin bool) error {
pet, err := s.dao.GetByID(ctx, uint(petID))
if err != nil {
if errors.Is(err, gorm.ErrRecordNotFound) {
return errors.New("宠物不存在")
}
return fmt.Errorf("获取宠物失败: %w", err)
}
if !isAdmin && pet.UserID != userID {
return errors.New("无权下架此宠物")
}
if err := s.dao.UpdateStatus(ctx, uint(petID), model.PetStatusOffline); err != nil {
return fmt.Errorf("下架失败: %w", err)
}
return nil
}

func (s *PetService) GetRecommendedPets(ctx context.Context, limit int) ([]*model.PetResponse, error) {
if limit < 1 || limit > 50 { limit = 10 }
pets, err := s.dao.GetRecommended(ctx, limit)
if err != nil {
return nil, fmt.Errorf("获取推荐宠物失败: %w", err)
}
responses := make([]*model.PetResponse, len(pets))
for i, pet := range pets {
responses[i] = s.convertToResponse(pet)
}
return responses, nil
}

func (s *PetService) GetStatistics(ctx context.Context) (*dao.PetStatistics, error) {
	return s.dao.GetStatistics(ctx)
}

func (s *PetService) convertToResponse(pet *model.Pet) *model.PetResponse {
resp := &model.PetResponse{
ID: pet.ID, Name: pet.Name, Type: pet.Type, Breed: pet.Breed, Gender: pet.Gender,
Age: pet.Age, Size: pet.Size, Color: pet.Color, Weight: pet.Weight,
IsVaccinated: pet.IsVaccinated, IsSterilized: pet.IsSterilized,
HealthStatus: pet.HealthStatus, Description: pet.Description,
Character: pet.Character, Photos: dao.ConvertPhotosToSlice(pet.Photos),
CoverPhoto: pet.CoverPhoto, Province: pet.Province, City: pet.City,
District: pet.District, Address: pet.Address, Status: pet.Status,
ViewCount: pet.ViewCount, FavoriteCount: pet.FavoriteCount,
UserID: uint(pet.UserID), AdoptedBy: convertInt64PtrToUintPtr(pet.AdoptedBy),
AdoptedAt: pet.AdoptedAt, CreatedAt: pet.CreatedAt, UpdatedAt: pet.UpdatedAt,
}
if pet.User != nil {
resp.UserInfo = &model.UserBasicInfo{
ID: uint(pet.User.ID), Username: pet.User.Username,
Nickname: pet.User.RealName, Phone: pet.User.Phone,
Role: string(pet.User.Role),
}
}
return resp
}

func convertInt64PtrToUintPtr(val *int64) *uint {
if val == nil {
return nil
}
u := uint(*val)
return &u
}
