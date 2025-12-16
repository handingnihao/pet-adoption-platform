package controller

import (
	"pet-adoption-platform/internal/dao"
	"pet-adoption-platform/internal/model"
	"pet-adoption-platform/internal/service"
	"pet-adoption-platform/pkg/logger"
	"pet-adoption-platform/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// DonationController 捐赠控制器
type DonationController struct {
	donationService *service.DonationService
}

// NewDonationController 创建DonationController实例
func NewDonationController(db *gorm.DB) *DonationController {
	donationDAO := dao.NewDonationDAO(db)
	orgDAO := dao.NewOrganizationDAO(db)
	donationService := service.NewDonationService(donationDAO, orgDAO)
	return &DonationController{donationService: donationService}
}

// CreateDonation 创建捐赠
func (c *DonationController) CreateDonation(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Unauthorized(ctx, "请先登录")
		return
	}

	var req model.DonationCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Warn("创建捐赠参数错误", zap.Error(err))
		response.ParamError(ctx, err.Error())
		return
	}

	donation, err := c.donationService.CreateDonation(ctx.Request.Context(), &req, uint64(userID.(int64)))
	if err != nil {
		logger.Error("创建捐赠失败", zap.Error(err))
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, donation.ToInfo())
}

// GetDonation 获取捐赠详情
func (c *DonationController) GetDonation(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.ParamError(ctx, "无效的捐赠ID")
		return
	}

	donation, err := c.donationService.GetDonation(ctx.Request.Context(), id)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, donation)
}

// ListDonations 获取捐赠列表
func (c *DonationController) ListDonations(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	var donationType *model.DonationType
	if typeStr := ctx.Query("type"); typeStr != "" {
		t := model.DonationType(typeStr)
		donationType = &t
	}

	var status *model.DonationStatus
	if statusStr := ctx.Query("status"); statusStr != "" {
		s, _ := strconv.Atoi(statusStr)
		st := model.DonationStatus(s)
		status = &st
	}

	donations, total, err := c.donationService.ListDonations(ctx.Request.Context(), page, pageSize, donationType, status)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.SuccessWithPagination(ctx, donations, page, pageSize, total)
}

// GetMyDonations 获取我的捐赠记录
func (c *DonationController) GetMyDonations(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Unauthorized(ctx, "请先登录")
		return
	}

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	donations, total, err := c.donationService.GetMyDonations(ctx.Request.Context(), uint64(userID.(int64)), page, pageSize)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.SuccessWithPagination(ctx, donations, page, pageSize, total)
}

// ConfirmDonation 确认捐赠
func (c *DonationController) ConfirmDonation(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Unauthorized(ctx, "请先登录")
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.ParamError(ctx, "无效的捐赠ID")
		return
	}

	var req model.DonationConfirmRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ParamError(ctx, err.Error())
		return
	}

	if err := c.donationService.ConfirmDonation(ctx.Request.Context(), id, uint64(userID.(int64)), req.TransactionID); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, nil)
}

// CancelDonation 取消捐赠
func (c *DonationController) CancelDonation(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Unauthorized(ctx, "请先登录")
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.ParamError(ctx, "无效的捐赠ID")
		return
	}

	if err := c.donationService.CancelDonation(ctx.Request.Context(), id, uint64(userID.(int64))); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, nil)
}

// CompleteDonation 完成捐赠（管理员）
func (c *DonationController) CompleteDonation(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Unauthorized(ctx, "请先登录")
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.ParamError(ctx, "无效的捐赠ID")
		return
	}

	if err := c.donationService.CompleteDonation(ctx.Request.Context(), id, uint64(userID.(int64))); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, nil)
}

// GetStatistics 获取捐赠统计
func (c *DonationController) GetStatistics(ctx *gin.Context) {
	stats, err := c.donationService.GetStatistics(ctx.Request.Context())
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.Success(ctx, stats)
}

// GetPublicDonations 获取公开捐赠列表（爱心墙）
func (c *DonationController) GetPublicDonations(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))
	if limit > 100 {
		limit = 100
	}

	donations, err := c.donationService.GetPublicDonations(ctx.Request.Context(), limit)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}

	response.Success(ctx, donations)
}
