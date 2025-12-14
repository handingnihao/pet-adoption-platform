-- 创建所有测试用户
-- 包含多个密码选项供测试使用
USE pet_adoption;

-- 1. admin 用户（密码: Test@123456）
INSERT INTO users (username, password, phone, email, nickname, role, status, created_at, updated_at) 
VALUES (
    'admin',
    '$2a$10$1MVnDio.71vfRKPznJAdI.jE2046pebq6rBc0UYGZTt31/pJCv1GO',
    '13800138000',
    'admin@pet.com',
    '系统管理员',
    'admin',
    1,
    NOW(),
    NOW()
) ON DUPLICATE KEY UPDATE 
    password = '$2a$10$1MVnDio.71vfRKPznJAdI.jE2046pebq6rBc0UYGZTt31/pJCv1GO',
    nickname = '系统管理员',
    updated_at = NOW();

-- 2. testuser 用户（密码: 123456）
INSERT INTO users (username, password, phone, email, nickname, role, status, created_at, updated_at) 
VALUES (
    'testuser',
    '$2a$10$N9qo8uLOickgx2ZMRZoMye.nZ/HgkcGfHnwm9kYi0PL9bM/rwNLaC',
    '13900000001',
    'test@pet.com',
    '测试用户',
    'user',
    1,
    NOW(),
    NOW()
) ON DUPLICATE KEY UPDATE 
    password = '$2a$10$N9qo8uLOickgx2ZMRZoMye.nZ/HgkcGfHnwm9kYi0PL9bM/rwNLaC',
    updated_at = NOW();

-- 3. organization 机构用户（密码: 123456）
INSERT INTO users (username, password, phone, email, nickname, role, status, created_at, updated_at) 
VALUES (
    'org_user',
    '$2a$10$N9qo8uLOickgx2ZMRZoMye.nZ/HgkcGfHnwm9kYi0PL9bM/rwNLaC',
    '13900000002',
    'org@pet.com',
    '爱心机构',
    'organization',
    1,
    NOW(),
    NOW()
) ON DUPLICATE KEY UPDATE 
    password = '$2a$10$N9qo8uLOickgx2ZMRZoMye.nZ/HgkcGfHnwm9kYi0PL9bM/rwNLaC',
    updated_at = NOW();

-- 4. user01 普通用户（密码: Test@123456）
INSERT INTO users (username, password, phone, email, nickname, gender, role, status, created_at, updated_at) 
VALUES (
    'user01',
    '$2a$10$1MVnDio.71vfRKPznJAdI.jE2046pebq6rBc0UYGZTt31/pJCv1GO',
    '13900000003',
    'user01@pet.com',
    '爱宠人士01',
    'male',
    'user',
    1,
    NOW(),
    NOW()
) ON DUPLICATE KEY UPDATE 
    password = '$2a$10$1MVnDio.71vfRKPznJAdI.jE2046pebq6rBc0UYGZTt31/pJCv1GO',
    updated_at = NOW();

-- 验证创建结果
SELECT 
    id,
    username,
    nickname,
    phone,
    email,
    role,
    status,
    DATE_FORMAT(created_at, '%Y-%m-%d %H:%i') as created
FROM users 
WHERE username IN ('admin', 'testuser', 'org_user', 'user01')
ORDER BY id;

-- ===============================================
-- 测试账号信息汇总
-- ===============================================
-- 
-- 管理员账号:
--   用户名: admin
--   密码: Test@123456
--   角色: admin
--
-- 测试用户:
--   用户名: testuser
--   密码: 123456
--   角色: user
--
-- 机构用户:
--   用户名: org_user
--   密码: 123456
--   角色: organization
--
-- 普通用户:
--   用户名: user01
--   密码: Test@123456
--   角色: user
-- ===============================================
