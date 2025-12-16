-- ============================================
-- 修复 adoption_applications 表的 status 字段
-- 问题1：字段类型是tinyint(1)会被解释为布尔值
-- 问题2：存储的是字符串值而非整数
-- 解决：修改字段类型为TINYINT（不带括号），转换数据为整数
-- ============================================

USE pet_adoption;

-- 步骤1: 查看当前表结构
DESCRIBE adoption_applications;

-- 步骤2: 查看当前数据
SELECT id, user_id, pet_id, status, applicant_name, created_at 
FROM adoption_applications 
ORDER BY created_at DESC;

-- 步骤3: 修改字段类型为TINYINT（避免TINYINT(1)被解释为布尔值）
ALTER TABLE adoption_applications 
MODIFY COLUMN status TINYINT NOT NULL DEFAULT 0 COMMENT '状态';

-- 步骤4: 将字符串状态转换为整数
-- pending -> 0
UPDATE adoption_applications SET status = 0 WHERE status = 'pending' OR status = '0';

-- reviewing -> 1
UPDATE adoption_applications SET status = 1 WHERE status = 'reviewing' OR status = '1';

-- interview -> 2
UPDATE adoption_applications SET status = 2 WHERE status = 'interview' OR status = '2';

-- home_visit -> 3
UPDATE adoption_applications SET status = 3 WHERE status = 'home_visit' OR status = '3';

-- approved -> 4
UPDATE adoption_applications SET status = 4 WHERE status = 'approved' OR status = '4';

-- rejected -> 5
UPDATE adoption_applications SET status = 5 WHERE status = 'rejected' OR status = '5';

-- cancelled -> 6
UPDATE adoption_applications SET status = 6 WHERE status = 'cancelled' OR status = '6';

-- 步骤5: 验证转换结果
SELECT id, user_id, pet_id, status, applicant_name, created_at 
FROM adoption_applications 
ORDER BY created_at DESC;

-- 步骤6: 确认所有status都是数字
SELECT DISTINCT status FROM adoption_applications;

-- 步骤7: 再次查看表结构，确认类型正确
DESCRIBE adoption_applications;
