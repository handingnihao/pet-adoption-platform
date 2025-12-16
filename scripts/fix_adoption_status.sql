-- ============================================
-- 修复 adoption_applications 表的 status 字段数据
-- 问题：字段类型是tinyint，但存储的是字符串值
-- 解决：将字符串值转换为对应的整数值
-- ============================================

USE pet_adoption;

-- 步骤1: 查看当前数据
SELECT id, user_id, pet_id, status, applicant_name, created_at 
FROM adoption_applications 
ORDER BY created_at DESC;

-- 步骤2: 将字符串状态转换为整数
-- pending -> 0
UPDATE adoption_applications SET status = 0 WHERE status = 'pending';

-- reviewing -> 1
UPDATE adoption_applications SET status = 1 WHERE status = 'reviewing';

-- interview -> 2
UPDATE adoption_applications SET status = 2 WHERE status = 'interview';

-- home_visit -> 3
UPDATE adoption_applications SET status = 3 WHERE status = 'home_visit';

-- approved -> 4
UPDATE adoption_applications SET status = 4 WHERE status = 'approved';

-- rejected -> 5
UPDATE adoption_applications SET status = 5 WHERE status = 'rejected';

-- cancelled -> 6
UPDATE adoption_applications SET status = 6 WHERE status = 'cancelled';

-- 步骤3: 验证转换结果
SELECT id, user_id, pet_id, status, applicant_name, created_at 
FROM adoption_applications 
ORDER BY created_at DESC;

-- 步骤4: 确认所有status都是数字
SELECT DISTINCT status FROM adoption_applications;
