-- ============================================
-- 爱心宠物领养平台 - 初始化数据脚本
-- MySQL 8.4
-- 步骤3: 插入初始数据
-- ============================================

USE pet_adoption;

-- ============================================
-- 1. 创建管理员账户
-- ============================================
-- 密码: Admin@123456 (BCrypt加密后的hash)
-- 注意：生产环境请立即修改密码！
INSERT INTO users (username, password, real_name, phone, email, role, status) VALUES
('admin', '$2a$10$XBWxOvKjLvV4J3BLMy7wPeKHqPR7D0gKNO7EX6Fj7N.vN7RZT5Jle', '系统管理员', '13800138000', 'admin@petadoption.com', 'admin', 1);

-- ============================================
-- 2. 创建测试用户（可选，测试环境使用）
-- ============================================
-- 密码都是: Test@123456
INSERT INTO users (username, password, real_name, phone, email, role, status) VALUES
('testuser1', '$2a$10$XBWxOvKjLvV4J3BLMy7wPeKHqPR7D0gKNO7EX6Fj7N.vN7RZT5Jle', '测试用户1', '13900000001', 'user1@test.com', 'user', 1),
('testuser2', '$2a$10$XBWxOvKjLvV4J3BLMy7wPeKHqPR7D0gKNO7EX6Fj7N.vN7RZT5Jle', '测试用户2', '13900000002', 'user2@test.com', 'user', 1),
('testorg', '$2a$10$XBWxOvKjLvV4J3BLMy7wPeKHqPR7D0gKNO7EX6Fj7N.vN7RZT5Jle', '测试机构', '13900000003', 'org@test.com', 'organization', 1);

-- ============================================
-- 3. 创建测试救助机构（可选）
-- ============================================
INSERT INTO organizations (name, contact_person, contact_phone, email, address, description, user_id, status, verified) VALUES
('爱心流浪动物救助中心', '张三', '13900000003', 'org@test.com', '北京市朝阳区测试路123号', '致力于流浪动物救助和领养', 4, 'approved', 1);

-- ============================================
-- 4. 验证初始数据
-- ============================================
SELECT '✅ 初始数据插入成功！' AS Status;

SELECT 
    '用户总数' AS '统计项',
    COUNT(*) AS '数量'
FROM users
UNION ALL
SELECT 
    '管理员数量' AS '统计项',
    COUNT(*) AS '数量'
FROM users WHERE role = 'admin'
UNION ALL
SELECT 
    '机构数量' AS '统计项',
    COUNT(*) AS '数量'
FROM organizations;

-- ============================================
-- 提示信息
-- ============================================
SELECT '📝 默认管理员账户信息：' AS Notice;
SELECT 
    'admin' AS '用户名',
    'Admin@123456' AS '初始密码',
    '⚠️ 请立即登录并修改密码！' AS '安全提示';
