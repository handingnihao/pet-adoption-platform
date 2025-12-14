package controller

import (
	"pet-adoption-platform/internal/model"
	"pet-adoption-platform/internal/service"
	"pet-adoption-platform/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdoptionController struct {
	service *service.AdoptionService
}

func NewAdoptionController(db *gorm.DB) *AdoptionController {
	return &AdoptionController{
		service: service.NewAdoptionService(db),
	}
}

// ========== Application接口 ==========

// CreateApplication 创建领养申请
func (c *AdoptionController) CreateApplication(ctx *gin.Context) {
	var req model.ApplicationCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ParamError(ctx, err.Error())
		return
	}

	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Unauthorized(ctx, "未登录")
		return
	}

	app, err := c.service.CreateApplication(ctx.Request.Context(), userID.(int64), &req)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, gin.H{
		"id":             app.ID,
		"application_no": app.ApplicationNo,
		"message":        "申请已提交，等待审核",
	})
}

// GetApplicationByID 获取申请详情
func (c *AdoptionController) GetApplicationByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.ParamError(ctx, "无效的申请ID")
		return
	}

	userID, _ := ctx.Get("user_id")
	role, _ := ctx.Get("role")
	isAdmin := role == string(model.RoleAdmin)

	app, err := c.service.GetApplicationByID(ctx.Request.Context(), id, userID.(int64), isAdmin)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, app)
}

// UpdateApplication 更新申请
func (c *AdoptionController) UpdateApplication(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.ParamError(ctx, "无效的申请ID")
		return
	}

	var req model.ApplicationUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ParamError(ctx, err.Error())
		return
	}

	userID, _ := ctx.Get("user_id")
	if err := c.service.UpdateApplication(ctx.Request.Context(), id, userID.(int64), &req); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, nil)
}

// CancelApplication 取消申请
func (c *AdoptionController) CancelApplication(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.ParamError(ctx, "无效的申请ID")
		return
	}

	userID, _ := ctx.Get("user_id")
	if err := c.service.CancelApplication(ctx.Request.Context(), id, userID.(int64)); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, nil)
}

// ReviewApplication 审核申请（管理员）
func (c *AdoptionController) ReviewApplication(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.ParamError(ctx, "无效的申请ID")
		return
	}

	var req model.ApplicationReviewRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ParamError(ctx, err.Error())
		return
	}

	reviewerID, _ := ctx.Get("user_id")
	if err := c.service.ReviewApplication(ctx.Request.Context(), id, reviewerID.(int64), &req); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, nil)
}

// GetMyApplications 获取我的申请列表
func (c *AdoptionController) GetMyApplications(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Unauthorized(ctx, "未登录")
		return
	}

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	apps, total, err := c.service.GetMyApplications(ctx.Request.Context(), userID.(int64), page, pageSize)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.SuccessWithPagination(ctx, apps, page, pageSize, total)
}

// ListApplications 管理员查看所有申请
func (c *AdoptionController) ListApplications(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	apps, total, err := c.service.ListApplications(ctx.Request.Context(), page, pageSize)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.SuccessWithPagination(ctx, apps, page, pageSize, total)
}

// GetPendingApplications 获取待审核申请
func (c *AdoptionController) GetPendingApplications(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	apps, total, err := c.service.GetPendingApplications(ctx.Request.Context(), page, pageSize)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.SuccessWithPagination(ctx, apps, page, pageSize, total)
}

// QueryApplications 条件查询申请
func (c *AdoptionController) QueryApplications(ctx *gin.Context) {
	var query model.ApplicationQueryRequest
	if err := ctx.ShouldBindQuery(&query); err != nil {
		response.ParamError(ctx, err.Error())
		return
	}

	// 设置默认值
	if query.Page == 0 {
		query.Page = 1
	}
	if query.PageSize == 0 {
		query.PageSize = 20
	}

	apps, total, err := c.service.QueryApplications(ctx.Request.Context(), &query)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.SuccessWithPagination(ctx, apps, query.Page, query.PageSize, total)
}

// ========== Adoption接口 ==========

// CreateAdoption 创建领养记录（管理员）
func (c *AdoptionController) CreateAdoption(ctx *gin.Context) {
	var req model.AdoptionCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ParamError(ctx, err.Error())
		return
	}

	adoption, err := c.service.CreateAdoption(ctx.Request.Context(), &req)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, gin.H{
		"id":      adoption.ID,
		"message": "领养记录创建成功",
	})
}

// GetAdoptionByID 获取领养记录详情
func (c *AdoptionController) GetAdoptionByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.ParamError(ctx, "无效的记录ID")
		return
	}

	userID, _ := ctx.Get("user_id")
	role, _ := ctx.Get("role")
	isAdmin := role == string(model.RoleAdmin)

	adoption, err := c.service.GetAdoptionByID(ctx.Request.Context(), id, userID.(int64), isAdmin)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, adoption)
}

// GetMyAdoptions 获取我的领养记录
func (c *AdoptionController) GetMyAdoptions(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Unauthorized(ctx, "未登录")
		return
	}

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	adoptions, total, err := c.service.GetMyAdoptions(ctx.Request.Context(), userID.(int64), page, pageSize)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.SuccessWithPagination(ctx, adoptions, page, pageSize, total)
}

// ListAdoptions 管理员查看所有领养记录
func (c *AdoptionController) ListAdoptions(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	adoptions, total, err := c.service.ListAdoptions(ctx.Request.Context(), page, pageSize)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.SuccessWithPagination(ctx, adoptions, page, pageSize, total)
}

// UpdateAdoptionStatus 更新领养状态（管理员）
func (c *AdoptionController) UpdateAdoptionStatus(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.ParamError(ctx, "无效的记录ID")
		return
	}

	var req struct {
		Status model.AdoptionStatus `json:"status" binding:"required,oneof=active returned deceased"`
		Notes  string               `json:"notes" binding:"max=1000"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ParamError(ctx, err.Error())
		return
	}

	if err := c.service.UpdateAdoptionStatus(ctx.Request.Context(), id, req.Status, req.Notes); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, nil)
}

// GetStatistics 获取统计信息（管理员）
func (c *AdoptionController) GetStatistics(ctx *gin.Context) {
	stats, err := c.service.GetStatistics(ctx.Request.Context())
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.Success(ctx, stats)
}
