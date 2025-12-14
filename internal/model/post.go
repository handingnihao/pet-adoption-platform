package model

import (
	"time"
)

// PostType 动态类型
type PostType string

const (
	PostTypeStory     PostType = "story"     // 领养故事
	PostTypeKnowledge PostType = "knowledge" // 养宠知识
	PostTypeDaily     PostType = "daily"     // 日常分享
	PostTypeOther     PostType = "other"     // 其他
)

// PostStatus 动态状态
type PostStatus int

const (
	PostStatusHidden  PostStatus = 0 // 隐藏
	PostStatusVisible PostStatus = 1 // 正常
)

// Post 社区动态
type Post struct {
	ID           uint64     `json:"id" gorm:"primaryKey;autoIncrement;comment:动态ID"`
	UserID       uint64     `json:"user_id" gorm:"not null;index;comment:发布者ID"`
	Title        string     `json:"title" gorm:"size:200;comment:标题"`
	Content      string     `json:"content" gorm:"type:text;not null;comment:内容"`
	Images       string     `json:"images" gorm:"type:text;comment:图片(JSON)"`
	VideoURL     string     `json:"video_url" gorm:"size:255;comment:视频URL"`
	TopicIDs     string     `json:"topic_ids" gorm:"type:text;comment:话题ID(JSON)"`
	PetID        *uint64    `json:"pet_id" gorm:"comment:关联宠物ID"`
	Type         PostType   `json:"type" gorm:"type:enum('story','knowledge','daily','other');default:'daily';comment:类型"`
	ViewCount    int        `json:"view_count" gorm:"default:0;comment:浏览次数"`
	LikeCount    int        `json:"like_count" gorm:"default:0;comment:点赞数"`
	CommentCount int        `json:"comment_count" gorm:"default:0;comment:评论数"`
	ShareCount   int        `json:"share_count" gorm:"default:0;comment:分享数"`
	Status       PostStatus `json:"status" gorm:"default:1;comment:状态: 0-隐藏, 1-正常"`
	IsTop        int        `json:"is_top" gorm:"default:0;comment:是否置顶"`
	CreatedAt    time.Time  `json:"created_at" gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"autoUpdateTime;comment:更新时间"`
	DeletedAt    *time.Time `json:"deleted_at" gorm:"index;comment:删除时间"`

	// 关联
	User *User `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Pet  *Pet  `json:"pet,omitempty" gorm:"foreignKey:PetID"`
}

// TableName 表名
func (Post) TableName() string {
	return "posts"
}

// PostCreateRequest 创建动态请求
type PostCreateRequest struct {
	Title    string   `json:"title" binding:"max=200"`
	Content  string   `json:"content" binding:"required,max=10000"`
	Images   []string `json:"images"`
	VideoURL string   `json:"video_url"`
	TopicIDs []int64  `json:"topic_ids"`
	PetID    *uint64  `json:"pet_id"`
	Type     PostType `json:"type" binding:"omitempty,oneof=story knowledge daily other"`
}

// PostUpdateRequest 更新动态请求
type PostUpdateRequest struct {
	Title    *string   `json:"title"`
	Content  *string   `json:"content"`
	Images   *[]string `json:"images"`
	VideoURL *string   `json:"video_url"`
	TopicIDs *[]int64  `json:"topic_ids"`
	PetID    *uint64   `json:"pet_id"`
	Type     *PostType `json:"type"`
}

// PostInfo 动态信息响应
type PostInfo struct {
	ID           uint64    `json:"id"`
	UserID       uint64    `json:"user_id"`
	Username     string    `json:"username"`
	UserAvatar   string    `json:"user_avatar"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Images       []string  `json:"images"`
	VideoURL     string    `json:"video_url"`
	PetID        *uint64   `json:"pet_id,omitempty"`
	Type         PostType  `json:"type"`
	ViewCount    int       `json:"view_count"`
	LikeCount    int       `json:"like_count"`
	CommentCount int       `json:"comment_count"`
	ShareCount   int       `json:"share_count"`
	IsTop        int       `json:"is_top"`
	IsLiked      bool      `json:"is_liked"`
	CreatedAt    time.Time `json:"created_at"`
}

// ToPostInfo 转换为动态信息
func (p *Post) ToPostInfo() *PostInfo {
	info := &PostInfo{
		ID:           p.ID,
		UserID:       p.UserID,
		Title:        p.Title,
		Content:      p.Content,
		VideoURL:     p.VideoURL,
		PetID:        p.PetID,
		Type:         p.Type,
		ViewCount:    p.ViewCount,
		LikeCount:    p.LikeCount,
		CommentCount: p.CommentCount,
		ShareCount:   p.ShareCount,
		IsTop:        p.IsTop,
		CreatedAt:    p.CreatedAt,
	}

	// 解析图片JSON
	if p.Images != "" {
		info.Images = ParseJSONStringArray(p.Images)
	} else {
		info.Images = []string{}
	}

	// 用户信息
	if p.User != nil {
		info.Username = p.User.Username
		info.UserAvatar = p.User.Avatar
	}

	return info
}
