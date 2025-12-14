package model

import (
	"time"
)

// UserRole 用户角色枚举
type UserRole string

const (
	RoleUser         UserRole = "user"         // 普通用户
	RoleOrganization UserRole = "organization" // 机构用户
	RoleVolunteer    UserRole = "volunteer"    // 志愿者
	RoleAdmin        UserRole = "admin"        // 管理员
)

// UserStatus 用户状态
type UserStatus int

const (
	UserStatusDisabled UserStatus = 0 // 禁用
	UserStatusNormal   UserStatus = 1 // 正常
)

// Gender 性别
type Gender int

const (
	GenderUnknown Gender = 0 // 未知
	GenderMale    Gender = 1 // 男
	GenderFemale  Gender = 2 // 女
)

// User 用户模型
type User struct {
	ID          int64      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Username    string     `gorm:"column:username;type:varchar(50);uniqueIndex;not null" json:"username"`
	Password    string     `gorm:"column:password;type:varchar(255);not null" json:"-"` // 不返回给前端
	RealName    string     `gorm:"column:real_name;type:varchar(50)" json:"real_name,omitempty"`
	Phone       string     `gorm:"column:phone;type:varchar(20);uniqueIndex" json:"phone,omitempty"`
	Email       string     `gorm:"column:email;type:varchar(100);uniqueIndex" json:"email,omitempty"`
	Avatar      string     `gorm:"column:avatar;type:varchar(255)" json:"avatar,omitempty"`
	Gender      Gender     `gorm:"column:gender;type:tinyint" json:"gender"`
	Birthday    *time.Time `gorm:"column:birthday;type:date" json:"birthday,omitempty"`
	IDCard      string     `gorm:"column:id_card;type:varchar(18)" json:"id_card,omitempty"`
	Address     string     `gorm:"column:address;type:text" json:"address,omitempty"`
	Role        UserRole   `gorm:"column:role;type:enum('user','organization','volunteer','admin');default:user" json:"role"`
	Status      UserStatus `gorm:"column:status;type:tinyint;default:1" json:"status"`
	LastLoginAt *time.Time `gorm:"column:last_login_at;type:timestamp" json:"last_login_at,omitempty"`
	LastLoginIP string     `gorm:"column:last_login_ip;type:varchar(50)" json:"last_login_ip,omitempty"`
	CreatedAt   time.Time  `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at;type:timestamp;index" json:"-"` // 软删除
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

// IsAdmin 判断是否是管理员
func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

// IsOrganization 判断是否是机构用户
func (u *User) IsOrganization() bool {
	return u.Role == RoleOrganization
}

// IsActive 判断用户是否激活
func (u *User) IsActive() bool {
	return u.Status == UserStatusNormal
}

// MaskSensitiveInfo 脱敏处理敏感信息（用于日志等）
func (u *User) MaskSensitiveInfo() *User {
	masked := *u
	masked.Password = "******"
	if len(masked.Phone) > 7 {
		masked.Phone = masked.Phone[:3] + "****" + masked.Phone[7:]
	}
	if len(masked.Email) > 3 {
		masked.Email = masked.Email[:3] + "****@****"
	}
	if len(masked.IDCard) > 6 {
		masked.IDCard = masked.IDCard[:6] + "********" + masked.IDCard[14:]
	}
	return &masked
}

// UserRegisterRequest 用户注册请求
type UserRegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6,max=20"`
	Phone    string `json:"phone" binding:"required,len=11"`
	Email    string `json:"email" binding:"required,email"`
}

// UserLoginRequest 用户登录请求
type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserLoginResponse 用户登录响应
type UserLoginResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	User      *UserInfo `json:"user"`
}

// UserInfo 用户信息（脱敏后）
type UserInfo struct {
	ID       int64    `json:"id"`
	Username string   `json:"username"`
	RealName string   `json:"real_name,omitempty"`
	Phone    string   `json:"phone,omitempty"`
	Email    string   `json:"email,omitempty"`
	Avatar   string   `json:"avatar,omitempty"`
	Gender   Gender   `json:"gender"`
	Role     UserRole `json:"role"`
	Status   int      `json:"status"`
}

// ToUserInfo 转换为UserInfo
func (u *User) ToUserInfo() *UserInfo {
	return &UserInfo{
		ID:       u.ID,
		Username: u.Username,
		RealName: u.RealName,
		Phone:    u.Phone,
		Email:    u.Email,
		Avatar:   u.Avatar,
		Gender:   u.Gender,
		Role:     u.Role,
		Status:   int(u.Status),
	}
}

// UserUpdateRequest 用户更新请求
type UserUpdateRequest struct {
	RealName string     `json:"real_name"`
	Avatar   string     `json:"avatar"`
	Gender   Gender     `json:"gender"`
	Birthday *time.Time `json:"birthday"`
	Address  string     `json:"address"`
}

// UpdatePasswordRequest 修改密码请求
type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6,max=20"`
}
