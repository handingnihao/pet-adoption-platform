# 测试指南

本目录包含爱心宠物领养平台的所有测试代码和文档。

## 目录结构

```
test/
├── api/                    # 后端API接口测试
│   └── api_test.go        # Go API测试代码
├── postman/               # Postman测试集合
│   └── pet-adoption-api.postman_collection.json
├── API测试文档.md         # API测试用例文档
├── 前端测试方案.md         # 前端测试方案和示例
└── README.md             # 本文件
```

## 快速开始

### 后端测试

#### 1. 运行Go API测试

```bash
# 在项目根目录运行
make test-api

# 或直接使用go test
cd test/api
go test -v

# 生成覆盖率报告
make test-coverage
```

#### 2. 使用Postman测试

1. 打开Postman
2. 点击 Import
3. 选择 `test/postman/pet-adoption-api.postman_collection.json`
4. 导入后即可使用所有测试用例

### 前端测试

#### 1. 安装测试依赖

```bash
cd web
npm install
```

#### 2. 运行单元测试和组件测试

```bash
# 监听模式（开发时使用）
npm test

# 运行一次
npm run test:run

# UI模式
npm run test:ui

# 生成覆盖率报告
npm run test:coverage
```

#### 3. 运行E2E测试

```bash
# 首次运行需要安装浏览器
npx playwright install

# 运行E2E测试
npm run test:e2e

# UI模式运行
npm run test:e2e:ui
```

### 配置测试环境

在运行测试前，确保：

1. 后端服务已启动（默认端口8080）
2. 数据库已初始化
3. Redis已启动
4. 前端开发服务器已启动（E2E测试需要，端口80）

可以使用以下命令启动服务：

```bash
# 启动后端服务
make run

# 启动前端服务
cd web
npm run dev

# 或使用docker-compose
cd docker
docker-compose up -d
```

## 测试覆盖范围

### 后端API接口测试

- ✅ 用户管理（注册、登录、个人信息）
- ✅ 宠物管理（CRUD、搜索、推荐）
- ✅ 领养管理（申请、审核、记录）
- ✅ 机构管理（CRUD、查询）
- ✅ 社区互动（帖子、评论、点赞）
- ✅ 捐赠管理（创建、查询、统计）

详细测试用例请查看 [API测试文档.md](./API测试文档.md)

### 前端测试

- ✅ 单元测试：工具函数、Hooks、Store
- ✅ 组件测试：UI组件、表单组件、页面组件
- ✅ E2E测试：用户流程、关键业务场景

详细测试方案请查看 [前端测试方案.md](./前端测试方案.md)

## 测试文件位置

### 后端测试
- API测试：`test/api/api_test.go`
- 单元测试：`internal/*/dao/*_test.go`（如需要）

### 前端测试
- 单元测试：`web/src/**/__tests__/*.test.ts(x)`
- E2E测试：`web/e2e/*.spec.ts`
- 测试配置：`web/vitest.config.ts`, `web/playwright.config.ts`
- 测试设置：`web/src/test/setup.ts`

## 测试最佳实践

### 后端测试
1. **测试隔离**：每个测试用例应该独立运行
2. **数据清理**：测试后清理测试数据
3. **Mock数据**：使用真实但不敏感的测试数据
4. **错误处理**：测试正常流程和异常情况

### 前端测试
1. **测试行为，不是实现**：关注用户交互和结果
2. **保持测试简单**：一个测试只验证一个功能点
3. **使用真实场景**：模拟真实用户操作
4. **避免测试耦合**：测试之间相互独立

## 测试数据

### 测试账号

| 用户名 | 密码 | 角色 |
|--------|------|------|
| admin | Test@123456 | 管理员 |
| testuser | Test@123456 | 普通用户 |

### 测试环境

- 本地后端: http://localhost:8080
- 本地前端: http://localhost:80

## 常见问题

### Q: 后端测试失败怎么办？

A: 检查以下几点：
1. 服务是否正常运行
2. 数据库连接是否正常
3. Redis是否启动
4. 测试数据是否准备好

### Q: 前端测试失败怎么办？

A: 检查以下几点：
1. 依赖是否安装完整（运行 `npm install`）
2. 测试配置文件是否正确
3. 查看错误日志定位问题

### Q: 如何查看测试覆盖率？

A: 
- 后端：运行 `make test-coverage`
- 前端：运行 `npm run test:coverage`

### Q: 如何添加新的测试用例？

A: 
- 后端：在 `test/api/api_test.go` 中添加新的测试函数
- 前端：在对应模块的 `__tests__` 目录下创建测试文件

## 持续集成

项目使用GitHub Actions进行CI/CD，每次提交都会自动运行测试。

查看 `.github/workflows/` 目录了解CI配置。

## 贡献指南

### 添加后端测试
1. 在 `test/api/api_test.go` 中添加测试函数
2. 更新 `API测试文档.md` 添加测试用例说明
3. 在Postman集合中添加对应的请求
4. 确保所有测试通过后再提交

### 添加前端测试
1. 在对应模块的 `__tests__` 目录下创建测试文件
2. 遵循测试命名规范：`*.test.ts` 或 `*.test.tsx`
3. 使用 `describe` 和 `it` 组织测试用例
4. 运行测试确保通过后再提交

## 相关文档

- [API测试文档](./API测试文档.md) - 后端API测试用例详细说明
- [前端测试方案](./前端测试方案.md) - 前端测试配置和示例代码
- [开发指南](../docs/DEVELOPMENT.md) - 项目开发指南
- [部署指南](../docs/部署指南.md) - 项目部署说明
