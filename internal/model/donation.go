package model

import (
	"time"
)

// DonationType 捐赠类型
type DonationType string

const (
	DonationTypeMoney   DonationType = "money"   // 资金捐赠
	DonationTypeSupply  DonationType = "supply"  // 物资捐赠
	DonationTypeService DonationType = "service" // 服务捐赠
)

// DonationStatus 捐赠状态
type DonationStatus int

const (
	DonationStatusPending   DonationStatus = 0 // 待确认
	DonationStatusConfirmed DonationStatus = 1 // 已确认
	DonationStatusCompleted DonationStatus = 2 // 已完成
	DonationStatusCancelled DonationStatus = 3 // 已取消
)

// Donation 捐赠模型
type Donation struct {
	ID             uint64         `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID         uint64         `gorm:"column:user_id;not null;index" json:"user_id"`
	OrganizationID *uint64        `gorm:"column:organization_id;index" json:"organization_id,omitempty"` // 捐赠给哪个机构，可选
	Type           DonationType   `gorm:"column:type;type:varchar(20);not null" json:"type"`
	Amount         float64        `gorm:"column:amount;type:decimal(10,2)" json:"amount"`           // 资金金额
	SupplyItems    string         `gorm:"column:supply_items;type:text" json:"supply_items"`        // 物资清单（JSON）
	ServiceDesc    string         `gorm:"column:service_desc;type:text" json:"service_desc"`        // 服务描述
	Message        string         `gorm:"column:message;type:text" json:"message"`                  // 捐赠留言
	IsAnonymous    bool           `gorm:"column:is_anonymous;default:false" json:"is_anonymous"`    // 是否匿名
	Status         DonationStatus `gorm:"column:status;type:tinyint;default:0" json:"status"`
	PaymentMethod  string         `gorm:"column:payment_method;type:varchar(50)" json:"payment_method"` // 支付方式
	TransactionID  string         `gorm:"column:transaction_id;type:varchar(100)" json:"transaction_id"` // 交易ID
	ConfirmedAt    *time.Time     `gorm:"column:confirmed_at" json:"confirmed_at,omitempty"`
	ConfirmedBy    *uint64        `gorm:"column:confirmed_by" json:"confirmed_by,omitempty"`
	CreatedAt      time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	// 关联
	User         *User         `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Organization *Organization `gorm:"foreignKey:OrganizationID" json:"organization,omitempty"`
}

func (Donation) TableName() string {
	return "donations"
}

// DonationInfo 捐赠信息响应
type DonationInfo struct {
	ID               uint64         `json:"id"`
	UserID           uint64         `json:"user_id"`
	Username         string         `json:"username"`
	UserAvatar       string         `json:"user_avatar"`
	OrganizationID   *uint64        `json:"organization_id,omitempty"`
	OrganizationName string         `json:"organization_name,omitempty"`
	Type             DonationType   `json:"type"`
	Amount           float64        `json:"amount"`
	SupplyItems      []string       `json:"supply_items"`
	ServiceDesc      string         `json:"service_desc"`
	Message          string         `json:"message"`
	IsAnonymous      bool           `json:"is_anonymous"`
	Status           DonationStatus `json:"status"`
	StatusText       string         `json:"status_text"`
	CreatedAt        time.Time      `json:"created_at"`
}

// ToInfo 转换为响应信息
func (d *Donation) ToInfo() *DonationInfo {
	info := &DonationInfo{
		ID:          d.ID,
		UserID:      d.UserID,
		Type:        d.Type,
		Amount:      d.Amount,
		ServiceDesc: d.ServiceDesc,
		Message:     d.Message,
		IsAnonymous: d.IsAnonymous,
		Status:      d.Status,
		CreatedAt:   d.CreatedAt,
	}

	// 状态文本
	switch d.Status {
	case DonationStatusPending:
		info.StatusText = "待确认"
	case DonationStatusConfirmed:
		info.StatusText = "已确认"
	case DonationStatusCompleted:
		info.StatusText = "已完成"
	case DonationStatusCancelled:
		info.StatusText = "已取消"
	}

	// 用户信息
	if d.User != nil {
		if d.IsAnonymous {
			info.Username = "匿名用户"
			info.UserAvatar = ""
		} else {
			info.Username = d.User.Username
			info.UserAvatar = d.User.Avatar
		}
	}

	// 机构信息
	if d.OrganizationID != nil {
		info.OrganizationID = d.OrganizationID
		if d.Organization != nil {
			info.OrganizationName = d.Organization.Name
		}
	}

	// 物资清单
	if d.SupplyItems != "" {
		info.SupplyItems = ParseJSONStringArray(d.SupplyItems)
	} else {
		info.SupplyItems = []string{}
	}

	return info
}

// DonationCreateRequest 创建捐赠请求
type DonationCreateRequest struct {
	OrganizationID *uint64      `json:"organization_id"`
	Type           DonationType `json:"type" binding:"required,oneof=money supply service"`
	Amount         float64      `json:"amount"`                              // 资金金额
	SupplyItems    []string     `json:"supply_items"`                        // 物资清单
	ServiceDesc    string       `json:"service_desc"`                        // 服务描述
	Message        string       `json:"message" binding:"max=500"`           // 留言
	IsAnonymous    bool         `json:"is_anonymous"`                        // 是否匿名
	PaymentMethod  string       `json:"payment_method"`                      // 支付方式
}

// DonationConfirmRequest 确认捐赠请求
type DonationConfirmRequest struct {
	TransactionID string `json:"transaction_id"` // 交易ID
}

// DonationStatistics 捐赠统计
type DonationStatistics struct {
	TotalDonations   int64   `json:"total_donations"`    // 总捐赠次数
	TotalAmount      float64 `json:"total_amount"`       // 总捐赠金额
	TotalDonors      int64   `json:"total_donors"`       // 捐赠人数
	MonthlyDonations int64   `json:"monthly_donations"`  // 本月捐赠次数
	MonthlyAmount    float64 `json:"monthly_amount"`     // 本月捐赠金额
}
