package controller

import (
	"pet-adoption-platform/internal/model"
	"pet-adoption-platform/internal/service"
	"pet-adoption-platform/pkg/logger"
	"pet-adoption-platform/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// UserController 用户控制器
type UserController struct {
	userService *service.UserService
}

// NewUserController 创建用户控制器
func NewUserController(db *gorm.DB, rdb *redis.Client) *UserController {
	return &UserController{
		userService: service.NewUserService(db, rdb),
	}
}

// Register 用户注册
// @Router /api/v1/users/register [post]
func (ctrl *UserController) Register(c *gin.Context) {
	var req model.UserRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("注册参数验证失败", zap.Error(err))
		response.ParamError(c, err.Error())
		return
	}

	user, err := ctrl.userService.Register(&req)
	if err != nil {
		switch err {
		case service.ErrUserExists:
			response.BadRequest(c, "用户名已存在")
		case service.ErrPhoneExists:
			response.BadRequest(c, "手机号已被注册")
		case service.ErrEmailExists:
			response.BadRequest(c, "邮箱已被注册")
		default:
			logger.Error("用户注册失败", zap.Error(err))
			response.Error(c, "注册失败")
		}
		return
	}

	response.Success(c, gin.H{
		"user_id": user.ID,
		"message": "注册成功",
	})
}

// Login 用户登录
// @Router /api/v1/users/login [post]
func (ctrl *UserController) Login(c *gin.Context) {
	var req model.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("登录参数验证失败", zap.Error(err))
		response.ParamError(c, err.Error())
		return
	}

	// 获取客户端IP
	ip := c.ClientIP()

	loginResp, err := ctrl.userService.Login(&req, ip)
	if err != nil {
		switch err {
		case service.ErrUserNotFound:
			response.BadRequest(c, "用户不存在")
		case service.ErrInvalidPassword:
			response.BadRequest(c, "密码错误")
		case service.ErrUserDisabled:
			response.Forbidden(c, "用户已被禁用")
		default:
			logger.Error("用户登录失败", zap.Error(err))
			response.Error(c, "登录失败")
		}
		return
	}

	response.Success(c, loginResp)
}

// GetProfile 获取个人信息
// @Router /api/v1/users/profile [get]
func (ctrl *UserController) GetProfile(c *gin.Context) {
	// 从上下文获取用户ID（由认证中间件设置）
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "请先登录")
		return
	}

	userInfo, err := ctrl.userService.GetProfile(userID.(int64))
	if err != nil {
		if err == service.ErrUserNotFound {
			response.NotFound(c, "用户不存在")
		} else {
			logger.Error("获取用户信息失败", zap.Error(err))
			response.Error(c, "获取失败")
		}
		return
	}

	response.Success(c, userInfo)
}

// UpdateProfile 更新个人信息
// @Router /api/v1/users/profile [put]
func (ctrl *UserController) UpdateProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "请先登录")
		return
	}

	var req model.UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	if err := ctrl.userService.UpdateProfile(userID.(int64), &req); err != nil {
		logger.Error("更新用户信息失败", zap.Error(err))
		response.Error(c, "更新失败")
		return
	}

	response.Success(c, gin.H{
		"message": "更新成功",
	})
}

// UpdatePassword 修改密码
// @Router /api/v1/users/password [put]
func (ctrl *UserController) UpdatePassword(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "请先登录")
		return
	}

	var req model.UpdatePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	if err := ctrl.userService.UpdatePassword(userID.(int64), &req); err != nil {
		switch err {
		case service.ErrOldPasswordWrong:
			response.BadRequest(c, "原密码错误")
		default:
			logger.Error("修改密码失败", zap.Error(err))
			response.Error(c, "修改失败")
		}
		return
	}

	response.Success(c, gin.H{
		"message": "密码修改成功，请重新登录",
	})
}

// GetUserByID 根据ID获取用户信息（管理员功能）
// @Router /api/v1/users/:id [get]
func (ctrl *UserController) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ParamError(c, "用户ID格式错误")
		return
	}

	userInfo, err := ctrl.userService.GetProfile(id)
	if err != nil {
		if err == service.ErrUserNotFound {
			response.NotFound(c, "用户不存在")
		} else {
			logger.Error("获取用户信息失败", zap.Error(err))
			response.Error(c, "获取失败")
		}
		return
	}

	response.Success(c, userInfo)
}

// GetUserList 获取用户列表（管理员功能）
// @Router /api/v1/users [get]
func (ctrl *UserController) GetUserList(c *gin.Context) {
	// 解析分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	// 角色过滤
	roleStr := c.Query("role")
	var role model.UserRole
	if roleStr != "" {
		role = model.UserRole(roleStr)
	}

	// 状态过滤
	statusStr := c.Query("status")
	var status *model.UserStatus
	if statusStr != "" {
		s, _ := strconv.Atoi(statusStr)
		st := model.UserStatus(s)
		status = &st
	}

	users, total, err := ctrl.userService.GetUserList(page, pageSize, role, status)
	if err != nil {
		logger.Error("获取用户列表失败", zap.Error(err))
		response.Error(c, "获取失败")
		return
	}

	response.SuccessWithPagination(c, users, page, pageSize, total)
}

// SearchUsers 搜索用户（管理员功能）
// @Router /api/v1/users/search [get]
func (ctrl *UserController) SearchUsers(c *gin.Context) {
	keyword := c.Query("keyword")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	users, total, err := ctrl.userService.SearchUsers(keyword, page, pageSize)
	if err != nil {
		logger.Error("搜索用户失败", zap.Error(err))
		response.Error(c, "搜索失败")
		return
	}

	response.SuccessWithPagination(c, users, page, pageSize, total)
}

// DisableUser 禁用用户（管理员功能）
// @Router /api/v1/users/:id/disable [put]
func (ctrl *UserController) DisableUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ParamError(c, "用户ID格式错误")
		return
	}

	if err := ctrl.userService.DisableUser(id); err != nil {
		logger.Error("禁用用户失败", zap.Error(err))
		response.Error(c, "操作失败")
		return
	}

	response.Success(c, gin.H{
		"message": "用户已禁用",
	})
}

// EnableUser 启用用户（管理员功能）
// @Router /api/v1/users/:id/enable [put]
func (ctrl *UserController) EnableUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ParamError(c, "用户ID格式错误")
		return
	}

	if err := ctrl.userService.EnableUser(id); err != nil {
		logger.Error("启用用户失败", zap.Error(err))
		response.Error(c, "操作失败")
		return
	}

	response.Success(c, gin.H{
		"message": "用户已启用",
	})
}
