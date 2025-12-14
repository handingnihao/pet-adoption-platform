# 开发指南

## 核心框架已完成 ✅

### 已实现的功能

#### 1. 项目架构 ✅
- [x] Go 单体分层架构
- [x] 标准项目目录结构
- [x] 模块化设计

#### 2. 核心组件 ✅

**配置管理** (`config/`)
- [x] YAML 配置文件支持
- [x] Viper 配置加载
- [x] 多环境配置
- [x] 配置结构化定义

**数据库** (`pkg/database/`)
- [x] MySQL 连接池
- [x] GORM ORM 集成
- [x] 连接参数配置
- [x] 自动重连

**缓存** (`pkg/cache/`)
- [x] Redis 客户端
- [x] 连接池管理
- [x] 常用缓存操作封装
- [x] 自动过期管理

**日志系统** (`pkg/logger/`)
- [x] Zap 日志框架
- [x] 日志分级（Debug/Info/Warn/Error）
- [x] 日志文件切割
- [x] 控制台 + 文件双输出

**统一响应** (`pkg/response/`)
- [x] 标准 JSON 响应格式
- [x] 错误码定义
- [x] 分页数据结构
- [x] 常用响应方法

**工具函数** (`pkg/utils/`)
- [x] JWT Token 生成与解析
- [x] BCrypt 密码加密
- [x] 随机字符串生成
- [x] 时间格式化

#### 3. 中间件 ✅

**CORS 中间件** (`middleware/cors.go`)
- [x] 跨域请求支持
- [x] OPTIONS 预检处理

**日志中间件** (`middleware/logger.go`)
- [x] HTTP 请求日志记录
- [x] 响应时间统计
- [x] 错误日志追踪

**认证中间件** (`middleware/auth.go`)
- [x] JWT 认证
- [x] 用户信息注入
- [x] 管理员权限验证
- [x] 机构权限验证

**限流中间件** (`middleware/rate_limit.go`)
- [x] IP 维度限流
- [x] 用户维度限流
- [x] 令牌桶算法
- [x] 自动清理机制

#### 4. 路由系统 ✅

**路由配置** (`internal/router/`)
- [x] Gin 路由集成
- [x] 路由分组
- [x] API 版本控制 (v1)
- [x] 健康检查接口
- [x] 路由骨架（待实现控制器）

#### 5. Docker 支持 ✅
- [x] Dockerfile 多阶段构建
- [x] Docker Compose 编排
- [x] MySQL + Redis + App 容器化

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

## 下一步开发任务

### Week 5-6: 用户模块开发

#### 后端任务
1. **创建用户模型** (`internal/model/user.go`)
   - 定义 User 结构体
   - GORM 标签配置
   - 表关联关系

2. **用户 DAO 层** (`internal/dao/user_dao.go`)
   - CreateUser: 创建用户
   - GetUserByID: 根据 ID 查询
   - GetUserByPhone: 根据手机号查询
   - GetUserByEmail: 根据邮箱查询
   - UpdateUser: 更新用户信息

3. **用户 Service 层** (`internal/service/user_service.go`)
   - Register: 用户注册逻辑
   - Login: 用户登录逻辑
   - GetProfile: 获取个人信息
   - UpdateProfile: 更新个人信息
   - UploadAvatar: 上传头像

4. **用户 Controller** (`internal/controller/user_controller.go`)
   - RegisterHandler: 注册接口
   - LoginHandler: 登录接口
   - GetProfileHandler: 获取信息接口
   - UpdateProfileHandler: 更新信息接口
   - UploadAvatarHandler: 上传头像接口

5. **注册路由**
   - 在 `router.go` 中取消注释用户路由
   - 绑定控制器方法

#### 前端任务（另一个分支）
- 注册页面
- 登录页面
- 个人中心页面

### 开发工作流

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
- [项目设计文档](../爱心宠物领养平台设计方案.md)

---

**最后更新**: 2024-12-08  
**当前进度**: Week 1-4 核心框架搭建 ✅
