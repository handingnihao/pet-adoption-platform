package controller

import (
	"net/http"
	"pet-adoption-platform/internal/model"
	"pet-adoption-platform/internal/service"
	"pet-adoption-platform/pkg/logger"
	"pet-adoption-platform/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// OrganizationController 机构控制器
type OrganizationController struct {
	orgService *service.OrganizationService
}

// NewOrganizationController 创建机构控制器
func NewOrganizationController(db *gorm.DB) *OrganizationController {
	return &OrganizationController{
		orgService: service.NewOrganizationService(db),
	}
}

// CreateOrganization 创建机构
// @Summary 创建机构
// @Description 创建新的机构
// @Tags 机构管理
// @Accept json
// @Produce json
// @Param organization body model.OrganizationCreateRequest true "机构信息"
// @Success 200 {object} response.Response{data=model.OrganizationInfo} "成功返回创建的机构信息"
// @Router /api/v1/organizations [post]
func (ctrl *OrganizationController) CreateOrganization(c *gin.Context) {
	// 获取当前用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "请先登录")
		return
	}

	// 解析请求参数
	var req model.OrganizationCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("创建机构参数错误", zap.Error(err))
		response.ParamError(c, err.Error())
		return
	}

	// 创建机构
	org, err := ctrl.orgService.CreateOrganization(c.Request.Context(), &req, userID.(int64))
	if err != nil {
		logger.Error("创建机构失败", zap.Error(err))
		response.BadRequest(c, err.Error())
		return
	}

	// 返回结果
	response.Success(c, org.ToOrganizationInfo())
}

// UpdateOrganization 更新机构信息
// @Summary 更新机构信息
// @Description 更新机构信息
// @Tags 机构管理
// @Accept json
// @Produce json
// @Param id path int true "机构ID"
// @Param organization body model.OrganizationUpdateRequest true "机构信息"
// @Success 200 {object} response.Response "成功"
// @Router /api/v1/organizations/{id} [put]
func (ctrl *OrganizationController) UpdateOrganization(c *gin.Context) {
	// 获取当前用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "请先登录")
		return
	}

	// 解析机构ID
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.ParamError(c, "机构ID格式错误")
		return
	}

	// 解析请求参数
	var req model.OrganizationUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("更新机构参数错误", zap.Error(err))
		response.ParamError(c, err.Error())
		return
	}

	// 更新机构信息
	if err := ctrl.orgService.UpdateOrganization(c.Request.Context(), id, &req, userID.(int64)); err != nil {
		logger.Error("更新机构失败", zap.Error(err), zap.Int64("org_id", id))
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, nil)
}

// GetOrganization 获取机构详情
// @Summary 获取机构详情
// @Description 根据ID获取机构详情
// @Tags 机构管理
// @Produce json
// @Param id path int true "机构ID"
// @Success 200 {object} response.Response{data=model.OrganizationInfo} "成功返回机构信息"
// @Router /api/v1/organizations/{id} [get]
func (ctrl *OrganizationController) GetOrganization(c *gin.Context) {
	// 解析机构ID
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.ParamError(c, "机构ID格式错误")
		return
	}

	// 获取机构详情
	org, err := ctrl.orgService.GetOrganization(c.Request.Context(), id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.NotFound(c, "机构不存在")
		} else {
			logger.Error("获取机构详情失败", zap.Error(err), zap.Int64("org_id", id))
			response.Error(c, "获取机构详情失败")
		}
		return
	}

	if org == nil {
		response.NotFound(c, "机构不存在")
		return
	}

	response.Success(c, org.ToOrganizationInfo())
}

// ListOrganizations 获取机构列表
// @Summary 获取机构列表
// @Description 获取机构列表，支持分页和状态筛选
// @Tags 机构管理
// @Produce json
// @Param page query int false "页码，默认为1"
// @Param page_size query int false "每页数量，默认为20，最大100"
// @Param status query int false "状态筛选：0-待审核 1-已通过 2-已拒绝"
// @Success 200 {object} response.PageResponse{data=[]model.OrganizationInfo} "成功返回机构列表"
// @Router /api/v1/organizations [get]
func (ctrl *OrganizationController) ListOrganizations(c *gin.Context) {
	// 解析分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	// 解析状态参数
	var status *model.OrganizationStatus
	if statusStr := c.Query("status"); statusStr != "" {
		statusInt, err := strconv.Atoi(statusStr)
		if err == nil {
			s := model.OrganizationStatus(statusInt)
			status = &s
		}
	}

	// 获取机构列表
	orgs, total, err := ctrl.orgService.ListOrganizations(c.Request.Context(), page, pageSize, status)
	if err != nil {
		logger.Error("获取机构列表失败", zap.Error(err))
		response.Error(c, "获取机构列表失败")
		return
	}

	// 转换为响应数据
	orgInfos := make([]*model.OrganizationInfo, 0, len(orgs))
	for _, org := range orgs {
		orgInfos = append(orgInfos, org.ToOrganizationInfo())
	}

	response.SuccessWithPagination(c, orgInfos, page, pageSize, total)
}

// DeleteOrganization 删除机构
// @Summary 删除机构
// @Description 删除指定ID的机构（软删除）
// @Tags 机构管理
// @Produce json
// @Param id path int true "机构ID"
// @Success 200 {object} response.Response "成功"
// @Router /api/v1/organizations/{id} [delete]
func (ctrl *OrganizationController) DeleteOrganization(c *gin.Context) {
	// 解析机构ID
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.ParamError(c, "机构ID格式错误")
		return
	}

	// 删除机构
	if err := ctrl.orgService.DeleteOrganization(c.Request.Context(), id); err != nil {
		if err == gorm.ErrRecordNotFound {
			response.NotFound(c, "机构不存在")
		} else {
			logger.Error("删除机构失败", zap.Error(err), zap.Int64("org_id", id))
			response.Error(c, "删除机构失败")
		}
		return
	}

	response.Success(c, nil)
}

// UpdateOrganizationStatus 更新机构状态
// @Summary 更新机构状态
// @Description 更新机构状态（审核）
// @Tags 机构管理
// @Accept json
// @Produce json
// @Param id path int true "机构ID"
// @Param status body model.OrganizationStatusUpdateRequest true "状态信息"
// @Success 200 {object} response.Response "成功"
// @Router /api/v1/organizations/{id}/status [put]
func (ctrl *OrganizationController) UpdateOrganizationStatus(c *gin.Context) {
	// 获取当前用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "请先登录")
		return
	}

	// 解析机构ID
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.ParamError(c, "机构ID格式错误")
		return
	}

	// 解析请求参数
	var req model.OrganizationStatusUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("更新机构状态参数错误", zap.Error(err))
		response.ParamError(c, err.Error())
		return
	}

	// 更新状态
	if err := ctrl.orgService.UpdateOrganizationStatus(c.Request.Context(), id, &req, userID.(int64)); err != nil {
		if err == gorm.ErrRecordNotFound {
			response.NotFound(c, "机构不存在")
		} else {
			logger.Error("更新机构状态失败", zap.Error(err), zap.Int64("org_id", id))
			response.Error(c, "更新机构状态失败")
		}
		return
	}

	response.Success(c, nil)
}

// GetMyOrganizations 获取我创建的机构列表
// @Summary 获取我创建的机构列表
// @Description 获取当前用户创建的机构列表
// @Tags 机构管理
// @Produce json
// @Param page query int false "页码，默认为1"
// @Param page_size query int false "每页数量，默认为20，最大100"
// @Success 200 {object} response.PageResponse{data=[]model.OrganizationInfo} "成功返回机构列表"
// @Router /api/v1/organizations/my [get]
func (ctrl *OrganizationController) GetMyOrganizations(c *gin.Context) {
	// 获取当前用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "请先登录")
		return
	}

	// 解析分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	// 获取用户创建的机构列表
	orgs, total, err := ctrl.orgService.GetUserOrganizations(c.Request.Context(), userID.(int64), page, pageSize)
	if err != nil {
		logger.Error("获取用户机构列表失败", zap.Error(err), zap.Int64("user_id", userID.(int64)))
		response.Error(c, "获取机构列表失败")
		return
	}

	// 转换为响应数据
	orgInfos := make([]*model.OrganizationInfo, 0, len(orgs))
	for _, org := range orgs {
		orgInfos = append(orgInfos, org.ToOrganizationInfo())
	}

	response.SuccessWithPagination(c, orgInfos, page, pageSize, total)
}
