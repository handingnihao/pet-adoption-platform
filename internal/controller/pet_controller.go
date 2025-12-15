package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"pet-adoption-platform/internal/model"
	"pet-adoption-platform/internal/service"
	"pet-adoption-platform/pkg/response"
)

type PetController struct {
	service *service.PetService
}

func NewPetController(db *gorm.DB, cache *redis.Client) *PetController {
	return &PetController{
		service: service.NewPetService(db, cache),
	}
}

func (c *PetController) CreatePet(ctx *gin.Context) {
	var req model.PetCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ParamError(ctx, err.Error())
		return
	}

	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Unauthorized(ctx, "未登录")
		return
	}

	pet, err := c.service.CreatePet(ctx.Request.Context(), userID.(int64), &req)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, gin.H{
		"id":      pet.ID,
		"message": "宠物创建成功，等待审核",
	})
}

func (c *PetController) GetPetDetail(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.ParamError(ctx, "无效的宠物ID")
		return
	}

	pet, err := c.service.GetPetByID(ctx.Request.Context(), uint(id), true)
	if err != nil {
		response.NotFound(ctx, err.Error())
		return
	}

	response.Success(ctx, pet)
}

func (c *PetController) UpdatePet(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.ParamError(ctx, "无效的宠物ID")
		return
	}

	var req model.PetUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ParamError(ctx, err.Error())
		return
	}

	userID, _ := ctx.Get("user_id")
	if err := c.service.UpdatePet(ctx.Request.Context(), int64(id), userID.(int64), &req); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, nil)
}

func (c *PetController) DeletePet(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.ParamError(ctx, "无效的宠物ID")
		return
	}

	userID, _ := ctx.Get("user_id")
	role, _ := ctx.Get("role")
	isAdmin := role == string(model.RoleAdmin)

	if err := c.service.DeletePet(ctx.Request.Context(), int64(id), userID.(int64), isAdmin); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, nil)
}

// AdminUpdatePet 管理员更新宠物信息
func (c *PetController) AdminUpdatePet(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.ParamError(ctx, "无效的宠物ID")
		return
	}

	var req model.PetUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ParamError(ctx, err.Error())
		return
	}

	if err := c.service.AdminUpdatePet(ctx.Request.Context(), int64(id), &req); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, nil)
}

func (c *PetController) ListPets(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	pets, total, err := c.service.ListPets(ctx.Request.Context(), page, pageSize)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.SuccessWithPagination(ctx, pets, page, pageSize, total)
}

func (c *PetController) QueryPets(ctx *gin.Context) {
	var req model.PetQueryRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.ParamError(ctx, err.Error())
		return
	}

	if req.Page < 1 {
		req.Page = 1
	}
	if req.PageSize < 1 {
		req.PageSize = 20
	}

	pets, total, err := c.service.QueryPets(ctx.Request.Context(), &req)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.SuccessWithPagination(ctx, pets, req.Page, req.PageSize, total)
}

func (c *PetController) SearchPets(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	if keyword == "" {
		response.ParamError(ctx, "搜索关键词不能为空")
		return
	}

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	pets, total, err := c.service.SearchPets(ctx.Request.Context(), keyword, page, pageSize)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.SuccessWithPagination(ctx, pets, page, pageSize, total)
}

func (c *PetController) GetMyPets(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Unauthorized(ctx, "未登录")
		return
	}

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	pets, total, err := c.service.GetMyPets(ctx.Request.Context(), userID.(int64), page, pageSize)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.SuccessWithPagination(ctx, pets, page, pageSize, total)
}

func (c *PetController) GetRecommendedPets(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	pets, err := c.service.GetRecommendedPets(ctx.Request.Context(), limit)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.Success(ctx, pets)
}

func (c *PetController) GetPendingPets(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	pets, total, err := c.service.GetPendingPets(ctx.Request.Context(), page, pageSize)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.SuccessWithPagination(ctx, pets, page, pageSize, total)
}

func (c *PetController) ApprovePet(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.ParamError(ctx, "无效的宠物ID")
		return
	}

	userID, _ := ctx.Get("user_id")
	if err := c.service.ApprovePet(ctx.Request.Context(), int64(id), userID.(int64)); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, nil)
}

func (c *PetController) RejectPet(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.ParamError(ctx, "无效的宠物ID")
		return
	}

	var req struct {
		Reason string `json:"reason" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ParamError(ctx, "拒绝原因不能为空")
		return
	}

	userID, _ := ctx.Get("user_id")
	if err := c.service.RejectPet(ctx.Request.Context(), int64(id), userID.(int64), req.Reason); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, nil)
}

func (c *PetController) OfflinePet(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.ParamError(ctx, "无效的宠物ID")
		return
	}

	userID, _ := ctx.Get("user_id")
	role, _ := ctx.Get("role")
	isAdmin := role == string(model.RoleAdmin)

	if err := c.service.OfflinePet(ctx.Request.Context(), int64(id), userID.(int64), isAdmin); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, nil)
}

func (c *PetController) GetStatistics(ctx *gin.Context) {
	stats, err := c.service.GetStatistics(ctx.Request.Context())
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.Success(ctx, stats)
}
