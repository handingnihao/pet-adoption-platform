package dao

import (
	"context"
	"errors"
	"pet-adoption-platform/internal/model"
	"pet-adoption-platform/pkg/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// LikeDAO 点赞数据访问对象
type LikeDAO struct {
	db *gorm.DB
}

// NewLikeDAO 创建点赞DAO实例
func NewLikeDAO(db *gorm.DB) *LikeDAO {
	return &LikeDAO{db: db}
}

// Create 创建点赞记录
func (dao *LikeDAO) Create(ctx context.Context, like *model.Like) error {
	if like == nil {
		return errors.New("like is nil")
	}

	if err := dao.db.WithContext(ctx).Create(like).Error; err != nil {
		logger.Error("创建点赞记录失败", zap.Error(err))
		return err
	}

	return nil
}

// Delete 删除点赞记录
func (dao *LikeDAO) Delete(ctx context.Context, userID uint64, targetType string, targetID uint64) error {
	result := dao.db.WithContext(ctx).
		Where("user_id = ? AND target_type = ? AND target_id = ?", userID, targetType, targetID).
		Delete(&model.Like{})

	if result.Error != nil {
		logger.Error("删除点赞记录失败", zap.Error(result.Error))
		return result.Error
	}

	return nil
}

// Exists 检查是否已点赞
func (dao *LikeDAO) Exists(ctx context.Context, userID uint64, targetType string, targetID uint64) (bool, error) {
	var count int64
	err := dao.db.WithContext(ctx).Model(&model.Like{}).
		Where("user_id = ? AND target_type = ? AND target_id = ?", userID, targetType, targetID).
		Count(&count).Error

	if err != nil {
		logger.Error("检查点赞状态失败", zap.Error(err))
		return false, err
	}

	return count > 0, nil
}

// GetUserLikedTargets 获取用户点赞的目标ID列表
func (dao *LikeDAO) GetUserLikedTargets(ctx context.Context, userID uint64, targetType string, targetIDs []uint64) (map[uint64]bool, error) {
	if len(targetIDs) == 0 {
		return map[uint64]bool{}, nil
	}

	var likes []model.Like
	err := dao.db.WithContext(ctx).
		Where("user_id = ? AND target_type = ? AND target_id IN ?", userID, targetType, targetIDs).
		Find(&likes).Error

	if err != nil {
		logger.Error("获取用户点赞列表失败", zap.Error(err))
		return nil, err
	}

	result := make(map[uint64]bool)
	for _, like := range likes {
		result[like.TargetID] = true
	}

	return result, nil
}
