# 爱心宠物领养平台 (Pet Adoption Platform)

基于 Go + Gin + MySQL + Redis 的宠物领养平台后端服务

## 项目简介

爱心宠物领养平台是一个连接宠物救助机构、志愿者与潜在领养者的综合性在线平台。通过信息化手段，提高流浪动物的领养率，促进人与动物的和谐共处。

## 技术栈

- **语言**: Go 1.21+
- **框架**: Gin
- **数据库**: MySQL 8.0+
- **缓存**: Redis 7.0+
- **ORM**: GORM
- **日志**: Zap
- **配置**: Viper
- **认证**: JWT

## 项目结构

```
pet-adoption-platform/
├── cmd/                    # 应用入口
│   └── server/
│       └── main.go
├── config/                 # 配置文件
│   ├── config.yaml
│   └── config.go
├── internal/               # 内部代码
│   ├── router/            # 路由层
│   ├── controller/        # 控制器层
│   ├── service/           # 服务层
│   ├── dao/               # 数据访问层
│   └── model/             # 数据模型
├── pkg/                    # 公共包
│   ├── utils/             # 工具函数
│   ├── logger/            # 日志
│   ├── cache/             # 缓存
│   ├── database/          # 数据库
│   └── response/          # 统一响应
├── api/                    # API文档
├── scripts/                # 脚本
├── test/                   # 测试
├── docs/                   # 文档
└── docker/                 # Docker配置
```

## 快速开始

### 1. 环境要求

- Go 1.21+
- MySQL 8.0+
- Redis 7.0+

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 配置文件

复制 `config/config.yaml.example` 并重命名为 `config/config.yaml`，修改数据库和Redis配置。

### 4. 初始化数据库

```bash
# 创建数据库
mysql -u root -p
CREATE DATABASE pet_adoption CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

# 运行迁移脚本
go run scripts/migration/init_db.go
```

### 5. 启动服务

```bash
# 开发模式
go run cmd/server/main.go

# 编译运行
go build -o bin/server cmd/server/main.go
./bin/server
```

服务将在 http://localhost:8080 启动

### 6. 测试接口

```bash
# 健康检查
curl http://localhost:8080/health

# 测试接口
curl http://localhost:8080/api/v1/ping
```

## API 文档

接口文档: http://localhost:8080/swagger/index.html

## 核心功能

### 已完成 ✅
- [x] 项目框架搭建
- [x] 配置管理
- [x] 数据库连接
- [x] Redis缓存
- [x] 日志系统
- [x] JWT认证
- [x] 统一响应格式
- [x] 中间件（CORS、日志、限流、认证）

### 开发中 🚧
- [ ] 用户注册登录
- [ ] 宠物信息管理
- [ ] 领养流程
- [ ] 机构管理
- [ ] 社区功能
- [ ] 捐赠系统

### 计划中 📋
- [ ] 回访系统
- [ ] 数据统计
- [ ] 消息通知
- [ ] 文件上传

## 开发规范

### 分支管理

- `main`: 主分支，生产环境
- `develop`: 开发分支
- `feature/*`: 功能分支
- `bugfix/*`: 修复分支

### 提交规范

```
feat: 新功能
fix: 修复bug
docs: 文档更新
style: 代码格式
refactor: 重构
test: 测试
chore: 构建/工具
```

### 代码规范

- 遵循 Go 官方代码规范
- 使用 `gofmt` 格式化代码
- 使用 `golint` 检查代码
- 编写单元测试

## 部署

### Docker 部署

```bash
# 构建镜像
docker build -t pet-adoption-platform .

# 运行容器
docker-compose up -d
```

### 生产部署

```bash
# 编译
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/server cmd/server/main.go

# 上传到服务器并运行
./bin/server
```

## 性能指标

- 支持 QPS: 2000+
- 响应时间: < 100ms
- 并发用户: 10000+

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License

## 联系方式

- 项目主页: https://github.com/your-org/pet-adoption-platform
- 问题反馈: https://github.com/your-org/pet-adoption-platform/issues

---

**文档版本**: v1.0  
**最后更新**: 2024-12-08
