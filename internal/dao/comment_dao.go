package dao

import (
	"context"
	"errors"
	"pet-adoption-platform/internal/model"
	"pet-adoption-platform/pkg/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// CommentDAO 评论数据访问对象
type CommentDAO struct {
	db *gorm.DB
}

// NewCommentDAO 创建评论DAO实例
func NewCommentDAO(db *gorm.DB) *CommentDAO {
	return &CommentDAO{db: db}
}

// Create 创建评论
func (dao *CommentDAO) Create(ctx context.Context, comment *model.Comment) error {
	if comment == nil {
		return errors.New("comment is nil")
	}

	if err := dao.db.WithContext(ctx).Create(comment).Error; err != nil {
		logger.Error("创建评论失败", zap.Error(err))
		return err
	}

	return nil
}

// GetByID 根据ID获取评论
func (dao *CommentDAO) GetByID(ctx context.Context, id uint64) (*model.Comment, error) {
	if id == 0 {
		return nil, errors.New("invalid comment id")
	}

	var comment model.Comment
	if err := dao.db.WithContext(ctx).
		Preload("User").
		Where("deleted_at IS NULL").
		First(&comment, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.Error("查询评论失败", zap.Error(err), zap.Uint64("id", id))
		return nil, err
	}

	return &comment, nil
}

// ListByPostID 获取动态的评论列表（只获取一级评论）
func (dao *CommentDAO) ListByPostID(ctx context.Context, postID uint64, page, pageSize int) ([]*model.Comment, int64, error) {
	var (
		comments []*model.Comment
		total    int64
	)

	db := dao.db.WithContext(ctx).Model(&model.Comment{}).
		Where("post_id = ? AND parent_id = 0 AND status = ? AND deleted_at IS NULL",
			postID, model.CommentStatusVisible)

	// 计算总数
	if err := db.Count(&total).Error; err != nil {
		logger.Error("获取评论总数失败", zap.Error(err))
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := db.Preload("User").
		Offset(offset).Limit(pageSize).
		Order("created_at DESC").
		Find(&comments).Error; err != nil {
		logger.Error("获取评论列表失败", zap.Error(err))
		return nil, 0, err
	}

	// 获取子评论
	for _, comment := range comments {
		replies, err := dao.GetReplies(ctx, comment.ID)
		if err != nil {
			logger.Warn("获取子评论失败", zap.Error(err), zap.Uint64("parent_id", comment.ID))
			continue
		}
		comment.Replies = replies
	}

	return comments, total, nil
}

// GetReplies 获取评论的回复列表
func (dao *CommentDAO) GetReplies(ctx context.Context, parentID uint64) ([]*model.Comment, error) {
	var replies []*model.Comment

	if err := dao.db.WithContext(ctx).
		Preload("User").
		Preload("ReplyToUser").
		Where("parent_id = ? AND status = ? AND deleted_at IS NULL",
			parentID, model.CommentStatusVisible).
		Order("created_at ASC").
		Limit(10). // 限制子评论数量
		Find(&replies).Error; err != nil {
		logger.Error("获取回复列表失败", zap.Error(err))
		return nil, err
	}

	return replies, nil
}

// Delete 删除评论（软删除）
func (dao *CommentDAO) Delete(ctx context.Context, id uint64) error {
	if id == 0 {
		return errors.New("invalid comment id")
	}

	result := dao.db.WithContext(ctx).
		Model(&model.Comment{}).
		Where("id = ?", id).
		Update("deleted_at", gorm.Expr("CURRENT_TIMESTAMP"))

	if result.Error != nil {
		logger.Error("删除评论失败", zap.Error(result.Error), zap.Uint64("id", id))
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// IncrementLikeCount 增加点赞数
func (dao *CommentDAO) IncrementLikeCount(ctx context.Context, id uint64, delta int) error {
	return dao.db.WithContext(ctx).Model(&model.Comment{}).
		Where("id = ?", id).
		UpdateColumn("like_count", gorm.Expr("like_count + ?", delta)).Error
}

// CountByPostID 统计动态的评论数
func (dao *CommentDAO) CountByPostID(ctx context.Context, postID uint64) (int64, error) {
	var count int64
	err := dao.db.WithContext(ctx).Model(&model.Comment{}).
		Where("post_id = ? AND status = ? AND deleted_at IS NULL",
			postID, model.CommentStatusVisible).
		Count(&count).Error
	return count, err
}
