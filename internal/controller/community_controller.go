package controller

import (
	"pet-adoption-platform/internal/model"
	"pet-adoption-platform/internal/service"
	"pet-adoption-platform/pkg/logger"
	"pet-adoption-platform/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// CommunityController 社区控制器
type CommunityController struct {
	communityService *service.CommunityService
}

// NewCommunityController 创建社区控制器
func NewCommunityController(db *gorm.DB) *CommunityController {
	return &CommunityController{
		communityService: service.NewCommunityService(db),
	}
}

// ==================== 动态相关 ====================

// CreatePost 创建动态
// @Summary 创建动态
// @Description 发布新的社区动态
// @Tags 社区
// @Accept json
// @Produce json
// @Param post body model.PostCreateRequest true "动态信息"
// @Success 200 {object} response.Response{data=model.PostInfo}
// @Router /api/v1/community/posts [post]
func (ctrl *CommunityController) CreatePost(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "请先登录")
		return
	}

	var req model.PostCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("创建动态参数错误", zap.Error(err))
		response.ParamError(c, err.Error())
		return
	}

	post, err := ctrl.communityService.CreatePost(c.Request.Context(), &req, uint64(userID.(int64)))
	if err != nil {
		logger.Error("创建动态失败", zap.Error(err))
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, post.ToPostInfo())
}

// GetPost 获取动态详情
// @Summary 获取动态详情
// @Description 根据ID获取动态详情
// @Tags 社区
// @Produce json
// @Param id path int true "动态ID"
// @Success 200 {object} response.Response{data=model.PostInfo}
// @Router /api/v1/community/posts/{id} [get]
func (ctrl *CommunityController) GetPost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ParamError(c, "动态ID格式错误")
		return
	}

	// 获取当前用户ID（可选）
	var currentUserID *uint64
	if userID, exists := c.Get("user_id"); exists {
		uid := uint64(userID.(int64))
		currentUserID = &uid
	}

	postInfo, err := ctrl.communityService.GetPost(c.Request.Context(), id, currentUserID)
	if err != nil {
		logger.Error("获取动态详情失败", zap.Error(err), zap.Uint64("id", id))
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, postInfo)
}

// ListPosts 获取动态列表
// @Summary 获取动态列表
// @Description 获取社区动态列表
// @Tags 社区
// @Produce json
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Param type query string false "类型: story/knowledge/daily/other"
// @Success 200 {object} response.PageResponse{data=[]model.PostInfo}
// @Router /api/v1/community/posts [get]
func (ctrl *CommunityController) ListPosts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	// 类型筛选
	var postType *model.PostType
	if typeStr := c.Query("type"); typeStr != "" {
		t := model.PostType(typeStr)
		postType = &t
	}

	// 获取当前用户ID（可选）
	var currentUserID *uint64
	if userID, exists := c.Get("user_id"); exists {
		uid := uint64(userID.(int64))
		currentUserID = &uid
	}

	posts, total, err := ctrl.communityService.ListPosts(c.Request.Context(), page, pageSize, postType, currentUserID)
	if err != nil {
		logger.Error("获取动态列表失败", zap.Error(err))
		response.Error(c, "获取动态列表失败")
		return
	}

	response.SuccessWithPagination(c, posts, page, pageSize, total)
}

// UpdatePost 更新动态
// @Summary 更新动态
// @Description 更新动态内容
// @Tags 社区
// @Accept json
// @Produce json
// @Param id path int true "动态ID"
// @Param post body model.PostUpdateRequest true "动态信息"
// @Success 200 {object} response.Response
// @Router /api/v1/community/posts/{id} [put]
func (ctrl *CommunityController) UpdatePost(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "请先登录")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ParamError(c, "动态ID格式错误")
		return
	}

	var req model.PostUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("更新动态参数错误", zap.Error(err))
		response.ParamError(c, err.Error())
		return
	}

	if err := ctrl.communityService.UpdatePost(c.Request.Context(), id, &req, uint64(userID.(int64))); err != nil {
		logger.Error("更新动态失败", zap.Error(err), zap.Uint64("id", id))
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, nil)
}

// DeletePost 删除动态
// @Summary 删除动态
// @Description 删除指定的动态
// @Tags 社区
// @Produce json
// @Param id path int true "动态ID"
// @Success 200 {object} response.Response
// @Router /api/v1/community/posts/{id} [delete]
func (ctrl *CommunityController) DeletePost(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "请先登录")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ParamError(c, "动态ID格式错误")
		return
	}

	if err := ctrl.communityService.DeletePost(c.Request.Context(), id, uint64(userID.(int64))); err != nil {
		logger.Error("删除动态失败", zap.Error(err), zap.Uint64("id", id))
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, nil)
}

// SearchPosts 搜索动态
// @Summary 搜索动态
// @Description 根据关键词搜索动态
// @Tags 社区
// @Produce json
// @Param keyword query string true "搜索关键词"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} response.PageResponse{data=[]model.PostInfo}
// @Router /api/v1/community/posts/search [get]
func (ctrl *CommunityController) SearchPosts(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		response.ParamError(c, "请输入搜索关键词")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	// 获取当前用户ID（可选）
	var currentUserID *uint64
	if userID, exists := c.Get("user_id"); exists {
		uid := uint64(userID.(int64))
		currentUserID = &uid
	}

	posts, total, err := ctrl.communityService.SearchPosts(c.Request.Context(), keyword, page, pageSize, currentUserID)
	if err != nil {
		logger.Error("搜索动态失败", zap.Error(err))
		response.Error(c, "搜索动态失败")
		return
	}

	response.SuccessWithPagination(c, posts, page, pageSize, total)
}

// GetMyPosts 获取我的动态
// @Summary 获取我的动态
// @Description 获取当前用户发布的动态列表
// @Tags 社区
// @Produce json
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} response.PageResponse{data=[]model.PostInfo}
// @Router /api/v1/community/posts/my [get]
func (ctrl *CommunityController) GetMyPosts(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "请先登录")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	posts, total, err := ctrl.communityService.GetUserPosts(c.Request.Context(), uint64(userID.(int64)), page, pageSize)
	if err != nil {
		logger.Error("获取我的动态失败", zap.Error(err))
		response.Error(c, "获取动态列表失败")
		return
	}

	response.SuccessWithPagination(c, posts, page, pageSize, total)
}

// ==================== 评论相关 ====================

// CreateComment 创建评论
// @Summary 创建评论
// @Description 对动态发表评论
// @Tags 社区
// @Accept json
// @Produce json
// @Param comment body model.CommentCreateRequest true "评论信息"
// @Success 200 {object} response.Response{data=model.CommentInfo}
// @Router /api/v1/community/comments [post]
func (ctrl *CommunityController) CreateComment(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "请先登录")
		return
	}

	var req model.CommentCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("创建评论参数错误", zap.Error(err))
		response.ParamError(c, err.Error())
		return
	}

	comment, err := ctrl.communityService.CreateComment(c.Request.Context(), &req, uint64(userID.(int64)))
	if err != nil {
		logger.Error("创建评论失败", zap.Error(err))
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, comment.ToCommentInfo())
}

// GetPostComments 获取动态评论
// @Summary 获取动态评论
// @Description 获取动态的评论列表
// @Tags 社区
// @Produce json
// @Param id path int true "动态ID"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} response.PageResponse{data=[]model.CommentInfo}
// @Router /api/v1/community/posts/{id}/comments [get]
func (ctrl *CommunityController) GetPostComments(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ParamError(c, "动态ID格式错误")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	// 获取当前用户ID（可选）
	var currentUserID *uint64
	if userID, exists := c.Get("user_id"); exists {
		uid := uint64(userID.(int64))
		currentUserID = &uid
	}

	comments, total, err := ctrl.communityService.GetPostComments(c.Request.Context(), postID, page, pageSize, currentUserID)
	if err != nil {
		logger.Error("获取评论列表失败", zap.Error(err))
		response.Error(c, "获取评论列表失败")
		return
	}

	response.SuccessWithPagination(c, comments, page, pageSize, total)
}

// DeleteComment 删除评论
// @Summary 删除评论
// @Description 删除指定的评论
// @Tags 社区
// @Produce json
// @Param id path int true "评论ID"
// @Success 200 {object} response.Response
// @Router /api/v1/community/comments/{id} [delete]
func (ctrl *CommunityController) DeleteComment(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "请先登录")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ParamError(c, "评论ID格式错误")
		return
	}

	if err := ctrl.communityService.DeleteComment(c.Request.Context(), id, uint64(userID.(int64))); err != nil {
		logger.Error("删除评论失败", zap.Error(err), zap.Uint64("id", id))
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, nil)
}

// ==================== 点赞相关 ====================

// LikePost 点赞动态
// @Summary 点赞动态
// @Description 对动态点赞
// @Tags 社区
// @Produce json
// @Param id path int true "动态ID"
// @Success 200 {object} response.Response
// @Router /api/v1/community/posts/{id}/like [post]
func (ctrl *CommunityController) LikePost(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "请先登录")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ParamError(c, "动态ID格式错误")
		return
	}

	if err := ctrl.communityService.LikePost(c.Request.Context(), id, uint64(userID.(int64))); err != nil {
		logger.Error("点赞失败", zap.Error(err), zap.Uint64("id", id))
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, nil)
}

// UnlikePost 取消点赞动态
// @Summary 取消点赞动态
// @Description 取消对动态的点赞
// @Tags 社区
// @Produce json
// @Param id path int true "动态ID"
// @Success 200 {object} response.Response
// @Router /api/v1/community/posts/{id}/like [delete]
func (ctrl *CommunityController) UnlikePost(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "请先登录")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ParamError(c, "动态ID格式错误")
		return
	}

	if err := ctrl.communityService.UnlikePost(c.Request.Context(), id, uint64(userID.(int64))); err != nil {
		logger.Error("取消点赞失败", zap.Error(err), zap.Uint64("id", id))
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, nil)
}

// LikeComment 点赞评论
// @Summary 点赞评论
// @Description 对评论点赞
// @Tags 社区
// @Produce json
// @Param id path int true "评论ID"
// @Success 200 {object} response.Response
// @Router /api/v1/community/comments/{id}/like [post]
func (ctrl *CommunityController) LikeComment(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "请先登录")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ParamError(c, "评论ID格式错误")
		return
	}

	if err := ctrl.communityService.LikeComment(c.Request.Context(), id, uint64(userID.(int64))); err != nil {
		logger.Error("点赞评论失败", zap.Error(err), zap.Uint64("id", id))
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, nil)
}

// UnlikeComment 取消点赞评论
// @Summary 取消点赞评论
// @Description 取消对评论的点赞
// @Tags 社区
// @Produce json
// @Param id path int true "评论ID"
// @Success 200 {object} response.Response
// @Router /api/v1/community/comments/{id}/like [delete]
func (ctrl *CommunityController) UnlikeComment(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "请先登录")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ParamError(c, "评论ID格式错误")
		return
	}

	if err := ctrl.communityService.UnlikeComment(c.Request.Context(), id, uint64(userID.(int64))); err != nil {
		logger.Error("取消点赞评论失败", zap.Error(err), zap.Uint64("id", id))
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, nil)
}
