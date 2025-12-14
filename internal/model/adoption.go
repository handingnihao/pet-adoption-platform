package model

import "time"

// ApplicationStatus 申请状态
type ApplicationStatus string

const (
	ApplicationStatusPending    ApplicationStatus = "pending"     // 待审核
	ApplicationStatusReviewing  ApplicationStatus = "reviewing"   // 审核中
	ApplicationStatusInterview  ApplicationStatus = "interview"   // 待面试
	ApplicationStatusHomeVisit  ApplicationStatus = "home_visit"  // 待家访
	ApplicationStatusApproved   ApplicationStatus = "approved"    // 已通过
	ApplicationStatusRejected   ApplicationStatus = "rejected"    // 已拒绝
	ApplicationStatusCancelled  ApplicationStatus = "cancelled"   // 已取消
)

// HousingType 住房类型
type HousingType string

const (
	HousingTypeApartment HousingType = "apartment" // 公寓
	HousingTypeHouse     HousingType = "house"     // 住宅
	HousingTypeVilla     HousingType = "villa"     // 别墅
	HousingTypeOther     HousingType = "other"     // 其他
)

// AdoptionStatus 领养状态
type AdoptionStatus string

const (
	AdoptionStatusActive   AdoptionStatus = "active"   // 正常
	AdoptionStatusReturned AdoptionStatus = "returned" // 已退回
	AdoptionStatusDeceased AdoptionStatus = "deceased" // 已去世
)

// AdoptionApplication 领养申请模型
type AdoptionApplication struct {
	ID             uint64            `json:"id" gorm:"primaryKey"`
	ApplicationNo  string            `json:"application_no" gorm:"size:50;uniqueIndex;not null;comment:申请编号"`
	UserID         int64             `json:"user_id" gorm:"not null;index;comment:申请人ID"`
	PetID          uint64            `json:"pet_id" gorm:"not null;index;comment:宠物ID"`
	OrganizationID uint64            `json:"organization_id" gorm:"not null;index;comment:机构ID"`

	// 申请人信息
	ApplicantName   string `json:"applicant_name" gorm:"size:50;not null;comment:申请人姓名"`
	ApplicantPhone  string `json:"applicant_phone" gorm:"size:20;not null;comment:申请人电话"`
	ApplicantIDCard string `json:"applicant_id_card" gorm:"size:18;comment:身份证号"`
	ApplicantAddr   string `json:"applicant_address" gorm:"column:applicant_address;type:text;comment:地址"`

	// 家庭情况
	HousingType    HousingType `json:"housing_type" gorm:"size:20;comment:住房类型"`
	HousingArea    int         `json:"housing_area" gorm:"comment:住房面积(平方米)"`
	HasYard        bool        `json:"has_yard" gorm:"default:0;comment:是否有院子"`
	FamilyMembers  int         `json:"family_members" gorm:"comment:家庭成员数"`
	HasChildren    bool        `json:"has_children" gorm:"default:0;comment:是否有孩子"`
	ChildrenAge    string      `json:"children_age" gorm:"size:50;comment:孩子年龄"`
	FamilyAgree    bool        `json:"family_agree" gorm:"default:0;comment:家人是否同意"`

	// 养宠经验
	HasPetExperience bool   `json:"has_pet_experience" gorm:"default:0;comment:是否有养宠经验"`
	PetExperience    string `json:"pet_experience" gorm:"type:text;comment:养宠经历"`
	CurrentPets      string `json:"current_pets" gorm:"type:text;comment:现有宠物"`

	// 领养原因
	AdoptionReason string `json:"adoption_reason" gorm:"type:text;comment:领养原因"`
	HowToCare      string `json:"how_to_care" gorm:"type:text;comment:如何照顾"`
	EmergencyPlan  string `json:"emergency_plan" gorm:"type:text;comment:应急计划"`

	// 附件
	IDCardImage     string `json:"id_card_image" gorm:"size:255;comment:身份证照片URL"`
	HousingProof    string `json:"housing_proof" gorm:"size:255;comment:住房证明URL"`
	AdditionalFiles string `json:"additional_files" gorm:"type:text;comment:其他附件(JSON)"`

	// 审核流程
	Status           ApplicationStatus `json:"status" gorm:"size:20;default:pending;index;comment:状态"`
	ReviewerID       *int64            `json:"reviewer_id,omitempty" gorm:"comment:审核人ID"`
	ReviewComment    string            `json:"review_comment" gorm:"type:text;comment:审核意见"`
	InterviewTime    *time.Time        `json:"interview_time,omitempty" gorm:"comment:面试时间"`
	HomeVisitTime    *time.Time        `json:"home_visit_time,omitempty" gorm:"comment:家访时间"`
	ApprovedAt       *time.Time        `json:"approved_at,omitempty" gorm:"comment:批准时间"`
	RejectedAt       *time.Time        `json:"rejected_at,omitempty" gorm:"comment:拒绝时间"`
	RejectionReason  string            `json:"rejection_reason" gorm:"type:text;comment:拒绝原因"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`

	// 关联
	User         *User         `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Pet          *Pet          `json:"pet,omitempty" gorm:"foreignKey:PetID"`
	Organization *Organization `json:"organization,omitempty" gorm:"foreignKey:OrganizationID"`
	Reviewer     *User         `json:"reviewer,omitempty" gorm:"foreignKey:ReviewerID"`
}

func (AdoptionApplication) TableName() string {
	return "adoption_applications"
}

// Adoption 领养记录模型
type Adoption struct {
	ID             uint64    `json:"id" gorm:"primaryKey"`
	ApplicationID  uint64    `json:"application_id" gorm:"uniqueIndex;not null;comment:申请ID"`
	UserID         int64     `json:"user_id" gorm:"not null;index;comment:领养人ID"`
	PetID          uint64    `json:"pet_id" gorm:"not null;index;comment:宠物ID"`
	OrganizationID uint64    `json:"organization_id" gorm:"not null;index;comment:机构ID"`
	AdoptionDate   time.Time `json:"adoption_date" gorm:"type:date;not null;index;comment:领养日期"`

	HandoverLocation string `json:"handover_location" gorm:"size:255;comment:交接地点"`
	AgreementURL     string `json:"agreement_url" gorm:"size:255;comment:协议文件URL"`
	AgreementSigned  bool   `json:"agreement_signed" gorm:"default:0;comment:协议是否签署"`
	FollowUpPlan     string `json:"follow_up_plan" gorm:"type:text;comment:回访计划(JSON)"`

	Status AdoptionStatus `json:"status" gorm:"size:20;default:active;index;comment:状态"`
	Notes  string         `json:"notes" gorm:"type:text;comment:备注"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`

	// 关联
	Application  *AdoptionApplication `json:"application,omitempty" gorm:"foreignKey:ApplicationID"`
	User         *User                `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Pet          *Pet                 `json:"pet,omitempty" gorm:"foreignKey:PetID"`
	Organization *Organization        `json:"organization,omitempty" gorm:"foreignKey:OrganizationID"`
}

func (Adoption) TableName() string {
	return "adoptions"
}

// Organization 机构模型（临时占位，后续实现）
type Organization struct {
	ID   uint64 `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"size:100"`
}

func (Organization) TableName() string {
	return "organizations"
}

// ========== 请求和响应结构 ==========

// ApplicationCreateRequest 创建领养申请请求
type ApplicationCreateRequest struct {
	PetID          uint64      `json:"pet_id" binding:"required"`
	OrganizationID uint64      `json:"organization_id" binding:"required"`
	ApplicantName  string      `json:"applicant_name" binding:"required,min=2,max=50"`
	ApplicantPhone string      `json:"applicant_phone" binding:"required,len=11"`
	ApplicantAddr  string      `json:"applicant_address" binding:"required,max=500"`
	HousingType    HousingType `json:"housing_type" binding:"required,oneof=apartment house villa other"`
	HousingArea    int         `json:"housing_area" binding:"required,min=10,max=10000"`
	HasYard        bool        `json:"has_yard"`
	FamilyMembers  int         `json:"family_members" binding:"required,min=1,max=20"`
	HasChildren    bool        `json:"has_children"`
	ChildrenAge    string      `json:"children_age" binding:"max=50"`
	FamilyAgree    bool        `json:"family_agree" binding:"required"`
	AdoptionReason string      `json:"adoption_reason" binding:"required,min=10,max=2000"`
	HowToCare      string      `json:"how_to_care" binding:"required,min=10,max=2000"`
	EmergencyPlan  string      `json:"emergency_plan" binding:"required,min=10,max=1000"`
}

// ApplicationUpdateRequest 更新申请请求
type ApplicationUpdateRequest struct {
	ApplicantName    *string      `json:"applicant_name" binding:"omitempty,min=2,max=50"`
	ApplicantPhone   *string      `json:"applicant_phone" binding:"omitempty,len=11"`
	ApplicantAddr    *string      `json:"applicant_address" binding:"omitempty,max=500"`
	HousingType      *HousingType `json:"housing_type" binding:"omitempty,oneof=apartment house villa other"`
	HousingArea      *int         `json:"housing_area" binding:"omitempty,min=10,max=10000"`
	HasYard          *bool        `json:"has_yard"`
	FamilyMembers    *int         `json:"family_members" binding:"omitempty,min=1,max=20"`
	HasChildren      *bool        `json:"has_children"`
	ChildrenAge      *string      `json:"children_age" binding:"omitempty,max=50"`
	FamilyAgree      *bool        `json:"family_agree"`
	HasPetExperience *bool        `json:"has_pet_experience"`
	PetExperience    *string      `json:"pet_experience" binding:"omitempty,max=2000"`
	CurrentPets      *string      `json:"current_pets" binding:"omitempty,max=500"`
	AdoptionReason   *string      `json:"adoption_reason" binding:"omitempty,min=10,max=2000"`
	HowToCare        *string      `json:"how_to_care" binding:"omitempty,min=10,max=2000"`
	EmergencyPlan    *string      `json:"emergency_plan" binding:"omitempty,min=10,max=1000"`
}

// ApplicationReviewRequest 审核申请请求
type ApplicationReviewRequest struct {
	Action  string `json:"action" binding:"required,oneof=approve reject interview home_visit"`
	Comment string `json:"comment" binding:"max=1000"`
	Reason  string `json:"reason" binding:"required_if=Action reject,max=500"`
}

// AdoptionCreateRequest 创建领养记录请求
type AdoptionCreateRequest struct {
	ApplicationID    uint64    `json:"application_id" binding:"required"`
	AdoptionDate     time.Time `json:"adoption_date" binding:"required"`
	HandoverLocation string    `json:"handover_location" binding:"required,max=255"`
	AgreementURL     string    `json:"agreement_url" binding:"max=255"`
	Notes            string    `json:"notes" binding:"max=1000"`
}

// ApplicationResponse 申请响应
type ApplicationResponse struct {
	ID              uint64             `json:"id"`
	ApplicationNo   string             `json:"application_no"`
	UserID          int64              `json:"user_id"`
	PetID           uint64             `json:"pet_id"`
	OrganizationID  uint64             `json:"organization_id"`
	ApplicantName   string             `json:"applicant_name"`
	ApplicantPhone  string             `json:"applicant_phone"`
	ApplicantAddr   string             `json:"applicant_address"`
	HousingType     HousingType        `json:"housing_type"`
	HousingArea     int                `json:"housing_area"`
	HasYard         bool               `json:"has_yard"`
	FamilyMembers   int                `json:"family_members"`
	HasChildren     bool               `json:"has_children"`
	FamilyAgree     bool               `json:"family_agree"`
	AdoptionReason  string             `json:"adoption_reason"`
	HowToCare       string             `json:"how_to_care"`
	EmergencyPlan   string             `json:"emergency_plan"`
	Status          ApplicationStatus  `json:"status"`
	ReviewComment   string             `json:"review_comment,omitempty"`
	RejectionReason string             `json:"rejection_reason,omitempty"`
	InterviewTime   *time.Time         `json:"interview_time,omitempty"`
	HomeVisitTime   *time.Time         `json:"home_visit_time,omitempty"`
	ApprovedAt      *time.Time         `json:"approved_at,omitempty"`
	RejectedAt      *time.Time         `json:"rejected_at,omitempty"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
	PetInfo         *PetBasicInfo      `json:"pet_info,omitempty"`
	UserInfo        *UserBasicInfo     `json:"user_info,omitempty"`
	ReviewerInfo    *UserBasicInfo     `json:"reviewer_info,omitempty"`
}

// AdoptionResponse 领养记录响应
type AdoptionResponse struct {
	ID               uint64         `json:"id"`
	ApplicationID    uint64         `json:"application_id"`
	UserID           int64          `json:"user_id"`
	PetID            uint64         `json:"pet_id"`
	OrganizationID   uint64         `json:"organization_id"`
	AdoptionDate     time.Time      `json:"adoption_date"`
	HandoverLocation string         `json:"handover_location"`
	AgreementURL     string         `json:"agreement_url"`
	AgreementSigned  bool           `json:"agreement_signed"`
	Status           AdoptionStatus `json:"status"`
	Notes            string         `json:"notes"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	PetInfo          *PetBasicInfo  `json:"pet_info,omitempty"`
	UserInfo         *UserBasicInfo `json:"user_info,omitempty"`
}

// PetBasicInfo 宠物基本信息
type PetBasicInfo struct {
	ID         uint64     `json:"id"`
	Name       string     `json:"name"`
	Type       PetType    `json:"type"`
	Breed      string     `json:"breed"`
	Age        int        `json:"age"`
	Gender     PetGender  `json:"gender"`
	CoverPhoto string     `json:"cover_photo"`
	Status     PetStatus  `json:"status"`
}

// ApplicationQueryRequest 申请查询请求
type ApplicationQueryRequest struct {
	Status         ApplicationStatus `json:"status" form:"status"`
	PetID          uint64            `json:"pet_id" form:"pet_id"`
	OrganizationID uint64            `json:"organization_id" form:"organization_id"`
	StartDate      string            `json:"start_date" form:"start_date"`
	EndDate        string            `json:"end_date" form:"end_date"`
	Page           int               `json:"page" form:"page" binding:"min=1"`
	PageSize       int               `json:"page_size" form:"page_size" binding:"min=1,max=100"`
}
