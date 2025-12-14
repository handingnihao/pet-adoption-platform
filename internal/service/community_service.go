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

// CommunityService 社区服务
type CommunityService struct {
	postDAO    *dao.PostDAO
	commentDAO *dao.CommentDAO
	likeDAO    *dao.LikeDAO
	db         *gorm.DB
}

// NewCommunityService 创建社区服务实例
func NewCommunityService(db *gorm.DB) *CommunityService {
	return &CommunityService{
		postDAO:    dao.NewPostDAO(db),
		commentDAO: dao.NewCommentDAO(db),
		likeDAO:    dao.NewLikeDAO(db),
		db:         db,
	}
}

// ==================== 动态相关 ====================

// CreatePost 创建动态
func (s *CommunityService) CreatePost(ctx context.Context, req *model.PostCreateRequest, userID uint64) (*model.Post, error) {
	post := &model.Post{
		UserID:   userID,
		Title:    req.Title,
		Content:  req.Content,
		VideoURL: req.VideoURL,
		PetID:    req.PetID,
		Type:     req.Type,
		Status:   model.PostStatusVisible,
	}

	// 处理图片
	if len(req.Images) > 0 {
		post.Images = model.ToJSONString(req.Images)
	}

	// 处理话题
	if len(req.TopicIDs) > 0 {
		post.TopicIDs = model.ToJSONInt64String(req.TopicIDs)
	}

	// 默认类型
	if post.Type == "" {
		post.Type = model.PostTypeDaily
	}

	if err := s.postDAO.Create(ctx, post); err != nil {
		logger.Error("创建动态失败", zap.Error(err))
		return nil, err
	}

	return post, nil
}

// UpdatePost 更新动态
func (s *CommunityService) UpdatePost(ctx context.Context, id uint64, req *model.PostUpdateRequest, userID uint64) error {
	// 检查动态是否存在
	post, err := s.postDAO.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if post == nil {
		return errors.New("动态不存在")
	}

	// 检查权限
	if post.UserID != userID {
		return errors.New("无权修改此动态")
	}

	// 构建更新字段
	updates := make(map[string]interface{})
	if req.Title != nil {
		updates["title"] = *req.Title
	}
	if req.Content != nil {
		updates["content"] = *req.Content
	}
	if req.Images != nil {
		updates["images"] = model.ToJSONString(*req.Images)
	}
	if req.VideoURL != nil {
		updates["video_url"] = *req.VideoURL
	}
	if req.TopicIDs != nil {
		updates["topic_ids"] = model.ToJSONInt64String(*req.TopicIDs)
	}
	if req.PetID != nil {
		updates["pet_id"] = *req.PetID
	}
	if req.Type != nil {
		updates["type"] = *req.Type
	}

	return s.postDAO.Update(ctx, id, updates)
}

// GetPost 获取动态详情
func (s *CommunityService) GetPost(ctx context.Context, id uint64, currentUserID *uint64) (*model.PostInfo, error) {
	post, err := s.postDAO.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, errors.New("动态不存在")
	}

	// 增加浏览量
	_ = s.postDAO.IncrementViewCount(ctx, id)

	info := post.ToPostInfo()

	// 检查当前用户是否点赞
	if currentUserID != nil && *currentUserID > 0 {
		liked, _ := s.likeDAO.Exists(ctx, *currentUserID, "post", id)
		info.IsLiked = liked
	}

	return info, nil
}

// ListPosts 获取动态列表
func (s *CommunityService) ListPosts(ctx context.Context, page, pageSize int, postType *model.PostType, currentUserID *uint64) ([]*model.PostInfo, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 50 {
		pageSize = 20
	}

	posts, total, err := s.postDAO.List(ctx, page, pageSize, postType, nil)
	if err != nil {
		return nil, 0, err
	}

	// 转换为响应数据
	postInfos := make([]*model.PostInfo, 0, len(posts))
	postIDs := make([]uint64, 0, len(posts))

	for _, post := range posts {
		postInfos = append(postInfos, post.ToPostInfo())
		postIDs = append(postIDs, post.ID)
	}

	// 批量检查当前用户的点赞状态
	if currentUserID != nil && *currentUserID > 0 && len(postIDs) > 0 {
		likedMap, _ := s.likeDAO.GetUserLikedTargets(ctx, *currentUserID, "post", postIDs)
		for _, info := range postInfos {
			info.IsLiked = likedMap[info.ID]
		}
	}

	return postInfos, total, nil
}

// DeletePost 删除动态
func (s *CommunityService) DeletePost(ctx context.Context, id uint64, userID uint64) error {
	post, err := s.postDAO.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if post == nil {
		return errors.New("动态不存在")
	}

	// 检查权限
	if post.UserID != userID {
		return errors.New("无权删除此动态")
	}

	return s.postDAO.Delete(ctx, id)
}

// SearchPosts 搜索动态
func (s *CommunityService) SearchPosts(ctx context.Context, keyword string, page, pageSize int, currentUserID *uint64) ([]*model.PostInfo, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 50 {
		pageSize = 20
	}

	posts, total, err := s.postDAO.Search(ctx, keyword, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	postInfos := make([]*model.PostInfo, 0, len(posts))
	postIDs := make([]uint64, 0, len(posts))

	for _, post := range posts {
		postInfos = append(postInfos, post.ToPostInfo())
		postIDs = append(postIDs, post.ID)
	}

	// 批量检查点赞状态
	if currentUserID != nil && *currentUserID > 0 && len(postIDs) > 0 {
		likedMap, _ := s.likeDAO.GetUserLikedTargets(ctx, *currentUserID, "post", postIDs)
		for _, info := range postInfos {
			info.IsLiked = likedMap[info.ID]
		}
	}

	return postInfos, total, nil
}

// GetUserPosts 获取用户的动态
func (s *CommunityService) GetUserPosts(ctx context.Context, userID uint64, page, pageSize int) ([]*model.PostInfo, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 50 {
		pageSize = 20
	}

	posts, total, err := s.postDAO.GetUserPosts(ctx, userID, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	postInfos := make([]*model.PostInfo, 0, len(posts))
	for _, post := range posts {
		postInfos = append(postInfos, post.ToPostInfo())
	}

	return postInfos, total, nil
}

// ==================== 评论相关 ====================

// CreateComment 创建评论
func (s *CommunityService) CreateComment(ctx context.Context, req *model.CommentCreateRequest, userID uint64) (*model.Comment, error) {
	// 检查动态是否存在
	post, err := s.postDAO.GetByID(ctx, req.PostID)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, errors.New("动态不存在")
	}

	// 如果是回复，检查父评论是否存在
	if req.ParentID > 0 {
		parent, err := s.commentDAO.GetByID(ctx, req.ParentID)
		if err != nil {
			return nil, err
		}
		if parent == nil {
			return nil, errors.New("回复的评论不存在")
		}
	}

	comment := &model.Comment{
		PostID:        req.PostID,
		UserID:        userID,
		ParentID:      req.ParentID,
		ReplyToUserID: req.ReplyToUserID,
		Content:       req.Content,
		Status:        model.CommentStatusVisible,
	}

	if err := s.commentDAO.Create(ctx, comment); err != nil {
		logger.Error("创建评论失败", zap.Error(err))
		return nil, err
	}

	// 更新动态评论数
	_ = s.postDAO.IncrementCommentCount(ctx, req.PostID, 1)

	return comment, nil
}

// GetPostComments 获取动态的评论列表
func (s *CommunityService) GetPostComments(ctx context.Context, postID uint64, page, pageSize int, currentUserID *uint64) ([]*model.CommentInfo, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 50 {
		pageSize = 20
	}

	comments, total, err := s.commentDAO.ListByPostID(ctx, postID, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	commentInfos := make([]*model.CommentInfo, 0, len(comments))
	commentIDs := make([]uint64, 0, len(comments))

	for _, comment := range comments {
		commentInfos = append(commentInfos, comment.ToCommentInfo())
		commentIDs = append(commentIDs, comment.ID)
	}

	// 批量检查点赞状态
	if currentUserID != nil && *currentUserID > 0 && len(commentIDs) > 0 {
		likedMap, _ := s.likeDAO.GetUserLikedTargets(ctx, *currentUserID, "comment", commentIDs)
		for _, info := range commentInfos {
			info.IsLiked = likedMap[info.ID]
		}
	}

	return commentInfos, total, nil
}

// DeleteComment 删除评论
func (s *CommunityService) DeleteComment(ctx context.Context, id uint64, userID uint64) error {
	comment, err := s.commentDAO.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if comment == nil {
		return errors.New("评论不存在")
	}

	// 检查权限
	if comment.UserID != userID {
		return errors.New("无权删除此评论")
	}

	if err := s.commentDAO.Delete(ctx, id); err != nil {
		return err
	}

	// 更新动态评论数
	_ = s.postDAO.IncrementCommentCount(ctx, comment.PostID, -1)

	return nil
}

// ==================== 点赞相关 ====================

// LikePost 点赞动态
func (s *CommunityService) LikePost(ctx context.Context, postID uint64, userID uint64) error {
	// 检查动态是否存在
	post, err := s.postDAO.GetByID(ctx, postID)
	if err != nil {
		return err
	}
	if post == nil {
		return errors.New("动态不存在")
	}

	// 检查是否已点赞
	exists, err := s.likeDAO.Exists(ctx, userID, "post", postID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("已点赞")
	}

	// 创建点赞记录
	like := &model.Like{
		UserID:     userID,
		TargetType: "post",
		TargetID:   postID,
	}

	if err := s.likeDAO.Create(ctx, like); err != nil {
		return err
	}

	// 更新点赞数
	return s.postDAO.IncrementLikeCount(ctx, postID, 1)
}

// UnlikePost 取消点赞动态
func (s *CommunityService) UnlikePost(ctx context.Context, postID uint64, userID uint64) error {
	// 检查是否已点赞
	exists, err := s.likeDAO.Exists(ctx, userID, "post", postID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("未点赞")
	}

	// 删除点赞记录
	if err := s.likeDAO.Delete(ctx, userID, "post", postID); err != nil {
		return err
	}

	// 更新点赞数
	return s.postDAO.IncrementLikeCount(ctx, postID, -1)
}

// LikeComment 点赞评论
func (s *CommunityService) LikeComment(ctx context.Context, commentID uint64, userID uint64) error {
	// 检查评论是否存在
	comment, err := s.commentDAO.GetByID(ctx, commentID)
	if err != nil {
		return err
	}
	if comment == nil {
		return errors.New("评论不存在")
	}

	// 检查是否已点赞
	exists, err := s.likeDAO.Exists(ctx, userID, "comment", commentID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("已点赞")
	}

	// 创建点赞记录
	like := &model.Like{
		UserID:     userID,
		TargetType: "comment",
		TargetID:   commentID,
	}

	if err := s.likeDAO.Create(ctx, like); err != nil {
		return err
	}

	// 更新点赞数
	return s.commentDAO.IncrementLikeCount(ctx, commentID, 1)
}

// UnlikeComment 取消点赞评论
func (s *CommunityService) UnlikeComment(ctx context.Context, commentID uint64, userID uint64) error {
	// 检查是否已点赞
	exists, err := s.likeDAO.Exists(ctx, userID, "comment", commentID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("未点赞")
	}

	// 删除点赞记录
	if err := s.likeDAO.Delete(ctx, userID, "comment", commentID); err != nil {
		return err
	}

	// 更新点赞数
	return s.commentDAO.IncrementLikeCount(ctx, commentID, -1)
}
