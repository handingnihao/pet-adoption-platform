# 数据库迁移脚本说明

## 📋 文件清单

### SQL 脚本
- `01_create_user_and_database.sql` - 创建数据库和用户
- `02_create_tables.sql` - 创建所有数据表
- `03_init_data.sql` - 插入初始数据

### 自动化脚本
- `../deploy/deploy_database.sh` - Linux/Mac 一键部署脚本
- `../deploy/deploy_database.bat` - Windows 一键部署脚本

### 文档
- `数据库部署指南.md` - 详细部署文档

## 🚀 快速开始

### 方式1: 使用自动化脚本（推荐）

#### Linux/Mac
```bash
cd scripts/deploy
chmod +x deploy_database.sh
./deploy_database.sh
```

#### Windows
```bash
cd scripts\deploy
deploy_database.bat
```

### 方式2: 手动执行 SQL

#### 步骤1: 创建数据库和用户
```bash
mysql -u root -p < 01_create_user_and_database.sql
```

#### 步骤2: 创建数据表
```bash
mysql -u pet_admin -p pet_adoption < 02_create_tables.sql
```

#### 步骤3: 插入初始数据（可选）
```bash
mysql -u pet_admin -p pet_adoption < 03_init_data.sql
```

## 📊 数据库结构

### 创建的对象

| 对象 | 数量 | 说明 |
|------|------|------|
| 数据库 | 1 | pet_adoption |
| 用户 | 3 | pet_admin(读写), pet_admin@localhost, pet_readonly(只读) |
| 数据表 | 10 | 见下表 |
| 索引 | 60+ | 性能优化 |
| 外键 | 15+ | 数据完整性 |

### 数据表列表

1. **users** - 用户表
2. **organizations** - 救助机构表
3. **pets** - 宠物表
4. **adoption_applications** - 领养申请表
5. **adoptions** - 领养记录表
6. **follow_up_records** - 回访记录表
7. **posts** - 社区动态表
8. **comments** - 评论表
9. **donations** - 捐赠记录表
10. **notifications** - 消息通知表

## 🔐 默认配置

### 数据库用户

| 用户名 | 主机 | 权限 | 默认密码 |
|--------|------|------|---------|
| pet_admin | % | ALL | Pet@Admin#2024!Secure |
| pet_admin | localhost | ALL | Pet@Admin#2024!Secure |
| pet_readonly | % | SELECT | Pet@ReadOnly#2024 |

⚠️ **生产环境请立即修改默认密码！**

### 应用账户（如果插入了初始数据）

| 用户名 | 角色 | 默认密码 |
|--------|------|---------|
| admin | 管理员 | Admin@123456 |
| testuser1 | 用户 | Test@123456 |
| testuser2 | 用户 | Test@123456 |
| testorg | 机构 | Test@123456 |

## ⚙️ 配置应用

部署完成后，更新 `config/config.yaml`:

```yaml
database:
  mysql:
    host: "your-server-ip"
    port: 3306
    username: "pet_admin"
    password: "YourPassword"  # 修改为实际密码
    database: "pet_adoption"
    charset: "utf8mb4"
    max_idle_conns: 10
    max_open_conns: 100
    conn_max_lifetime: 3600
```

## 🔍 验证部署

```sql
-- 登录数据库
mysql -u pet_admin -p pet_adoption

-- 查看所有表
SHOW TABLES;

-- 查看用户数量
SELECT COUNT(*) FROM users;

-- 查看数据库大小
SELECT 
    ROUND(SUM(data_length + index_length) / 1024 / 1024, 2) AS 'Size(MB)'
FROM information_schema.TABLES 
WHERE table_schema = 'pet_adoption';
```

## 🛠️ 常用操作

### 重置数据库
```bash
# ⚠️ 警告：这将删除所有数据！
mysql -u root -p -e "DROP DATABASE IF EXISTS pet_adoption;"
mysql -u root -p < 01_create_user_and_database.sql
mysql -u pet_admin -p pet_adoption < 02_create_tables.sql
```

### 备份数据库
```bash
# 完整备份
mysqldump -u pet_admin -p pet_adoption > backup_$(date +%Y%m%d).sql

# 压缩备份
mysqldump -u pet_admin -p pet_adoption | gzip > backup_$(date +%Y%m%d).sql.gz
```

### 恢复数据库
```bash
mysql -u pet_admin -p pet_adoption < backup_20241208.sql
```

## 📚 更多信息

详细文档请参考：
- [数据库部署指南.md](./数据库部署指南.md) - 完整的部署和维护文档
- [项目设计文档](../../爱心宠物领养平台设计方案.md) - 整体设计方案

## ❓ 常见问题

### 连接失败
检查：
1. MySQL 服务是否启动
2. 用户名密码是否正确
3. 防火墙是否开放 3306 端口
4. 远程访问权限是否配置

### 字符集问题
确保：
1. 数据库字符集: utf8mb4
2. 表字符集: utf8mb4
3. 连接字符集: utf8mb4
4. 配置文件 charset: utf8mb4

### 权限不足
重新授权：
```sql
GRANT ALL PRIVILEGES ON pet_adoption.* TO 'pet_admin'@'%';
FLUSH PRIVILEGES;
```

## 🔒 安全建议

1. ✅ 立即修改所有默认密码
2. ✅ 限制远程访问（如果不需要）
3. ✅ 启用 SSL 连接（生产环境）
4. ✅ 定期备份数据
5. ✅ 监控慢查询
6. ✅ 定期更新 MySQL 版本

---

**最后更新**: 2024-12-08  
**MySQL 版本**: 8.4+  
**维护者**: 项目团队
