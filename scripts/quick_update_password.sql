-- 快速更新所有测试用户密码为 Test@123456
USE pet_adoption;

-- 生成的bcrypt哈希值 (Test@123456)
SET @password_hash = '$2a$10$1MVnDio.71vfRKPznJAdI.jE2046pebq6rBc0UYGZTt31/pJCv1GO';

-- 更新admin
UPDATE users SET password = @password_hash WHERE username = 'admin';

-- 更新testuser1
UPDATE users SET password = @password_hash WHERE username = 'testuser1';

-- 更新testuser2
UPDATE users SET password = @password_hash WHERE username = 'testuser2';

-- 更新testorg
UPDATE users SET password = @password_hash WHERE username = 'testorg';

-- 验证更新
SELECT username, role, LEFT(password, 30) as password_hash FROM users WHERE username IN ('admin', 'testuser1', 'testuser2', 'testorg');

-- 所有用户密码已更新为: Test@123456
