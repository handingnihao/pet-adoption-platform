-- ============================================
-- 修复 adoption_applications 表的 status 字段类型问题
-- 问题：GORM期望VARCHAR类型，但数据库可能是INT类型
-- ============================================

USE pet_adoption;

-- 步骤1: 查看当前表结构
DESCRIBE adoption_applications;

-- 步骤2: 查看现有数据
SELECT id, user_id, pet_id, status, applicant_name, created_at 
FROM adoption_applications 
ORDER BY created_at DESC;

-- 步骤3: 删除现有测试数据（如果需要）
-- DELETE FROM adoption_applications;

-- 步骤4: 修改status字段类型为VARCHAR(20)
ALTER TABLE adoption_applications 
MODIFY COLUMN status VARCHAR(20) NOT NULL DEFAULT 'pending' COMMENT '状态';

-- 步骤5: 验证修改后的表结构
DESCRIBE adoption_applications;

-- 步骤6: 确认数据
SELECT id, user_id, pet_id, status, applicant_name, created_at 
FROM adoption_applications 
ORDER BY created_at DESC;
