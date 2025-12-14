-- ============================================
-- 爱心宠物领养平台 - 数据库初始化脚本
-- MySQL 8.4
-- 步骤1: 创建用户和数据库
-- ============================================

-- 注意：此脚本需要使用 root 用户执行
-- 使用方法: mysql -u root -p < 01_create_user_and_database.sql

-- 1. 创建数据库
DROP DATABASE IF EXISTS pet_adoption;
CREATE DATABASE pet_adoption 
    CHARACTER SET utf8mb4 
    COLLATE utf8mb4_unicode_ci
    COMMENT '爱心宠物领养平台数据库';

-- 2. 创建专用应用用户（生产环境建议）
-- 注意：请修改密码为强密码
DROP USER IF EXISTS 'pet_admin'@'%';
DROP USER IF EXISTS 'pet_admin'@'localhost';

-- 创建用户并设置密码（请修改为强密码）
CREATE USER 'pet_admin'@'%' IDENTIFIED BY 'Pet@Admin#2024!Secure';
CREATE USER 'pet_admin'@'localhost' IDENTIFIED BY 'Pet@Admin#2024!Secure';

-- MySQL 8.4 密码策略（可选）
-- ALTER USER 'pet_admin'@'%' PASSWORD EXPIRE INTERVAL 90 DAY;

-- 3. 授权
-- 授予数据库完全权限
GRANT ALL PRIVILEGES ON pet_adoption.* TO 'pet_admin'@'%';
GRANT ALL PRIVILEGES ON pet_adoption.* TO 'pet_admin'@'localhost';

-- 如果只需要基本权限（推荐）
-- GRANT SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, INDEX, ALTER, CREATE TEMPORARY TABLES, LOCK TABLES 
-- ON pet_adoption.* TO 'pet_admin'@'%';

-- 4. 创建只读用户（用于数据分析、报表等）
DROP USER IF EXISTS 'pet_readonly'@'%';
CREATE USER 'pet_readonly'@'%' IDENTIFIED BY 'Pet@ReadOnly#2024';
GRANT SELECT ON pet_adoption.* TO 'pet_readonly'@'%';

-- 5. 刷新权限
FLUSH PRIVILEGES;

-- 6. 切换到新建的数据库
USE pet_adoption;

-- 7. 验证数据库创建成功
SELECT 
    SCHEMA_NAME AS '数据库名',
    DEFAULT_CHARACTER_SET_NAME AS '字符集',
    DEFAULT_COLLATION_NAME AS '排序规则'
FROM information_schema.SCHEMATA 
WHERE SCHEMA_NAME = 'pet_adoption';

-- 8. 显示用户权限
SHOW GRANTS FOR 'pet_admin'@'%';
SHOW GRANTS FOR 'pet_admin'@'localhost';
SHOW GRANTS FOR 'pet_readonly'@'%';

-- ============================================
-- 创建完成提示
-- ============================================
SELECT '✅ 数据库和用户创建成功！' AS Status;
SELECT '📝 请记录以下连接信息：' AS Notice;
SELECT 
    'pet_admin' AS '用户名',
    'Pet@Admin#2024!Secure' AS '密码（请修改）',
    'pet_adoption' AS '数据库名',
    '3306' AS '端口',
    'utf8mb4' AS '字符集';
