# 爱心宠物领养平台 (Pet Adoption Platform)

基于 Go + Gin + MySQL + Redis + React + TypeScript 的全栈宠物领养平台

## 项目简介

爱心宠物领养平台是一个连接宠物救助机构、志愿者与潜在领养者的综合性在线平台。通过信息化手段，提高流浪动物的领养率，促进人与动物的和谐共处。

### 核心价值
- 🐾 **领养代替购买**：让每一只流浪动物都能找到温暖的家
- 🤝 **连接爱心**：搭建救助机构与领养者之间的桥梁
- 📱 **便捷体验**：现代化的Web界面，流畅的用户体验
- 🔒 **安全可靠**：完善的审核机制，保障领养双方权益

## 技术栈

### 后端
- **语言**: Go 1.21+
- **框架**: Gin
- **数据库**: MySQL 8.0+
- **缓存**: Redis 7.0+
- **ORM**: GORM
- **日志**: Zap
- **配置**: Viper
- **认证**: JWT

### 前端
- **框架**: React 18 + TypeScript
- **构建工具**: Vite 6
- **UI组件**: shadcn/ui + TailwindCSS
- **状态管理**: Zustand
- **路由**: React Router v6
- **数据请求**: TanStack Query + Axios
- **图标**: Lucide React

## 项目结构

```
pet-adoption-platform/
├── cmd/                    # 后端应用入口
│   └── server/
│       └── main.go
├── config/                 # 后端配置文件
│   ├── config.yaml
│   └── config.go
├── internal/               # 后端内部代码
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
├── web/                    # 前端项目
│   ├── src/
│   │   ├── components/    # UI组件
│   │   │   ├── ui/       # 基础组件(Button, Card, Input...)
│   │   │   └── layout/   # 布局组件(Header, Footer, Layout...)
│   │   ├── pages/         # 页面组件
│   │   │   ├── admin/    # 管理后台页面
│   │   │   └── ...       # 用户端页面
│   │   ├── lib/           # 工具库(API客户端, 工具函数)
│   │   ├── store/         # 状态管理(Zustand)
│   │   ├── router/        # 路由配置
│   │   └── index.css      # 全局样式
│   ├── package.json
│   └── vite.config.ts
├── docs/                   # 项目文档
├── scripts/                # 脚本(数据库迁移等)
└── docker/                 # Docker配置
```

## 快速开始

### 1. 环境要求

- Go 1.21+
- Node.js 18+
- MySQL 8.0+
- Redis 7.0+

### 2. 后端启动

```bash
# 安装依赖
go mod tidy

# 配置文件
# 复制 config/config.yaml.example 为 config/config.yaml，修改数据库和Redis配置

# 初始化数据库
mysql -u root -p -e "CREATE DATABASE pet_adoption CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"

# 启动后端服务
go run cmd/server/main.go
```

后端服务运行在 http://localhost:8080

### 3. 前端启动

```bash
# 进入前端目录
cd web

# 安装依赖
npm install

# 启动开发服务器
npm run dev
```

前端服务运行在 http://localhost:3000

### 4. 测试账号

| 角色 | 用户名 | 密码 |
|------|--------|------|
| 管理员 | admin | Test@123456 |
| 普通用户 | testuser | Test@123456 |

### 5. 验证服务

```bash
# 后端健康检查
curl http://localhost:8080/health

# 前端访问
open http://localhost:3000
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

### 社区接口 `/api/v1/community`
| 方法 | 路径 | 权限 | 说明 |
|------|------|------|------|
| GET | `/posts` | 公开 | 动态列表 |
| GET | `/posts/search` | 公开 | 搜索动态 |
| GET | `/posts/:id` | 公开 | 动态详情 |
| GET | `/posts/:id/comments` | 公开 | 评论列表 |
| POST | `/posts` | 登录 | 发布动态 |
| GET | `/posts/my` | 登录 | 我的动态 |
| PUT | `/posts/:id` | 登录 | 更新动态 |
| DELETE | `/posts/:id` | 登录 | 删除动态 |
| POST | `/posts/:id/like` | 登录 | 点赞动态 |
| DELETE | `/posts/:id/like` | 登录 | 取消点赞 |
| POST | `/comments` | 登录 | 发表评论 |
| DELETE | `/comments/:id` | 登录 | 删除评论 |
| POST | `/comments/:id/like` | 登录 | 点赞评论 |
| DELETE | `/comments/:id/like` | 登录 | 取消点赞 |

### Swagger 文档
接口文档: http://localhost:8080/swagger/index.html

## 核心功能

### 已完成 ✅

**后端基础架构**
- [x] 项目框架搭建 (Go + Gin)
- [x] 配置管理 (Viper)
- [x] 数据库连接 (GORM + MySQL)
- [x] Redis缓存
- [x] 日志系统 (Zap)
- [x] JWT认证
- [x] 统一响应格式
- [x] 中间件（CORS、日志、限流、认证、权限）

**前端基础架构**
- [x] React + TypeScript + Vite 项目搭建
- [x] TailwindCSS + shadcn/ui 组件库
- [x] Zustand 状态管理
- [x] React Router 路由系统
- [x] Axios API客户端 + 拦截器
- [x] TanStack Query 数据请求

**用户模块**
- [x] 后端：注册/登录/个人信息/修改密码
- [x] 后端：用户列表/搜索/禁用启用（管理员）
- [x] 前端：登录/注册页面
- [x] 前端：个人中心页面
- [x] 前端：用户管理页面（管理员）

**宠物模块**
- [x] 后端：发布/列表/详情/搜索/条件查询
- [x] 后端：推荐/我的宠物/更新/删除
- [x] 后端：审核/统计（管理员）
- [x] 前端：宠物列表页（筛选/搜索）
- [x] 前端：宠物详情页
- [x] 前端：发布宠物页面
- [x] 前端：我的宠物页面
- [x] 前端：宠物管理页面（管理员，含编辑功能）

**领养模块**
- [x] 后端：提交申请/我的申请/详情/更新/取消
- [x] 后端：我的领养记录
- [x] 后端：申请审核/记录管理/统计（管理员）
- [x] 前端：领养申请页面
- [x] 前端：我的申请页面
- [x] 前端：领养管理页面（管理员）

**机构模块**
- [x] 后端：入驻申请/列表/详情/更新/删除
- [x] 后端：我的机构/审核（管理员）
- [x] 前端：机构审核页面（管理员）

**社区模块**
- [x] 后端：发布/更新/删除动态
- [x] 后端：动态列表/详情/搜索/我的动态
- [x] 后端：评论功能（发表/删除/回复）
- [x] 后端：点赞功能（动态/评论）
- [x] 前端：社区动态列表页
- [x] 前端：发布动态页面

**管理后台**
- [x] 管理员布局（侧边栏导航）
- [x] 仪表盘（统计数据）
- [x] 用户管理
- [x] 宠物管理（含编辑功能）
- [x] 领养管理
- [x] 机构审核

**其他页面**
- [x] 首页（推荐宠物、统计、介绍）
- [x] 领养指南页面

### 计划中 📋
- [ ] 捐赠系统
- [ ] 回访系统
- [ ] 消息通知
- [ ] 文件上传（OSS）
- [ ] 数据统计可视化图表
- [ ] 关于我们页面

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

## 许可证

MIT License

## 前端页面预览

| 页面 | 路径 | 说明 |
|------|------|------|
| 首页 | `/` | 平台介绍、推荐宠物、统计数据 |
| 宠物列表 | `/pets` | 浏览所有宠物，支持筛选搜索 |
| 宠物详情 | `/pets/:id` | 查看宠物详细信息 |
| 发布宠物 | `/pets/create` | 发布新宠物 |
| 领养申请 | `/adopt/:id` | 提交领养申请 |
| 领养指南 | `/guide` | 详细的领养须知 |
| 社区 | `/community` | 社区动态列表 |
| 发布动态 | `/community/create` | 发布社区动态 |
| 登录 | `/login` | 用户登录 |
| 注册 | `/register` | 用户注册 |
| 个人中心 | `/profile` | 个人信息管理 |
| 我的宠物 | `/my-pets` | 管理发布的宠物 |
| 我的申请 | `/my-applications` | 查看领养申请 |
| 管理后台 | `/admin` | 管理员仪表盘 |
| 用户管理 | `/admin/users` | 管理用户 |
| 宠物管理 | `/admin/pets` | 管理所有宠物 |
| 领养管理 | `/admin/adoptions` | 审核领养申请 |
| 机构审核 | `/admin/organizations` | 审核机构入驻 |

---

**文档版本**: v2.0  
**最后更新**: 2025-12-15
