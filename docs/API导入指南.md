# API文档导入指南

## 📦 文档文件

- **API文档.md** - 完整的API接口文档
- **apifox-api.json** - OpenAPI 3.0 格式，可导入Apifox/Postman

## 🚀 如何导入到Apifox

1. 打开 Apifox 应用
2. 新建项目或选择已有项目
3. 点击 **导入** → **从文件导入**
4. 选择格式：**OpenAPI/Swagger**
5. 选择文件：`docs/apifox-api.json`
6. 点击**确定**导入

## 📝 已实现的接口

### 系统接口
- `GET /health` - 健康检查
- `GET /api/v1/ping` - 测试接口

### 用户模块 (9个接口)
- `POST /api/v1/users/register` - 用户注册
- `POST /api/v1/users/login` - 用户登录
- `GET /api/v1/users/profile` - 获取个人信息
- `PUT /api/v1/users/profile` - 更新个人信息
- `PUT /api/v1/users/password` - 修改密码
- `GET /api/v1/users` - 用户列表（管理员）
- `GET /api/v1/users/search` - 搜索用户（管理员）
- `PUT /api/v1/users/{id}/disable` - 禁用用户（管理员）
- `PUT /api/v1/users/{id}/enable` - 启用用户（管理员）

### 宠物模块 (13个接口)
- `GET /api/v1/pets` - 宠物列表
- `GET /api/v1/pets/query` - 条件查询
- `GET /api/v1/pets/search` - 搜索宠物
- `GET /api/v1/pets/recommended` - 推荐宠物
- `GET /api/v1/pets/{id}` - 宠物详情
- `POST /api/v1/pets` - 发布宠物
- `GET /api/v1/pets/my` - 我的宠物
- `PUT /api/v1/pets/{id}` - 更新宠物
- `DELETE /api/v1/pets/{id}` - 删除宠物
- `PUT /api/v1/pets/{id}/offline` - 下架宠物
- `GET /api/v1/pets/pending` - 待审核列表（管理员）
- `PUT /api/v1/pets/{id}/approve` - 审核通过（管理员）
- `PUT /api/v1/pets/{id}/reject` - 审核拒绝（管理员）

### 领养模块 (12个接口)
- `POST /api/v1/adoptions/applications` - 提交申请
- `GET /api/v1/adoptions/applications/my` - 我的申请
- `GET /api/v1/adoptions/applications/{id}` - 申请详情
- `PUT /api/v1/adoptions/applications/{id}` - 更新申请
- `DELETE /api/v1/adoptions/applications/{id}` - 取消申请
- `GET /api/v1/adoptions/records/my` - 我的领养记录
- `GET /api/v1/adoptions/records/{id}` - 记录详情
- `GET /api/v1/adoptions/applications` - 申请列表（管理员）
- `GET /api/v1/adoptions/applications/pending` - 待审核申请（管理员）
- `PUT /api/v1/adoptions/applications/{id}/review` - 审核申请（管理员）
- `POST /api/v1/adoptions/records` - 创建记录（管理员）
- `GET /api/v1/adoptions/statistics` - 统计信息（管理员）

### 机构模块 (7个接口)
- `GET /api/v1/organizations` - 机构列表
- `GET /api/v1/organizations/{id}` - 机构详情
- `POST /api/v1/organizations` - 申请入驻
- `GET /api/v1/organizations/my` - 我的机构
- `PUT /api/v1/organizations/{id}` - 更新机构
- `DELETE /api/v1/organizations/{id}` - 删除机构
- `PUT /api/v1/organizations/{id}/status` - 审核机构（管理员）

## 🔑 测试账号

| 用户名 | 密码 | 角色 |
|--------|------|------|
| admin | Test@123456 | 管理员 |
| testuser1 | Test@123456 | 普通用户 |
| testuser2 | Test@123456 | 普通用户 |
| testorg | Test@123456 | 机构用户 |

## 🌐 服务器地址

### 本地开发环境
```
http://localhost:8080
```

### 测试服务器
```
http://115.190.125.177:8080
```

## 🔐 认证方式

大部分接口需要JWT Token认证：

1. 先调用登录接口获取token
2. 在需要认证的接口header中添加：
   ```
   Authorization: Bearer <你的token>
   ```

## 💡 使用提示

### 1. 设置环境变量（推荐）

在Apifox中设置环境变量：

```javascript
// 本地环境
{
  "baseUrl": "http://localhost:8080",
  "token": ""
}

// 测试环境
{
  "baseUrl": "http://115.190.125.177:8080",
  "token": ""
}
```

### 2. 自动设置Token

在登录接口的**后置操作**中添加脚本：

```javascript
// 自动保存token到环境变量
const res = pm.response.json();
if (res.code === 200 && res.data.token) {
  pm.environment.set("token", res.data.token);
  console.log("Token已保存");
}
```

### 3. 全局认证头

在项目设置中配置全局Header：

```
Authorization: Bearer {{token}}
```

## 📖 更多资源

- [Apifox官方文档](https://www.apifox.cn/help/)
- [OpenAPI规范](https://swagger.io/specification/)
- [完整API文档](./API文档.md)

## ⚠️ 注意事项

1. **敏感信息**：不要在生产环境使用测试账号
2. **Token有效期**：JWT Token有效期为7天
3. **权限**：管理员接口需要admin角色的Token

## 🐛 常见问题

### Q: 导入后接口显示不全？
A: 检查是否选择了正确的OpenAPI版本（3.0）

### Q: 认证失败？
A: 确认Token格式正确，应为 `Bearer <token>`

### Q: 提示跨域错误？
A: 本地开发时需要配置CORS，服务器已启用跨域支持

---

**最后更新**: 2025-12-14  
**版本**: 2.0.0
