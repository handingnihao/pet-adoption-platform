package model

import (
	"time"
)

// CommentStatus 评论状态
type CommentStatus int

const (
	CommentStatusHidden  CommentStatus = 0 // 隐藏
	CommentStatusVisible CommentStatus = 1 // 正常
)

// Comment 评论
type Comment struct {
	ID            uint64        `json:"id" gorm:"primaryKey;autoIncrement;comment:评论ID"`
	PostID        uint64        `json:"post_id" gorm:"not null;index;comment:动态ID"`
	UserID        uint64        `json:"user_id" gorm:"not null;index;comment:评论者ID"`
	ParentID      uint64        `json:"parent_id" gorm:"default:0;index;comment:父评论ID, 0为一级评论"`
	ReplyToUserID *uint64       `json:"reply_to_user_id" gorm:"comment:回复的用户ID"`
	Content       string        `json:"content" gorm:"type:text;not null;comment:评论内容"`
	LikeCount     int           `json:"like_count" gorm:"default:0;comment:点赞数"`
	Status        CommentStatus `json:"status" gorm:"default:1;comment:状态: 0-隐藏, 1-正常"`
	CreatedAt     time.Time     `json:"created_at" gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt     time.Time     `json:"updated_at" gorm:"autoUpdateTime;comment:更新时间"`
	DeletedAt     *time.Time    `json:"deleted_at" gorm:"index;comment:删除时间"`

	// 关联
	User        *User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	ReplyToUser *User      `json:"reply_to_user,omitempty" gorm:"foreignKey:ReplyToUserID"`
	Replies     []*Comment `json:"replies,omitempty" gorm:"-"` // 子评论，不存储
}

// TableName 表名
func (Comment) TableName() string {
	return "comments"
}

// CommentCreateRequest 创建评论请求
type CommentCreateRequest struct {
	PostID        uint64  `json:"post_id" binding:"required"`
	ParentID      uint64  `json:"parent_id"`
	ReplyToUserID *uint64 `json:"reply_to_user_id"`
	Content       string  `json:"content" binding:"required,max=1000"`
}

// CommentInfo 评论信息响应
type CommentInfo struct {
	ID              uint64         `json:"id"`
	PostID          uint64         `json:"post_id"`
	UserID          uint64         `json:"user_id"`
	Username        string         `json:"username"`
	UserAvatar      string         `json:"user_avatar"`
	ParentID        uint64         `json:"parent_id"`
	ReplyToUserID   *uint64        `json:"reply_to_user_id,omitempty"`
	ReplyToUsername string         `json:"reply_to_username,omitempty"`
	Content         string         `json:"content"`
	LikeCount       int            `json:"like_count"`
	IsLiked         bool           `json:"is_liked"`
	CreatedAt       time.Time      `json:"created_at"`
	Replies         []*CommentInfo `json:"replies,omitempty"`
}

// ToCommentInfo 转换为评论信息
func (c *Comment) ToCommentInfo() *CommentInfo {
	info := &CommentInfo{
		ID:            c.ID,
		PostID:        c.PostID,
		UserID:        c.UserID,
		ParentID:      c.ParentID,
		ReplyToUserID: c.ReplyToUserID,
		Content:       c.Content,
		LikeCount:     c.LikeCount,
		CreatedAt:     c.CreatedAt,
	}

	if c.User != nil {
		info.Username = c.User.Username
		info.UserAvatar = c.User.Avatar
	}

	if c.ReplyToUser != nil {
		info.ReplyToUsername = c.ReplyToUser.Username
	}

	// 转换子评论
	if len(c.Replies) > 0 {
		info.Replies = make([]*CommentInfo, 0, len(c.Replies))
		for _, reply := range c.Replies {
			info.Replies = append(info.Replies, reply.ToCommentInfo())
		}
	}

	return info
}

// Like 点赞记录
type Like struct {
	ID         uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID     uint64    `json:"user_id" gorm:"not null;index:idx_user_target,priority:1;comment:用户ID"`
	TargetType string    `json:"target_type" gorm:"size:20;not null;index:idx_user_target,priority:2;comment:目标类型: post/comment"`
	TargetID   uint64    `json:"target_id" gorm:"not null;index:idx_user_target,priority:3;comment:目标ID"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// TableName 表名
func (Like) TableName() string {
	return "likes"
}
