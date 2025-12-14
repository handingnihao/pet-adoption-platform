# API文档导入指南

## 📦 生成的文件

- **apifox-api.json** - OpenAPI 3.0 格式的API文档，可直接导入Apifox

## 🚀 如何导入到Apifox

### 方法1：导入文件

1. 打开 Apifox 应用
2. 点击左上角的项目名称
3. 选择 **导入** → **从文件导入**
4. 选择格式：**OpenAPI/Swagger**
5. 选择文件：`docs/apifox-api.json`
6. 点击**确定**导入

### 方法2：直接导入

1. 打开 Apifox
2. 新建项目或选择已有项目
3. 点击**导入数据**
4. 选择 **OpenAPI/Swagger** 格式
5. 拖拽或选择 `docs/apifox-api.json` 文件
6. 配置导入选项后确认

## 📝 包含的接口

### 系统接口
- `GET /health` - 健康检查

### 用户认证
- `POST /api/v1/users/login` - 用户登录
- `POST /api/v1/users/register` - 用户注册  
- `POST /api/v1/users/send-code` - 发送验证码

### 用户信息管理
- `GET /api/v1/users/profile` - 获取个人信息
- `PUT /api/v1/users/profile` - 更新个人信息
- `PUT /api/v1/users/password` - 修改密码

### 用户管理（管理员）
- `GET /api/v1/users` - 获取用户列表
- `GET /api/v1/users/search` - 搜索用户
- `GET /api/v1/users/{id}` - 获取用户详情
- `PUT /api/v1/users/{id}/disable` - 禁用用户
- `PUT /api/v1/users/{id}/enable` - 启用用户

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

## 🔄 重新生成文档

如果API有更新，运行以下命令重新生成：

```bash
go run scripts/generate_apifox.go
```

## 📖 更多资源

- [Apifox官方文档](https://www.apifox.cn/help/)
- [OpenAPI规范](https://swagger.io/specification/)
- [项目完整API文档](./API测试文档-用户模块.md)

## ⚠️ 注意事项

1. **敏感信息**：不要在生产环境使用测试账号
2. **Token有效期**：JWT Token有效期为7天
3. **验证码**：测试环境验证码会打印在服务器日志中
4. **权限**：管理员接口需要admin角色的Token

## 🐛 常见问题

### Q: 导入后接口显示不全？
A: 检查是否选择了正确的OpenAPI版本（3.0）

### Q: 认证失败？
A: 确认Token格式正确，应为 `Bearer <token>`

### Q: 提示跨域错误？
A: 本地开发时需要配置CORS，服务器已启用跨域支持

---

生成时间：2025-12-09
版本：1.0.0
