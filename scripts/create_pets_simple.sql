-- 简化版pets表创建脚本（无外键约束）
USE pet_adoption;

-- 如果表存在则删除
DROP TABLE IF EXISTS pets;

-- 创建宠物表（无外键）
CREATE TABLE pets (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL COMMENT '宠物名称',
    type VARCHAR(20) NOT NULL COMMENT '宠物类型',
    breed VARCHAR(100) COMMENT '品种',
    gender VARCHAR(10) COMMENT '性别',
    age INT COMMENT '年龄(月)',
    size VARCHAR(20) COMMENT '体型',
    color VARCHAR(50) COMMENT '毛色',
    weight DECIMAL(5,2) COMMENT '体重(kg)',
    is_vaccinated BOOLEAN DEFAULT FALSE,
    is_sterilized BOOLEAN DEFAULT FALSE,
    health_status VARCHAR(200),
    description TEXT,
    `character` TEXT,
    photos TEXT,
    cover_photo VARCHAR(500),
    province VARCHAR(50),
    city VARCHAR(50),
    district VARCHAR(50),
    address VARCHAR(200),
    user_id BIGINT NOT NULL COMMENT '发布用户ID',
    status TINYINT DEFAULT 0,
    view_count INT DEFAULT 0,
    favorite_count INT DEFAULT 0,
    adopted_by BIGINT,
    adopted_at DATETIME,
    reviewed_by BIGINT,
    reviewed_at DATETIME,
    reject_reason VARCHAR(500),
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    INDEX idx_user_id (user_id),
    INDEX idx_type (type),
    INDEX idx_status (status),
    INDEX idx_city (city),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 插入测试数据
INSERT INTO pets (name, type, breed, gender, age, size, color, weight, is_vaccinated, is_sterilized, health_status, description, `character`, cover_photo, province, city, district, address, user_id, status) VALUES
('小白', 'dog', '金毛寻回犬', 'male', 18, 'large', '金黄色', 28.50, TRUE, FALSE, '健康', '温顺的金毛', '活泼', 'https://example.com/1.jpg', '广东省', '深圳市', '南山区', '科技园', 2, 1),
('咪咪', 'cat', '英短', 'female', 12, 'medium', '银渐层', 4.20, TRUE, TRUE, '健康', '可爱的猫咪', '安静', 'https://example.com/2.jpg', '广东省', '深圳市', '福田区', 'CBD', 2, 1);

SELECT * FROM pets;
