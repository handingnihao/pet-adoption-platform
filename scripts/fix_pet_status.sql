-- ============================================
-- 修复 pets 表的 status 字段数据
-- 问题：status字段存储的是字符串值，需要转换为整数
-- ============================================

USE pet_adoption;

-- 步骤1: 查看当前pets表结构
DESCRIBE pets;

-- 步骤2: 查看当前数据
SELECT id, name, status FROM pets;

-- 步骤3: 修改字段类型为TINYINT（如果需要）
ALTER TABLE pets MODIFY COLUMN status TINYINT NOT NULL DEFAULT 0 COMMENT '状态';

-- 步骤4: 将字符串状态转换为整数
-- pending -> 0
UPDATE pets SET status = 0 WHERE status = 'pending' OR status = '0';

-- available -> 1
UPDATE pets SET status = 1 WHERE status = 'available' OR status = '1';

-- adopted/approved -> 2
UPDATE pets SET status = 2 WHERE status = 'adopted' OR status = 'approved' OR status = '2';

-- offline -> 3
UPDATE pets SET status = 3 WHERE status = 'offline' OR status = '3';

-- 步骤5: 验证转换结果
SELECT id, name, status FROM pets;

-- 步骤6: 确认所有status都是数字
SELECT DISTINCT status FROM pets;
