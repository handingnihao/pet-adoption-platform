package dao

import (
	"context"
	"errors"
	"pet-adoption-platform/internal/model"
	"pet-adoption-platform/pkg/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// PostDAO 动态数据访问对象
type PostDAO struct {
	db *gorm.DB
}

// NewPostDAO 创建动态DAO实例
func NewPostDAO(db *gorm.DB) *PostDAO {
	return &PostDAO{db: db}
}

// Create 创建动态
func (dao *PostDAO) Create(ctx context.Context, post *model.Post) error {
	if post == nil {
		return errors.New("post is nil")
	}

	if err := dao.db.WithContext(ctx).Create(post).Error; err != nil {
		logger.Error("创建动态失败", zap.Error(err))
		return err
	}

	return nil
}

// Update 更新动态
func (dao *PostDAO) Update(ctx context.Context, id uint64, updates map[string]interface{}) error {
	if id == 0 {
		return errors.New("invalid post id")
	}

	result := dao.db.WithContext(ctx).Model(&model.Post{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(updates)

	if result.Error != nil {
		logger.Error("更新动态失败", zap.Error(result.Error), zap.Uint64("id", id))
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// GetByID 根据ID获取动态
func (dao *PostDAO) GetByID(ctx context.Context, id uint64) (*model.Post, error) {
	if id == 0 {
		return nil, errors.New("invalid post id")
	}

	var post model.Post
	if err := dao.db.WithContext(ctx).
		Preload("User").
		Where("deleted_at IS NULL").
		First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.Error("查询动态失败", zap.Error(err), zap.Uint64("id", id))
		return nil, err
	}

	return &post, nil
}

// List 获取动态列表
func (dao *PostDAO) List(ctx context.Context, page, pageSize int, postType *model.PostType, userID *uint64) ([]*model.Post, int64, error) {
	var (
		posts []*model.Post
		total int64
	)

	db := dao.db.WithContext(ctx).Model(&model.Post{}).
		Where("status = ? AND deleted_at IS NULL", model.PostStatusVisible)

	// 类型筛选
	if postType != nil {
		db = db.Where("type = ?", *postType)
	}

	// 用户筛选
	if userID != nil {
		db = db.Where("user_id = ?", *userID)
	}

	// 计算总数
	if err := db.Count(&total).Error; err != nil {
		logger.Error("获取动态总数失败", zap.Error(err))
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := db.Preload("User").
		Offset(offset).Limit(pageSize).
		Order("is_top DESC, created_at DESC").
		Find(&posts).Error; err != nil {
		logger.Error("获取动态列表失败", zap.Error(err))
		return nil, 0, err
	}

	return posts, total, nil
}

// Delete 删除动态（软删除）
func (dao *PostDAO) Delete(ctx context.Context, id uint64) error {
	if id == 0 {
		return errors.New("invalid post id")
	}

	result := dao.db.WithContext(ctx).
		Model(&model.Post{}).
		Where("id = ?", id).
		Update("deleted_at", gorm.Expr("CURRENT_TIMESTAMP"))

	if result.Error != nil {
		logger.Error("删除动态失败", zap.Error(result.Error), zap.Uint64("id", id))
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// IncrementViewCount 增加浏览次数
func (dao *PostDAO) IncrementViewCount(ctx context.Context, id uint64) error {
	return dao.db.WithContext(ctx).Model(&model.Post{}).
		Where("id = ?", id).
		UpdateColumn("view_count", gorm.Expr("view_count + 1")).Error
}

// IncrementLikeCount 增加点赞数
func (dao *PostDAO) IncrementLikeCount(ctx context.Context, id uint64, delta int) error {
	return dao.db.WithContext(ctx).Model(&model.Post{}).
		Where("id = ?", id).
		UpdateColumn("like_count", gorm.Expr("like_count + ?", delta)).Error
}

// IncrementCommentCount 增加评论数
func (dao *PostDAO) IncrementCommentCount(ctx context.Context, id uint64, delta int) error {
	return dao.db.WithContext(ctx).Model(&model.Post{}).
		Where("id = ?", id).
		UpdateColumn("comment_count", gorm.Expr("comment_count + ?", delta)).Error
}

// Search 搜索动态
func (dao *PostDAO) Search(ctx context.Context, keyword string, page, pageSize int) ([]*model.Post, int64, error) {
	var (
		posts []*model.Post
		total int64
	)

	db := dao.db.WithContext(ctx).Model(&model.Post{}).
		Where("status = ? AND deleted_at IS NULL", model.PostStatusVisible).
		Where("title LIKE ? OR content LIKE ?", "%"+keyword+"%", "%"+keyword+"%")

	// 计算总数
	if err := db.Count(&total).Error; err != nil {
		logger.Error("搜索动态总数失败", zap.Error(err))
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := db.Preload("User").
		Offset(offset).Limit(pageSize).
		Order("created_at DESC").
		Find(&posts).Error; err != nil {
		logger.Error("搜索动态失败", zap.Error(err))
		return nil, 0, err
	}

	return posts, total, nil
}

// GetUserPosts 获取用户的动态列表
func (dao *PostDAO) GetUserPosts(ctx context.Context, userID uint64, page, pageSize int) ([]*model.Post, int64, error) {
	return dao.List(ctx, page, pageSize, nil, &userID)
}
