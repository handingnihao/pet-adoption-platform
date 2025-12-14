# 开发指南

## 项目进度 ✅

### 已完成的模块

#### 1. 基础架构 ✅
- [x] Go 单体分层架构 (Controller → Service → DAO → Model)
- [x] 标准项目目录结构
- [x] 配置管理 (Viper + YAML)
- [x] 数据库 (GORM + MySQL)
- [x] 缓存 (Redis)
- [x] 日志系统 (Zap)
- [x] 统一响应格式
- [x] JWT认证
- [x] 中间件 (CORS、日志、限流、认证、权限)
- [x] Docker支持

#### 2. 用户模块 ✅ (`internal/*/user*`)
- [x] 用户注册/登录
- [x] 获取/更新个人信息
- [x] 修改密码
- [x] 用户列表（管理员）
- [x] 用户搜索（管理员）
- [x] 禁用/启用用户（管理员）

#### 3. 宠物模块 ✅ (`internal/*/pet*`)
- [x] 宠物发布
- [x] 宠物列表/详情
- [x] 宠物搜索/条件查询
- [x] 推荐宠物
- [x] 我的宠物
- [x] 更新/删除/下架宠物
- [x] 宠物审核（管理员）
- [x] 宠物统计（管理员）

#### 4. 领养模块 ✅ (`internal/*/adoption*`)
- [x] 提交领养申请
- [x] 我的申请列表
- [x] 申请详情/更新/取消
- [x] 我的领养记录
- [x] 申请审核（管理员）
- [x] 领养记录管理（管理员）
- [x] 领养统计（管理员）

#### 5. 机构模块 ✅ (`internal/*/organization*`)
- [x] 机构入驻申请
- [x] 机构列表/详情
- [x] 我的机构
- [x] 更新/删除机构
- [x] 机构审核（管理员）

#### 6. 社区模块 ✅ (`internal/*/community*`, `internal/*/post*`, `internal/*/comment*`)
- [x] 发布/更新/删除动态
- [x] 动态列表/详情/搜索
- [x] 我的动态
- [x] 评论功能（发表/删除/回复）
- [x] 点赞功能（动态/评论）

### 待开发的模块 📋

- [ ] 捐赠系统
- [ ] 回访系统
- [ ] 消息通知
- [ ] 文件上传（OSS）
- [ ] 数据统计仪表盘

## 开发环境配置

### 1. 安装 Go 环境

```bash
# 下载 Go 1.21+
# https://go.dev/dl/

# 验证安装
go version
```

### 2. 安装 MySQL

```bash
# Windows: 下载 MySQL 8.0
# https://dev.mysql.com/downloads/mysql/

# 创建数据库
mysql -u root -p
CREATE DATABASE pet_adoption CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 3. 安装 Redis

```bash
# Windows: 下载 Redis
# https://github.com/microsoftarchive/redis/releases

# 或使用 Docker
docker run -d -p 6379:6379 redis:7-alpine
```

### 4. 配置项目

```bash
# 1. 克隆项目
cd pet-adoption-platform

# 2. 安装依赖
go mod tidy

# 3. 复制配置文件
cp config/config.yaml.example config/config.yaml

# 4. 修改配置
# 编辑 config/config.yaml，修改数据库和 Redis 配置
```

### 5. 启动服务

```bash
# 方式1: 直接运行
go run cmd/server/main.go

# 方式2: 编译后运行
go build -o bin/server cmd/server/main.go
./bin/server

# 方式3: 使用 Make（推荐）
make run
```

### 6. 验证服务

```bash
# 健康检查
curl http://localhost:8080/health

# 测试接口
curl http://localhost:8080/api/v1/ping
```

## 开发工作流

```bash
# 1. 创建功能分支
git checkout -b feature/user-module

# 2. 开发功能
# 按照 Model -> DAO -> Service -> Controller 顺序开发

# 3. 编写测试
# 在 test/ 目录下编写单元测试

# 4. 运行测试
make test

# 5. 提交代码
git add .
git commit -m "feat: 完成用户模块开发"
git push origin feature/user-module

# 6. 合并到 develop
git checkout develop
git merge feature/user-module
```

## 代码规范

### 目录规范

```
internal/model/        # 数据模型
internal/dao/          # 数据访问层（数据库操作）
internal/service/      # 业务逻辑层
internal/controller/   # 控制器层（HTTP 处理）
```

### 命名规范

- **文件名**: 小写下划线 `user_service.go`
- **结构体**: 大驼峰 `UserService`
- **方法**: 大驼峰（导出）`GetUser`，小驼峰（私有）`getUser`
- **变量**: 小驼峰 `userId`
- **常量**: 大写下划线 `MAX_SIZE`

### 错误处理

```go
// 使用统一的错误响应
if err != nil {
    response.Error(c, "操作失败")
    return
}

// 带自定义状态码
if user == nil {
    response.NotFound(c, "用户不存在")
    return
}
```

### 日志记录

```go
// 使用结构化日志
logger.Info("用户登录", 
    zap.Int64("user_id", userID),
    zap.String("username", username),
)

logger.Error("数据库操作失败",
    zap.Error(err),
    zap.String("operation", "insert"),
)
```

## 常用命令

```bash
# 安装依赖
make install

# 运行服务
make run

# 编译
make build

# 测试
make test

# 代码格式化
make fmt

# Docker 相关
make docker-build    # 构建镜像
make docker-up       # 启动容器
make docker-down     # 停止容器
make docker-logs     # 查看日志

# 清理
make clean
```

## 性能优化建议

### 数据库优化
1. 为常用查询字段添加索引
2. 使用连接池复用连接
3. 避免 N+1 查询
4. 使用批量操作

### 缓存优化
1. 缓存热点数据
2. 设置合理的过期时间
3. 使用缓存预热
4. 防止缓存穿透

### 接口优化
1. 使用分页查询
2. 返回必要字段
3. 异步处理耗时操作
4. 使用 CDN 加速静态资源

## 调试技巧

### 1. 查看日志
```bash
# 实时查看日志
tail -f logs/app.log

# 查看错误日志
grep "ERROR" logs/app.log
```

### 2. 使用 Postman
- 导入 API 文档
- 测试接口
- 调试请求参数

### 3. 数据库调试
```bash
# 查看执行的 SQL
# 在配置中设置日志级别为 debug
log:
  level: "debug"
```

## 问题排查

### 数据库连接失败
```bash
# 检查配置
config/config.yaml

# 测试连接
mysql -h localhost -u root -p

# 查看日志
tail -f logs/app.log | grep "database"
```

### Redis 连接失败
```bash
# 检查 Redis 是否启动
redis-cli ping

# 查看日志
tail -f logs/app.log | grep "redis"
```

### 端口被占用
```bash
# Windows 查看端口占用
netstat -ano | findstr :8080

# 杀掉进程
taskkill /PID <进程ID> /F
```

## 参考资料

- [Go 官方文档](https://go.dev/doc/)
- [Gin 框架文档](https://gin-gonic.com/docs/)
- [GORM 文档](https://gorm.io/docs/)
- [Redis 文档](https://redis.io/docs/)
- [API文档](./API文档.md)

---

**最后更新**: 2025-12-14  
**当前进度**: 核心模块开发完成 ✅
