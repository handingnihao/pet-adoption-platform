-- ============================================
-- 爱心宠物领养平台 - 数据表创建脚本
-- MySQL 8.4
-- 步骤2: 创建所有数据表
-- ============================================

-- 使用方法: mysql -u pet_admin -p pet_adoption < 02_create_tables.sql
-- 或在MySQL中: source /path/to/02_create_tables.sql

USE pet_adoption;

-- 设置存储引擎和字符集
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ============================================
-- 1. 用户表 (users)
-- ============================================
DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '用户ID',
    username VARCHAR(50) UNIQUE NOT NULL COMMENT '用户名',
    password VARCHAR(255) NOT NULL COMMENT '密码(BCrypt加密)',
    real_name VARCHAR(50) COMMENT '真实姓名',
    phone VARCHAR(20) UNIQUE COMMENT '手机号',
    email VARCHAR(100) UNIQUE COMMENT '邮箱',
    avatar VARCHAR(255) COMMENT '头像URL',
    gender TINYINT COMMENT '性别: 0-未知, 1-男, 2-女',
    birthday DATE COMMENT '生日',
    id_card VARCHAR(18) COMMENT '身份证号',
    address TEXT COMMENT '地址',
    role ENUM('user', 'organization', 'volunteer', 'admin') DEFAULT 'user' COMMENT '角色',
    status TINYINT DEFAULT 1 COMMENT '状态: 0-禁用, 1-正常',
    last_login_at TIMESTAMP NULL COMMENT '最后登录时间',
    last_login_ip VARCHAR(50) COMMENT '最后登录IP',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL COMMENT '删除时间(软删除)',
    
    INDEX idx_username (username),
    INDEX idx_phone (phone),
    INDEX idx_email (email),
    INDEX idx_role (role),
    INDEX idx_status (status),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- ============================================
-- 2. 救助机构表 (organizations)
-- ============================================
DROP TABLE IF EXISTS organizations;
CREATE TABLE organizations (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '机构ID',
    name VARCHAR(100) NOT NULL COMMENT '机构名称',
    license_number VARCHAR(50) UNIQUE COMMENT '营业执照号',
    legal_person VARCHAR(50) COMMENT '法人',
    contact_person VARCHAR(50) COMMENT '联系人',
    contact_phone VARCHAR(20) COMMENT '联系电话',
    email VARCHAR(100) COMMENT '邮箱',
    address TEXT COMMENT '地址',
    description TEXT COMMENT '机构介绍',
    logo VARCHAR(255) COMMENT 'Logo URL',
    images TEXT COMMENT '机构图片(JSON数组)',
    license_image VARCHAR(255) COMMENT '营业执照图片URL',
    status ENUM('pending', 'approved', 'rejected') DEFAULT 'pending' COMMENT '审核状态',
    verified TINYINT DEFAULT 0 COMMENT '是否认证: 0-未认证, 1-已认证',
    user_id BIGINT NOT NULL COMMENT '关联用户ID',
    rating DECIMAL(3,2) DEFAULT 5.00 COMMENT '评分(1.00-5.00)',
    adoption_count INT DEFAULT 0 COMMENT '领养成功数',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    
    INDEX idx_user_id (user_id),
    INDEX idx_status (status),
    INDEX idx_verified (verified),
    INDEX idx_license (license_number),
    CONSTRAINT fk_org_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='救助机构表';

-- ============================================
-- 3. 宠物表 (pets)
-- ============================================
DROP TABLE IF EXISTS pets;
CREATE TABLE pets (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '宠物ID',
    name VARCHAR(50) NOT NULL COMMENT '宠物名称',
    species ENUM('dog', 'cat', 'other') NOT NULL COMMENT '物种: dog-狗, cat-猫, other-其他',
    breed VARCHAR(50) COMMENT '品种',
    gender ENUM('male', 'female', 'unknown') COMMENT '性别',
    age_months INT COMMENT '年龄(月)',
    weight DECIMAL(5,2) COMMENT '体重(kg)',
    color VARCHAR(50) COMMENT '颜色',
    description TEXT COMMENT '描述',
    `character` TEXT COMMENT '性格特点',
    health_status TEXT COMMENT '健康状况',
    is_neutered TINYINT DEFAULT 0 COMMENT '是否绝育: 0-否, 1-是',
    vaccination_record TEXT COMMENT '疫苗记录(JSON)',
    main_image VARCHAR(255) COMMENT '主图URL',
    images TEXT COMMENT '图片集(JSON数组)',
    video_url VARCHAR(255) COMMENT '视频URL',
    organization_id BIGINT NOT NULL COMMENT '所属机构ID',
    location VARCHAR(100) COMMENT '所在地',
    status ENUM('available', 'reserved', 'adopted', 'fostered') DEFAULT 'available' COMMENT '状态',
    view_count INT DEFAULT 0 COMMENT '浏览次数',
    like_count INT DEFAULT 0 COMMENT '点赞数',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    
    INDEX idx_organization_id (organization_id),
    INDEX idx_species_status (species, status),
    INDEX idx_status (status),
    INDEX idx_location (location),
    INDEX idx_created_at (created_at),
    INDEX idx_view_count (view_count),
    CONSTRAINT fk_pet_org FOREIGN KEY (organization_id) REFERENCES organizations(id) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='宠物表';

-- ============================================
-- 4. 领养申请表 (adoption_applications)
-- ============================================
DROP TABLE IF EXISTS adoption_applications;
CREATE TABLE adoption_applications (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '申请ID',
    application_no VARCHAR(50) UNIQUE NOT NULL COMMENT '申请编号',
    user_id BIGINT NOT NULL COMMENT '申请人ID',
    pet_id BIGINT NOT NULL COMMENT '宠物ID',
    organization_id BIGINT NOT NULL COMMENT '机构ID',
    
    -- 申请人信息
    applicant_name VARCHAR(50) NOT NULL COMMENT '申请人姓名',
    applicant_phone VARCHAR(20) NOT NULL COMMENT '申请人电话',
    applicant_id_card VARCHAR(18) COMMENT '身份证号',
    applicant_address TEXT COMMENT '地址',
    
    -- 家庭情况
    housing_type ENUM('apartment', 'house', 'villa', 'other') COMMENT '住房类型',
    housing_area INT COMMENT '住房面积(平方米)',
    has_yard TINYINT DEFAULT 0 COMMENT '是否有院子',
    family_members INT COMMENT '家庭成员数',
    has_children TINYINT DEFAULT 0 COMMENT '是否有孩子',
    children_age VARCHAR(50) COMMENT '孩子年龄',
    family_agree TINYINT DEFAULT 0 COMMENT '家人是否同意',
    
    -- 养宠经验
    has_pet_experience TINYINT DEFAULT 0 COMMENT '是否有养宠经验',
    pet_experience TEXT COMMENT '养宠经历',
    current_pets TEXT COMMENT '现有宠物',
    
    -- 领养原因
    adoption_reason TEXT COMMENT '领养原因',
    how_to_care TEXT COMMENT '如何照顾',
    emergency_plan TEXT COMMENT '应急计划',
    
    -- 附件
    id_card_image VARCHAR(255) COMMENT '身份证照片URL',
    housing_proof VARCHAR(255) COMMENT '住房证明URL',
    additional_files TEXT COMMENT '其他附件(JSON)',
    
    -- 审核流程
    status ENUM('pending', 'reviewing', 'interview', 'home_visit', 'approved', 'rejected', 'cancelled') 
        DEFAULT 'pending' COMMENT '状态',
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
    UNIQUE KEY uk_user_pet (user_id, pet_id, created_at),
    CONSTRAINT fk_app_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT,
    CONSTRAINT fk_app_pet FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE RESTRICT,
    CONSTRAINT fk_app_org FOREIGN KEY (organization_id) REFERENCES organizations(id) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='领养申请表';

-- ============================================
-- 5. 领养记录表 (adoptions)
-- ============================================
DROP TABLE IF EXISTS adoptions;
CREATE TABLE adoptions (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '领养记录ID',
    application_id BIGINT UNIQUE NOT NULL COMMENT '申请ID',
    user_id BIGINT NOT NULL COMMENT '领养人ID',
    pet_id BIGINT NOT NULL COMMENT '宠物ID',
    organization_id BIGINT NOT NULL COMMENT '机构ID',
    adoption_date DATE NOT NULL COMMENT '领养日期',
    handover_location VARCHAR(255) COMMENT '交接地点',
    agreement_url VARCHAR(255) COMMENT '协议文件URL',
    agreement_signed TINYINT DEFAULT 0 COMMENT '协议是否签署',
    follow_up_plan TEXT COMMENT '回访计划(JSON)',
    status ENUM('active', 'returned', 'deceased') DEFAULT 'active' COMMENT '状态',
    notes TEXT COMMENT '备注',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    
    INDEX idx_user_id (user_id),
    INDEX idx_pet_id (pet_id),
    INDEX idx_org_id (organization_id),
    INDEX idx_adoption_date (adoption_date),
    INDEX idx_status (status),
    CONSTRAINT fk_adoption_app FOREIGN KEY (application_id) REFERENCES adoption_applications(id) ON DELETE RESTRICT,
    CONSTRAINT fk_adoption_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT,
    CONSTRAINT fk_adoption_pet FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE RESTRICT,
    CONSTRAINT fk_adoption_org FOREIGN KEY (organization_id) REFERENCES organizations(id) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='领养记录表';

-- ============================================
-- 6. 回访记录表 (follow_up_records)
-- ============================================
DROP TABLE IF EXISTS follow_up_records;
CREATE TABLE follow_up_records (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '回访记录ID',
    adoption_id BIGINT NOT NULL COMMENT '领养记录ID',
    follow_up_date DATE NOT NULL COMMENT '回访日期',
    follow_up_type ENUM('7days', '30days', '90days', '180days', 'custom') COMMENT '回访类型',
    visitor_id BIGINT COMMENT '回访人ID',
    pet_status TEXT COMMENT '宠物状况',
    living_environment TEXT COMMENT '生活环境',
    health_condition TEXT COMMENT '健康状况',
    images TEXT COMMENT '照片(JSON)',
    video_url VARCHAR(255) COMMENT '视频URL',
    issues TEXT COMMENT '发现的问题',
    suggestions TEXT COMMENT '建议',
    rating TINYINT COMMENT '评分(1-5)',
    next_follow_up_date DATE COMMENT '下次回访日期',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    
    INDEX idx_adoption_id (adoption_id),
    INDEX idx_follow_up_date (follow_up_date),
    INDEX idx_visitor_id (visitor_id),
    CONSTRAINT fk_followup_adoption FOREIGN KEY (adoption_id) REFERENCES adoptions(id) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='回访记录表';

-- ============================================
-- 7. 社区动态表 (posts)
-- ============================================
DROP TABLE IF EXISTS posts;
CREATE TABLE posts (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '动态ID',
    user_id BIGINT NOT NULL COMMENT '发布者ID',
    title VARCHAR(200) COMMENT '标题',
    content TEXT NOT NULL COMMENT '内容',
    images TEXT COMMENT '图片(JSON)',
    video_url VARCHAR(255) COMMENT '视频URL',
    topic_ids TEXT COMMENT '话题ID(JSON)',
    pet_id BIGINT COMMENT '关联宠物ID',
    type ENUM('story', 'knowledge', 'daily', 'other') DEFAULT 'daily' COMMENT '类型',
    view_count INT DEFAULT 0 COMMENT '浏览次数',
    like_count INT DEFAULT 0 COMMENT '点赞数',
    comment_count INT DEFAULT 0 COMMENT '评论数',
    share_count INT DEFAULT 0 COMMENT '分享数',
    status TINYINT DEFAULT 1 COMMENT '状态: 0-隐藏, 1-正常',
    is_top TINYINT DEFAULT 0 COMMENT '是否置顶',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    
    INDEX idx_user_id (user_id),
    INDEX idx_created_at (created_at),
    INDEX idx_status (status),
    INDEX idx_type (type),
    INDEX idx_is_top (is_top),
    FULLTEXT INDEX ft_title_content (title, content) WITH PARSER ngram,
    CONSTRAINT fk_post_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='社区动态表';

-- ============================================
-- 8. 评论表 (comments)
-- ============================================
DROP TABLE IF EXISTS comments;
CREATE TABLE comments (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '评论ID',
    post_id BIGINT NOT NULL COMMENT '动态ID',
    user_id BIGINT NOT NULL COMMENT '评论者ID',
    parent_id BIGINT DEFAULT 0 COMMENT '父评论ID, 0为一级评论',
    reply_to_user_id BIGINT COMMENT '回复的用户ID',
    content TEXT NOT NULL COMMENT '评论内容',
    like_count INT DEFAULT 0 COMMENT '点赞数',
    status TINYINT DEFAULT 1 COMMENT '状态: 0-隐藏, 1-正常',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    
    INDEX idx_post_id (post_id),
    INDEX idx_user_id (user_id),
    INDEX idx_parent_id (parent_id),
    INDEX idx_created_at (created_at),
    CONSTRAINT fk_comment_post FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    CONSTRAINT fk_comment_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='评论表';

-- ============================================
-- 9. 捐赠记录表 (donations)
-- ============================================
DROP TABLE IF EXISTS donations;
CREATE TABLE donations (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '捐赠ID',
    donation_no VARCHAR(50) UNIQUE NOT NULL COMMENT '捐赠单号',
    user_id BIGINT COMMENT '捐赠人ID',
    organization_id BIGINT COMMENT '受捐机构ID',
    pet_id BIGINT COMMENT '定向宠物ID',
    type ENUM('money', 'goods') NOT NULL COMMENT '捐赠类型',
    amount DECIMAL(10,2) COMMENT '金额',
    goods_description TEXT COMMENT '物资描述',
    payment_method ENUM('wechat', 'alipay', 'bank') COMMENT '支付方式',
    transaction_id VARCHAR(100) COMMENT '交易流水号',
    is_anonymous TINYINT DEFAULT 0 COMMENT '是否匿名',
    message TEXT COMMENT '留言',
    status ENUM('pending', 'paid', 'failed', 'refunded') DEFAULT 'pending' COMMENT '状态',
    receipt_url VARCHAR(255) COMMENT '电子收据URL',
    paid_at TIMESTAMP NULL COMMENT '支付时间',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    
    INDEX idx_user_id (user_id),
    INDEX idx_org_id (organization_id),
    INDEX idx_donation_no (donation_no),
    INDEX idx_status (status),
    INDEX idx_created_at (created_at),
    CONSTRAINT fk_donation_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT,
    CONSTRAINT fk_donation_org FOREIGN KEY (organization_id) REFERENCES organizations(id) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='捐赠记录表';

-- ============================================
-- 10. 消息通知表 (notifications)
-- ============================================
DROP TABLE IF EXISTS notifications;
CREATE TABLE notifications (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '通知ID',
    user_id BIGINT NOT NULL COMMENT '接收用户ID',
    type ENUM('system', 'adoption', 'follow_up', 'comment', 'like', 'donation') NOT NULL COMMENT '类型',
    title VARCHAR(200) NOT NULL COMMENT '标题',
    content TEXT COMMENT '内容',
    related_id BIGINT COMMENT '关联ID',
    related_type VARCHAR(50) COMMENT '关联类型',
    is_read TINYINT DEFAULT 0 COMMENT '是否已读: 0-未读, 1-已读',
    read_at TIMESTAMP NULL COMMENT '阅读时间',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    
    INDEX idx_user_read (user_id, is_read),
    INDEX idx_created_at (created_at),
    INDEX idx_type (type),
    CONSTRAINT fk_notification_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='消息通知表';

-- ============================================
-- 恢复外键检查
-- ============================================
SET FOREIGN_KEY_CHECKS = 1;

-- ============================================
-- 显示所有创建的表
-- ============================================
SELECT '✅ 所有数据表创建成功！' AS Status;
SHOW TABLES;

-- 显示表结构统计
SELECT 
    TABLE_NAME AS '表名',
    TABLE_ROWS AS '预估行数',
    AVG_ROW_LENGTH AS '平均行长度',
    TABLE_COMMENT AS '表说明'
FROM information_schema.TABLES
WHERE TABLE_SCHEMA = 'pet_adoption'
ORDER BY TABLE_NAME;
