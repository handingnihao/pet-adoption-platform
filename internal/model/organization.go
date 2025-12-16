package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// OrganizationStatus 机构状态
type OrganizationStatus int

const (
	// OrganizationStatusPending 待审核
	OrganizationStatusPending OrganizationStatus = iota
	// OrganizationStatusApproved 已通过
	OrganizationStatusApproved
	// OrganizationStatusRejected 已拒绝
	OrganizationStatusRejected
)

// String 返回状态的字符串表示
func (s OrganizationStatus) String() string {
	switch s {
	case OrganizationStatusPending:
		return "pending"
	case OrganizationStatusApproved:
		return "approved"
	case OrganizationStatusRejected:
		return "rejected"
	default:
		return "unknown"
	}
}

// MarshalJSON 自定义JSON序列化
func (s OrganizationStatus) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

// Scan 实现 sql.Scanner 接口，支持从字符串或整数读取
func (s *OrganizationStatus) Scan(value interface{}) error {
	if value == nil {
		*s = OrganizationStatusPending
		return nil
	}

	switch v := value.(type) {
	case int64:
		*s = OrganizationStatus(v)
	case []uint8:
		str := string(v)
		return s.parseString(str)
	case string:
		return s.parseString(v)
	default:
		return fmt.Errorf("cannot scan type %T into OrganizationStatus", value)
	}
	return nil
}

// parseString 解析字符串状态
func (s *OrganizationStatus) parseString(str string) error {
	switch str {
	case "pending", "0":
		*s = OrganizationStatusPending
	case "approved", "1":
		*s = OrganizationStatusApproved
	case "rejected", "2":
		*s = OrganizationStatusRejected
	default:
		return fmt.Errorf("unknown organization status string: %s", str)
	}
	return nil
}

// Value 实现 driver.Valuer 接口，写入数据库时转为整数
func (s OrganizationStatus) Value() (driver.Value, error) {
	return int64(s), nil
}

// Organization 机构模型
type Organization struct {
	ID             uint64             `json:"id" gorm:"primaryKey;autoIncrement;comment:机构ID"`
	Name           string             `json:"name" gorm:"size:100;not null;comment:机构名称"`
	Type           string             `json:"type" gorm:"size:50;comment:机构类型"`
	Logo           string             `json:"logo" gorm:"size:255;comment:机构logo"`
	Description    string             `json:"description" gorm:"type:text;comment:机构描述"`
	Province       string             `json:"province" gorm:"size:50;comment:省份"`
	City           string             `json:"city" gorm:"size:50;comment:城市"`
	Address        string             `json:"address" gorm:"size:255;comment:机构地址"`
	ContactName    string             `json:"contact_name" gorm:"size:50;comment:联系人姓名"`
	ContactPhone   string             `json:"contact_phone" gorm:"size:20;comment:联系人电话"`
	ContactEmail   string             `json:"contact_email" gorm:"size:100;comment:联系人邮箱"`
	Phone          string             `json:"phone" gorm:"size:20;comment:联系电话"`
	Email          string             `json:"email" gorm:"size:100;comment:联系邮箱"`
	CredentialUrls []string           `json:"credential_urls" gorm:"type:json;serializer:json;comment:资历证明图片"`
	Status         OrganizationStatus `json:"status" gorm:"default:0;comment:状态(0:待审核 1:已通过 2:已拒绝)"`
	RejectReason   string             `json:"reject_reason" gorm:"size:255;comment:拒绝原因"`
	CreatedBy      uint64             `json:"created_by" gorm:"comment:创建人ID"`
	UpdatedBy      uint64             `json:"updated_by" gorm:"comment:更新人ID"`
	CreatedAt      time.Time          `json:"created_at" gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt      time.Time          `json:"updated_at" gorm:"autoUpdateTime;comment:更新时间"`
	DeletedAt      *time.Time         `json:"deleted_at" gorm:"index;comment:删除时间"`
}

// TableName 指定表名
func (Organization) TableName() string {
	return "organizations"
}

// OrganizationCreateRequest 创建机构请求
type OrganizationCreateRequest struct {
	Name        string `json:"name" binding:"required,max=100"`
	Logo        string `json:"logo" binding:"max=255"`
	Description string `json:"description"`
	Address     string `json:"address" binding:"required,max=255"`
	Phone       string `json:"phone" binding:"required,max=20"`
	Email       string `json:"email" binding:"required,email,max=100"`
}

// OrganizationUpdateRequest 更新机构请求
type OrganizationUpdateRequest struct {
	Name        *string `json:"name,omitempty"`
	Logo        *string `json:"logo,omitempty"`
	Description *string `json:"description,omitempty"`
	Address     *string `json:"address,omitempty"`
	Phone       *string `json:"phone,omitempty"`
	Email       *string `json:"email,omitempty"`
}

// OrganizationStatusUpdateRequest 更新机构状态请求
type OrganizationStatusUpdateRequest struct {
	Status       OrganizationStatus `json:"status" binding:"required,oneof=0 1 2"`
	RejectReason string            `json:"reject_reason"`
}

// OrganizationInfo 机构信息
type OrganizationInfo struct {
	ID          uint64            `json:"id"`
	Name        string            `json:"name"`
	Logo        string            `json:"logo"`
	Description string            `json:"description"`
	Address     string            `json:"address"`
	Phone       string            `json:"phone"`
	Email       string            `json:"email"`
	Status      OrganizationStatus `json:"status"`
	CreatedAt   time.Time         `json:"created_at"`
}

// ToOrganizationInfo 转换为机构信息
func (o *Organization) ToOrganizationInfo() *OrganizationInfo {
	return &OrganizationInfo{
		ID:          o.ID,
		Name:        o.Name,
		Logo:        o.Logo,
		Description: o.Description,
		Address:     o.Address,
		Phone:       o.Phone,
		Email:       o.Email,
		Status:      o.Status,
		CreatedAt:   o.CreatedAt,
	}
}
