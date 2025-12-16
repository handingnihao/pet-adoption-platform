package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// PetType 宠物类型
type PetType string

const (
	PetTypeDog    PetType = "dog"    // 狗
	PetTypeCat    PetType = "cat"    // 猫
	PetTypeRabbit PetType = "rabbit" // 兔子
	PetTypeBird   PetType = "bird"   // 鸟
	PetTypeOther  PetType = "other"  // 其他
)

// PetGender 宠物性别
type PetGender string

const (
	PetGenderMale   PetGender = "male"   // 公
	PetGenderFemale PetGender = "female" // 母
	PetGenderUnknown PetGender = "unknown" // 未知
)

// PetStatus 宠物状态
type PetStatus int

const (
	PetStatusPending  PetStatus = 0 // 待审核
	PetStatusAvailable PetStatus = 1 // 可领养
	PetStatusAdopted   PetStatus = 2 // 已领养
	PetStatusOffline   PetStatus = 3 // 已下架
)

// String 返回状态的字符串表示
func (s PetStatus) String() string {
	switch s {
	case PetStatusPending:
		return "pending"
	case PetStatusAvailable:
		return "available"
	case PetStatusAdopted:
		return "adopted"
	case PetStatusOffline:
		return "offline"
	default:
		return "unknown"
	}
}

// MarshalJSON 自定义JSON序列化
func (s PetStatus) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

// Scan 实现 sql.Scanner 接口，支持从字符串或整数读取
func (s *PetStatus) Scan(value interface{}) error {
	if value == nil {
		*s = PetStatusPending
		return nil
	}

	switch v := value.(type) {
	case int64:
		*s = PetStatus(v)
	case []uint8:
		str := string(v)
		return s.parseString(str)
	case string:
		return s.parseString(v)
	default:
		return fmt.Errorf("cannot scan type %T into PetStatus", value)
	}
	return nil
}

// parseString 解析字符串状态
func (s *PetStatus) parseString(str string) error {
	switch str {
	case "pending", "0":
		*s = PetStatusPending
	case "available", "1":
		*s = PetStatusAvailable
	case "adopted", "approved", "2":
		*s = PetStatusAdopted
	case "offline", "3":
		*s = PetStatusOffline
	default:
		return fmt.Errorf("unknown pet status string: %s", str)
	}
	return nil
}

// Value 实现 driver.Valuer 接口，写入数据库时转为整数
func (s PetStatus) Value() (driver.Value, error) {
	return int64(s), nil
}

// PetSize 宠物体型
type PetSize string

const (
	PetSizeSmall  PetSize = "small"  // 小型
	PetSizeMedium PetSize = "medium" // 中型
	PetSizeLarge  PetSize = "large"  // 大型
)

// Pet 宠物模型
type Pet struct {
	ID             uint64    `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name" gorm:"size:100;not null;comment:宠物名称"`
	Type           PetType   `json:"type" gorm:"size:20;not null;comment:宠物类型"`
	Breed          string    `json:"breed" gorm:"size:100;comment:品种"`
	Gender         PetGender `json:"gender" gorm:"size:10;comment:性别"`
	Age            int       `json:"age" gorm:"comment:年龄(月)"`
	Size           PetSize   `json:"size" gorm:"size:20;comment:体型"`
	Color          string    `json:"color" gorm:"size:50;comment:毛色"`
	Weight         float64   `json:"weight" gorm:"type:decimal(5,2);comment:体重(kg)"`
	
	// 健康信息
	IsVaccinated   bool      `json:"is_vaccinated" gorm:"default:false;comment:是否接种疫苗"`
	IsSterilized   bool      `json:"is_sterilized" gorm:"default:false;comment:是否绝育"`
	HealthStatus   string    `json:"health_status" gorm:"size:200;comment:健康状况"`
	
	// 详细信息
	Description    string    `json:"description" gorm:"type:text;comment:详细描述"`
	Character      string    `json:"character" gorm:"type:text;comment:性格特点"`
	Photos         string    `json:"photos" gorm:"type:text;comment:照片URL,逗号分隔"`
	CoverPhoto     string    `json:"cover_photo" gorm:"size:500;comment:封面图"`
	
	// 位置信息
	Province       string    `json:"province" gorm:"size:50;comment:省份"`
	City           string    `json:"city" gorm:"size:50;comment:城市"`
	District       string    `json:"district" gorm:"size:50;comment:区县"`
	Address        string    `json:"address" gorm:"size:200;comment:详细地址"`
	
	// 关联信息
	UserID         int64     `json:"user_id" gorm:"not null;index;comment:发布用户ID"`
	User           *User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
	
	// 状态和统计
	Status         PetStatus `json:"status" gorm:"default:0;index;comment:状态"`
	ViewCount      int       `json:"view_count" gorm:"default:0;comment:浏览次数"`
	FavoriteCount  int       `json:"favorite_count" gorm:"default:0;comment:收藏次数"`
	
	// 领养信息
	AdoptedBy      *int64    `json:"adopted_by,omitempty" gorm:"index;comment:领养人ID"`
	AdoptedAt      *time.Time `json:"adopted_at,omitempty" gorm:"comment:领养时间"`
	
	// 审核信息
	ReviewedBy     *int64    `json:"reviewed_by,omitempty" gorm:"comment:审核人ID"`
	ReviewedAt     *time.Time `json:"reviewed_at,omitempty" gorm:"comment:审核时间"`
	RejectReason   string    `json:"reject_reason,omitempty" gorm:"size:500;comment:拒绝原因"`
	
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty" gorm:"index"`
}

// TableName 指定表名
func (Pet) TableName() string {
	return "pets"
}

// PetCreateRequest 创建宠物请求
type PetCreateRequest struct {
	Name         string    `json:"name" binding:"required,min=1,max=100"`
	Type         PetType   `json:"type" binding:"required,oneof=dog cat rabbit hamster bird other"`
	Breed        string    `json:"breed" binding:"max=100"`
	Gender       PetGender `json:"gender" binding:"oneof=male female unknown"`
	Age          int       `json:"age" binding:"min=0,max=300"`
	Size         PetSize   `json:"size" binding:"oneof=small medium large"`
	Color        string    `json:"color" binding:"max=50"`
	Weight       float64   `json:"weight" binding:"min=0,max=999"`
	IsVaccinated bool      `json:"is_vaccinated"`
	IsSterilized bool      `json:"is_sterilized"`
	HealthStatus string    `json:"health_status" binding:"max=200"`
	Description  string    `json:"description" binding:"required,max=2000"`
	Character    string    `json:"character" binding:"max=500"`
	Photos       []string  `json:"photos"`
	CoverPhoto   string    `json:"cover_photo"`
	Province     string    `json:"province" binding:"max=50"`
	City         string    `json:"city" binding:"max=50"`
	District     string    `json:"district" binding:"max=50"`
	Address      string    `json:"address" binding:"max=200"`
}

// PetUpdateRequest 更新宠物请求
type PetUpdateRequest struct {
	Name         *string    `json:"name" binding:"omitempty,min=1,max=100"`
	Type         *PetType   `json:"type" binding:"omitempty,oneof=dog cat rabbit hamster bird other"`
	Breed        *string    `json:"breed" binding:"omitempty,max=100"`
	Gender       *PetGender `json:"gender" binding:"omitempty,oneof=male female unknown"`
	Age          *int       `json:"age" binding:"omitempty,min=0,max=300"`
	Size         *PetSize   `json:"size" binding:"omitempty,oneof=small medium large"`
	Color        *string    `json:"color" binding:"omitempty,max=50"`
	Weight       *float64   `json:"weight" binding:"omitempty,min=0,max=999"`
	IsVaccinated *bool      `json:"is_vaccinated"`
	IsSterilized *bool      `json:"is_sterilized"`
	HealthStatus *string    `json:"health_status" binding:"omitempty,max=200"`
	Description  *string    `json:"description" binding:"omitempty,max=2000"`
	Character    *string    `json:"character" binding:"omitempty,max=500"`
	Photos       []string   `json:"photos"`
	CoverPhoto   *string    `json:"cover_photo"`
	Province     *string    `json:"province" binding:"omitempty,max=50"`
	City         *string    `json:"city" binding:"omitempty,max=50"`
	District     *string    `json:"district" binding:"omitempty,max=50"`
	Address      *string    `json:"address" binding:"omitempty,max=200"`
}

// PetQueryRequest 查询宠物请求
type PetQueryRequest struct {
	Type     PetType   `json:"type" form:"type"`
	Gender   PetGender `json:"gender" form:"gender"`
	Size     PetSize   `json:"size" form:"size"`
	Status   *PetStatus `json:"status" form:"status"`
	Province string    `json:"province" form:"province"`
	City     string    `json:"city" form:"city"`
	MinAge   *int      `json:"min_age" form:"min_age"`
	MaxAge   *int      `json:"max_age" form:"max_age"`
	Keyword  string    `json:"keyword" form:"keyword"` // 搜索名称、品种、描述
	Page     int       `json:"page" form:"page" binding:"min=1"`
	PageSize int       `json:"page_size" form:"page_size" binding:"min=1,max=100"`
}

// PetResponse 宠物响应
type PetResponse struct {
	ID            uint64     `json:"id"`
	Name          string     `json:"name"`
	Type          PetType    `json:"type"`
	Breed         string     `json:"breed"`
	Gender        PetGender  `json:"gender"`
	Age           int        `json:"age"`
	Size          PetSize    `json:"size"`
	Color         string     `json:"color"`
	Weight        float64    `json:"weight"`
	IsVaccinated  bool       `json:"is_vaccinated"`
	IsSterilized  bool       `json:"is_sterilized"`
	HealthStatus  string     `json:"health_status"`
	Description   string     `json:"description"`
	Character     string     `json:"character"`
	Photos        []string   `json:"photos"`
	CoverPhoto    string     `json:"cover_photo"`
	Province      string     `json:"province"`
	City          string     `json:"city"`
	District      string     `json:"district"`
	Address       string     `json:"address"`
	Status        PetStatus  `json:"status"`
	ViewCount     int        `json:"view_count"`
	FavoriteCount int        `json:"favorite_count"`
	UserID        uint       `json:"user_id"`
	UserInfo      *UserBasicInfo `json:"user_info,omitempty"`
	AdoptedBy     *uint      `json:"adopted_by,omitempty"`
	AdoptedAt     *time.Time `json:"adopted_at,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

// UserBasicInfo 用户基本信息（用于关联显示）
type UserBasicInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}
