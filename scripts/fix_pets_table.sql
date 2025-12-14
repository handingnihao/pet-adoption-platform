-- 修改pets表的字段类型以匹配Go模型
USE pet_adoption;

-- 修改user_id为BIGINT
ALTER TABLE pets MODIFY COLUMN user_id BIGINT NOT NULL COMMENT '发布用户ID';

-- 修改adopted_by为BIGINT
ALTER TABLE pets MODIFY COLUMN adopted_by BIGINT COMMENT '领养人ID';

-- 修改reviewed_by为BIGINT  
ALTER TABLE pets MODIFY COLUMN reviewed_by BIGINT COMMENT '审核人ID';

-- 查看修改后的表结构
DESC pets;

-- 插入测试数据
INSERT INTO pets (
    name, type, breed, gender, age, size, color, weight,
    is_vaccinated, is_sterilized, health_status,
    description, `character`, cover_photo,
    province, city, district, address,
    user_id, status
) VALUES
('小白', 'dog', '金毛寻回犬', 'male', 18, 'large', '金黄色', 28.50,
 TRUE, FALSE, '健康状况良好',
 '小白是一只温顺可爱的金毛。', '温顺、活泼', 'https://example.com/golden.jpg',
 '广东省', '深圳市', '南山区', '科技园', 2, 1),
('咪咪', 'cat', '英国短毛猫', 'female', 12, 'medium', '银渐层', 4.20,
 TRUE, TRUE, '健康',
 '咪咪是一只可爱的银渐层英短。', '安静、温和', 'https://example.com/cat.jpg',
 '广东省', '深圳市', '福田区', 'CBD', 2, 1);

-- 查询结果
SELECT id, name, type, breed, city, status FROM pets;
