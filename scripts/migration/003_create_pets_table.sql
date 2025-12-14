-- 创建宠物表
-- 执行时间：2025-12-09
-- 描述：宠物领养平台核心表 - 宠物信息表

USE pet_adoption;

-- 创建宠物表
CREATE TABLE IF NOT EXISTS pets (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '宠物ID',
    
    -- 基本信息
    name VARCHAR(100) NOT NULL COMMENT '宠物名称',
    type VARCHAR(20) NOT NULL COMMENT '宠物类型(dog/cat/rabbit/bird/other)',
    breed VARCHAR(100) COMMENT '品种',
    gender VARCHAR(10) COMMENT '性别(male/female/unknown)',
    age INT COMMENT '年龄(月)',
    size VARCHAR(20) COMMENT '体型(small/medium/large)',
    color VARCHAR(50) COMMENT '毛色',
    weight DECIMAL(5,2) COMMENT '体重(kg)',
    
    -- 健康信息
    is_vaccinated BOOLEAN DEFAULT FALSE COMMENT '是否接种疫苗',
    is_sterilized BOOLEAN DEFAULT FALSE COMMENT '是否绝育',
    health_status VARCHAR(200) COMMENT '健康状况',
    
    -- 详细信息
    description TEXT COMMENT '详细描述',
    `character` TEXT COMMENT '性格特点',
    photos TEXT COMMENT '照片URL,逗号分隔',
    cover_photo VARCHAR(500) COMMENT '封面图',
    
    -- 位置信息
    province VARCHAR(50) COMMENT '省份',
    city VARCHAR(50) COMMENT '城市',
    district VARCHAR(50) COMMENT '区县',
    address VARCHAR(200) COMMENT '详细地址',
    
    -- 关联信息
    user_id INT UNSIGNED NOT NULL COMMENT '发布用户ID',
    
    -- 状态和统计
    status TINYINT DEFAULT 0 COMMENT '状态(0待审核/1可领养/2已领养/3已下架)',
    view_count INT DEFAULT 0 COMMENT '浏览次数',
    favorite_count INT DEFAULT 0 COMMENT '收藏次数',
    
    -- 领养信息
    adopted_by INT UNSIGNED COMMENT '领养人ID',
    adopted_at DATETIME COMMENT '领养时间',
    
    -- 审核信息
    reviewed_by INT UNSIGNED COMMENT '审核人ID',
    reviewed_at DATETIME COMMENT '审核时间',
    reject_reason VARCHAR(500) COMMENT '拒绝原因',
    
    -- 时间戳
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME COMMENT '删除时间(软删除)',
    
    -- 索引
    INDEX idx_user_id (user_id),
    INDEX idx_type (type),
    INDEX idx_status (status),
    INDEX idx_city (city),
    INDEX idx_adopted_by (adopted_by),
    INDEX idx_created_at (created_at),
    INDEX idx_deleted_at (deleted_at),
    
    -- 外键约束
    CONSTRAINT fk_pets_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_pets_adopted_by FOREIGN KEY (adopted_by) REFERENCES users(id) ON DELETE SET NULL,
    CONSTRAINT fk_pets_reviewed_by FOREIGN KEY (reviewed_by) REFERENCES users(id) ON DELETE SET NULL
    
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='宠物信息表';

-- 创建全文索引用于搜索
ALTER TABLE pets ADD FULLTEXT INDEX idx_fulltext_search (name, breed, description) WITH PARSER ngram;

-- 插入测试数据
INSERT INTO pets (
    name, type, breed, gender, age, size, color, weight,
    is_vaccinated, is_sterilized, health_status,
    description, `character`, cover_photo,
    province, city, district, address,
    user_id, status
) VALUES
(
    '小白',
    'dog',
    '金毛寻回犬',
    'male',
    18,
    'large',
    '金黄色',
    28.50,
    TRUE,
    FALSE,
    '健康状况良好，已完成所有疫苗接种',
    '小白是一只温顺可爱的金毛，非常喜欢和人玩耍，性格活泼开朗。适合有院子的家庭饲养。',
    '温顺、活泼、聪明、忠诚',
    'https://example.com/pets/golden-retriever.jpg',
    '广东省',
    '深圳市',
    '南山区',
    '科技园',
    2,
    1
),
(
    '咪咪',
    'cat',
    '英国短毛猫',
    'female',
    12,
    'medium',
    '银渐层',
    4.20,
    TRUE,
    TRUE,
    '健康，已绝育，已接种疫苗',
    '咪咪是一只可爱的银渐层英短，性格温和，喜欢安静的环境。非常适合公寓饲养。',
    '安静、温和、独立、爱干净',
    'https://example.com/pets/british-shorthair.jpg',
    '广东省',
    '深圳市',
    '福田区',
    'CBD',
    4,
    1
),
(
    '旺财',
    'dog',
    '中华田园犬',
    'male',
    24,
    'medium',
    '黄白相间',
    15.00,
    TRUE,
    TRUE,
    '健康活泼',
    '旺财是一只非常忠诚的田园犬，看家护院的好帮手。已经两岁了，非常懂事。',
    '忠诚、勇敢、机警、看家好手',
    'https://example.com/pets/chinese-rural-dog.jpg',
    '广东省',
    '深圳市',
    '龙岗区',
    '布吉',
    2,
    1
),
(
    '雪球',
    'rabbit',
    '安哥拉兔',
    'female',
    6,
    'small',
    '纯白色',
    2.50,
    TRUE,
    FALSE,
    '健康',
    '雪球是一只纯白色的安哥拉兔，毛发柔软，非常可爱。性格温顺，适合陪伴。',
    '温顺、可爱、安静',
    'https://example.com/pets/angora-rabbit.jpg',
    '广东省',
    '深圳市',
    '宝安区',
    '西乡',
    4,
    1
),
(
    '虎妞',
    'cat',
    '橘猫',
    'female',
    8,
    'medium',
    '橘色',
    5.80,
    TRUE,
    FALSE,
    '健康，食欲旺盛',
    '虎妞是一只可爱的橘猫，性格活泼好动，喜欢吃东西。需要主人有耐心陪伴。',
    '活泼、贪吃、粘人',
    'https://example.com/pets/orange-cat.jpg',
    '广东省',
    '深圳市',
    '南山区',
    '后海',
    2,
    0
);

-- 查看插入结果
SELECT 
    id,
    name,
    type,
    breed,
    gender,
    age,
    city,
    status,
    DATE_FORMAT(created_at, '%Y-%m-%d %H:%i') as created
FROM pets
ORDER BY id;

-- 统计信息
SELECT 
    '总宠物数' as item,
    COUNT(*) as count
FROM pets
UNION ALL
SELECT 
    '待审核',
    COUNT(*) 
FROM pets 
WHERE status = 0
UNION ALL
SELECT 
    '可领养',
    COUNT(*) 
FROM pets 
WHERE status = 1
UNION ALL
SELECT 
    '已领养',
    COUNT(*) 
FROM pets 
WHERE status = 2;
