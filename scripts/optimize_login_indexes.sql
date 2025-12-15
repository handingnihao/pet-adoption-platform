-- 登录查询优化索引
-- 在 Navicat 中逐条执行（忽略DROP的错误，继续执行CREATE）

-- ========== 1. 查看现有索引 ==========
SHOW INDEX FROM users;
SHOW INDEX FROM posts;
SHOW INDEX FROM pets;

-- ========== 2. users 表优化 ==========
-- 如果 deleted_at 索引不存在，执行以下语句（已存在则跳过）
ALTER TABLE users ADD INDEX idx_users_deleted_at (deleted_at);

-- ========== 3. posts 表优化 ==========
ALTER TABLE posts ADD INDEX idx_posts_status_deleted (status, deleted_at);
ALTER TABLE posts ADD INDEX idx_posts_user_status (user_id, status, deleted_at);
ALTER TABLE posts ADD INDEX idx_posts_type_status (type, status, deleted_at);

-- ========== 4. pets 表优化 ==========
ALTER TABLE pets ADD INDEX idx_pets_status (status);
ALTER TABLE pets ADD INDEX idx_pets_type_status (type, status);
ALTER TABLE pets ADD INDEX idx_pets_user_status (user_id, status);

-- ========== 5. 更新统计信息 ==========
ANALYZE TABLE users;
ANALYZE TABLE posts;
ANALYZE TABLE pets;

-- 注意：如果提示"Duplicate key name"错误，说明索引已存在，可忽略该错误
