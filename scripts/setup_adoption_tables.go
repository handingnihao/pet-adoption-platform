package main

import (
	"fmt"
	"log"
	"pet-adoption-platform/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// 1. 加载配置
	config.Init()

	// 2. 连接数据库
	dsn := config.AppConfig.Database.MySQL.GetDSN()
	fmt.Println("连接字符串:", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	fmt.Println("✓ 数据库连接成功")

	// 3. 创建organizations表（临时表）
	fmt.Println("\n[1] 创建organizations表...")
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	createOrgSQL := `
CREATE TABLE IF NOT EXISTS organizations (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '机构ID',
    name VARCHAR(100) NOT NULL COMMENT '机构名称',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_name (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='机构表（临时）'`

	if err := db.Exec(createOrgSQL).Error; err != nil {
		log.Fatal("创建organizations表失败:", err)
	}
	fmt.Println("✓ organizations表创建成功")

	// 插入测试机构
	db.Exec("INSERT IGNORE INTO organizations (id, name) VALUES (1, '深圳宠物救助中心')")
	fmt.Println("✓ 测试机构数据插入完成")

	// 4. 创建adoption_applications表
	fmt.Println("\n[2] 创建adoption_applications表...")
	createAppSQL := `
CREATE TABLE IF NOT EXISTS adoption_applications (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '申请ID',
    application_no VARCHAR(50) UNIQUE NOT NULL COMMENT '申请编号',
    user_id BIGINT NOT NULL COMMENT '申请人ID',
    pet_id BIGINT UNSIGNED NOT NULL COMMENT '宠物ID',
    organization_id BIGINT UNSIGNED NOT NULL COMMENT '机构ID',
    
    applicant_name VARCHAR(50) NOT NULL COMMENT '申请人姓名',
    applicant_phone VARCHAR(20) NOT NULL COMMENT '申请人电话',
    applicant_id_card VARCHAR(18) COMMENT '身份证号',
    applicant_address TEXT COMMENT '地址',
    
    housing_type VARCHAR(20) COMMENT '住房类型',
    housing_area INT COMMENT '住房面积(平方米)',
    has_yard TINYINT DEFAULT 0 COMMENT '是否有院子',
    family_members INT COMMENT '家庭成员数',
    has_children TINYINT DEFAULT 0 COMMENT '是否有孩子',
    children_age VARCHAR(50) COMMENT '孩子年龄',
    family_agree TINYINT DEFAULT 0 COMMENT '家人是否同意',
    
    has_pet_experience TINYINT DEFAULT 0 COMMENT '是否有养宠经验',
    pet_experience TEXT COMMENT '养宠经历',
    current_pets TEXT COMMENT '现有宠物',
    
    adoption_reason TEXT COMMENT '领养原因',
    how_to_care TEXT COMMENT '如何照顾',
    emergency_plan TEXT COMMENT '应急计划',
    
    id_card_image VARCHAR(255) COMMENT '身份证照片URL',
    housing_proof VARCHAR(255) COMMENT '住房证明URL',
    additional_files TEXT COMMENT '其他附件(JSON)',
    
    status VARCHAR(20) DEFAULT 'pending' COMMENT '状态',
    reviewer_id BIGINT COMMENT '审核人ID',
    review_comment TEXT COMMENT '审核意见',
    interview_time TIMESTAMP NULL COMMENT '面试时间',
    home_visit_time TIMESTAMP NULL COMMENT '家访时间',
    approved_at TIMESTAMP NULL COMMENT '批准时间',
    rejected_at TIMESTAMP NULL COMMENT '拒绝时间',
    rejection_reason TEXT COMMENT '拒绝原因',
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    
    INDEX idx_user_id (user_id),
    INDEX idx_pet_id (pet_id),
    INDEX idx_org_id (organization_id),
    INDEX idx_status (status),
    INDEX idx_application_no (application_no),
    INDEX idx_created_at (created_at),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='领养申请表'`

	if err := db.Exec(createAppSQL).Error; err != nil {
		log.Fatal("创建adoption_applications表失败:", err)
	}
	fmt.Println("✓ adoption_applications表创建成功")

	// 5. 创建adoptions表
	fmt.Println("\n[3] 创建adoptions表...")
	// 先删除可能引用adoptions的表
	db.Exec("DROP TABLE IF EXISTS follow_up_records")
	fmt.Println("✓ 已删除follow_up_records表")
	
	createAdoptionSQL := `
CREATE TABLE IF NOT EXISTS adoptions (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '领养记录ID',
    application_id BIGINT UNSIGNED UNIQUE NOT NULL COMMENT '申请ID',
    user_id BIGINT NOT NULL COMMENT '领养人ID',
    pet_id BIGINT UNSIGNED NOT NULL COMMENT '宠物ID',
    organization_id BIGINT UNSIGNED NOT NULL COMMENT '机构ID',
    adoption_date DATE NOT NULL COMMENT '领养日期',
    handover_location VARCHAR(255) COMMENT '交接地点',
    agreement_url VARCHAR(255) COMMENT '协议文件URL',
    agreement_signed TINYINT DEFAULT 0 COMMENT '协议是否签署',
    follow_up_plan TEXT COMMENT '回访计划(JSON)',
    status VARCHAR(20) DEFAULT 'active' COMMENT '状态',
    notes TEXT COMMENT '备注',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    
    INDEX idx_user_id (user_id),
    INDEX idx_pet_id (pet_id),
    INDEX idx_org_id (organization_id),
    INDEX idx_adoption_date (adoption_date),
    INDEX idx_status (status),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='领养记录表'`

	if err := db.Exec(createAdoptionSQL).Error; err != nil {
		log.Fatal("创建adoptions表失败:", err)
	}
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	fmt.Println("✓ adoptions表创建成功")

	// 6. 验证
	fmt.Println("\n[4] 验证表结构...")
	var tables []string
	db.Raw("SHOW TABLES LIKE '%adoption%'").Scan(&tables)
	fmt.Println("✓ 找到以下表:")
	for _, table := range tables {
		fmt.Println("  -", table)
	}

	fmt.Println("\n✓ 领养模块数据库设置完成！")
	fmt.Println("现在可以测试领养API了")
}
