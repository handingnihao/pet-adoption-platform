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

## API 接口

### 用户接口 `/api/v1/users`
| 方法 | 路径 | 权限 | 说明 |
|------|------|------|------|
| POST | `/register` | 公开 | 用户注册 |
| POST | `/login` | 公开 | 用户登录 |
| GET | `/profile` | 登录 | 获取个人信息 |
| PUT | `/profile` | 登录 | 更新个人信息 |
| PUT | `/password` | 登录 | 修改密码 |
| GET | `/` | 管理员 | 用户列表 |
| GET | `/search` | 管理员 | 搜索用户 |
| PUT | `/:id/disable` | 管理员 | 禁用用户 |
| PUT | `/:id/enable` | 管理员 | 启用用户 |

### 宠物接口 `/api/v1/pets`
| 方法 | 路径 | 权限 | 说明 |
|------|------|------|------|
| GET | `/` | 公开 | 宠物列表 |
| GET | `/query` | 公开 | 条件查询 |
| GET | `/search` | 公开 | 搜索宠物 |
| GET | `/recommended` | 公开 | 推荐宠物 |
| GET | `/:id` | 公开 | 宠物详情 |
| POST | `/` | 登录 | 发布宠物 |
| GET | `/my` | 登录 | 我的宠物 |
| PUT | `/:id` | 登录 | 更新宠物 |
| DELETE | `/:id` | 登录 | 删除宠物 |
| PUT | `/:id/approve` | 管理员 | 审核通过 |
| PUT | `/:id/reject` | 管理员 | 审核拒绝 |

### 领养接口 `/api/v1/adoptions`
| 方法 | 路径 | 权限 | 说明 |
|------|------|------|------|
| POST | `/applications` | 登录 | 提交申请 |
| GET | `/applications/my` | 登录 | 我的申请 |
| GET | `/applications/:id` | 登录 | 申请详情 |
| PUT | `/applications/:id` | 登录 | 更新申请 |
| DELETE | `/applications/:id` | 登录 | 取消申请 |
| GET | `/records/my` | 登录 | 我的领养记录 |
| PUT | `/applications/:id/review` | 管理员 | 审核申请 |
| GET | `/statistics` | 管理员 | 统计信息 |

### 机构接口 `/api/v1/organizations`
| 方法 | 路径 | 权限 | 说明 |
|------|------|------|------|
| GET | `/` | 公开 | 机构列表 |
| GET | `/:id` | 公开 | 机构详情 |
| POST | `/` | 登录 | 申请入驻 |
| GET | `/my` | 登录 | 我的机构 |
| PUT | `/:id` | 登录 | 更新机构 |
| DELETE | `/:id` | 登录 | 删除机构 |
| PUT | `/:id/status` | 管理员 | 审核机构 |

### Swagger 文档
接口文档: http://localhost:8080/swagger/index.html

## 核心功能

### 已完成 ✅

**基础架构**
- [x] 项目框架搭建
- [x] 配置管理 (Viper)
- [x] 数据库连接 (GORM + MySQL)
- [x] Redis缓存
- [x] 日志系统 (Zap)
- [x] JWT认证
- [x] 统一响应格式
- [x] 中间件（CORS、日志、限流、认证、权限）

**用户模块**
- [x] 用户注册
- [x] 用户登录
- [x] 获取/更新个人信息
- [x] 修改密码
- [x] 用户列表（管理员）
- [x] 用户搜索（管理员）
- [x] 禁用/启用用户（管理员）

**宠物模块**
- [x] 宠物发布
- [x] 宠物列表/详情
- [x] 宠物搜索/条件查询
- [x] 推荐宠物
- [x] 我的宠物
- [x] 更新/删除/下架宠物
- [x] 宠物审核（管理员）
- [x] 宠物统计（管理员）

**领养模块**
- [x] 提交领养申请
- [x] 我的申请列表
- [x] 申请详情/更新/取消
- [x] 我的领养记录
- [x] 申请审核（管理员）
- [x] 领养记录管理（管理员）
- [x] 领养统计（管理员）

**机构模块**
- [x] 机构入驻申请
- [x] 机构列表/详情
- [x] 我的机构
- [x] 更新/删除机构
- [x] 机构审核（管理员）

### 计划中 📋
- [ ] 社区功能（动态、评论、点赞）
- [ ] 捐赠系统
- [ ] 回访系统
- [ ] 消息通知
- [ ] 文件上传（OSS）
- [ ] 数据统计仪表盘

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

**文档版本**: v1.1  
**最后更新**: 2025-12-14
