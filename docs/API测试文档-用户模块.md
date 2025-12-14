# 用户模块API测试文档

## 📝 基础信息

- **基础URL**: `http://localhost:8080/api/v1`
- **Content-Type**: `application/json`
- **认证方式**: Bearer Token

---

## 1. 发送验证码

### 请求

```
POST /users/send-code
Content-Type: application/json
```

```json
{
  "phone": "13800138000"
}
```

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "message": "验证码已发送"
  },
  "timestamp": 1733629200
}
```

### cURL 测试

```bash
curl -X POST http://localhost:8080/api/v1/users/send-code \
  -H "Content-Type: application/json" \
  -d '{"phone":"13800138000"}'
```

### 注意事项
- 同一手机号5分钟内只能发送一次
- 验证码5分钟内有效
- 开发环境验证码会打印在控制台日志中

---

## 2. 用户注册

### 请求

```
POST /users/register
Content-Type: application/json
```

```json
{
  "username": "testuser",
  "password": "123456",
  "phone": "13800138000",
  "email": "test@example.com",
  "code": "123456"
}
```

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "user_id": 1,
    "message": "注册成功"
  },
  "timestamp": 1733629200
}
```

### cURL 测试

```bash
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "123456",
    "phone": "13800138000",
    "email": "test@example.com",
    "code": "123456"
  }'
```

### 参数说明

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| username | string | 是 | 用户名，3-20字符 |
| password | string | 是 | 密码，6-20字符 |
| phone | string | 是 | 手机号，11位 |
| email | string | 是 | 邮箱 |
| code | string | 是 | 验证码，6位数字 |

### 错误响应

```json
{
  "code": 400,
  "message": "用户名已存在",
  "data": null,
  "timestamp": 1733629200
}
```

---

## 3. 用户登录

### 请求

```
POST /users/login
Content-Type: application/json
```

```json
{
  "username": "testuser",
  "password": "123456"
}
```

**说明**: username 支持以下三种方式：
- 用户名：`testuser`
- 手机号：`13800138000`
- 邮箱：`test@example.com`

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_at": "2024-12-15T12:00:00Z",
    "user": {
      "id": 1,
      "username": "testuser",
      "real_name": "",
      "phone": "13800138000",
      "email": "test@example.com",
      "avatar": "",
      "gender": 0,
      "role": "user",
      "status": 1
    }
  },
  "timestamp": 1733629200
}
```

### cURL 测试

```bash
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "123456"
  }'
```

### 错误响应

```json
{
  "code": 400,
  "message": "密码错误",
  "data": null,
  "timestamp": 1733629200
}
```

---

## 4. 获取个人信息

### 请求

```
GET /users/profile
Authorization: Bearer {token}
```

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "username": "testuser",
    "real_name": "张三",
    "phone": "13800138000",
    "email": "test@example.com",
    "avatar": "https://example.com/avatar.jpg",
    "gender": 1,
    "role": "user",
    "status": 1
  },
  "timestamp": 1733629200
}
```

### cURL 测试

```bash
curl -X GET http://localhost:8080/api/v1/users/profile \
  -H "Authorization: Bearer YOUR_TOKEN"
```

---

## 5. 更新个人信息

### 请求

```
PUT /users/profile
Authorization: Bearer {token}
Content-Type: application/json
```

```json
{
  "real_name": "张三",
  "avatar": "https://example.com/avatar.jpg",
  "gender": 1,
  "birthday": "1990-01-01",
  "address": "北京市朝阳区"
}
```

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "message": "更新成功"
  },
  "timestamp": 1733629200
}
```

### cURL 测试

```bash
curl -X PUT http://localhost:8080/api/v1/users/profile \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "real_name": "张三",
    "gender": 1,
    "address": "北京市朝阳区"
  }'
```

---

## 6. 修改密码

### 请求

```
PUT /users/password
Authorization: Bearer {token}
Content-Type: application/json
```

```json
{
  "old_password": "123456",
  "new_password": "654321"
}
```

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "message": "密码修改成功，请重新登录"
  },
  "timestamp": 1733629200
}
```

### cURL 测试

```bash
curl -X PUT http://localhost:8080/api/v1/users/password \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "old_password": "123456",
    "new_password": "654321"
  }'
```

---

## 7. 获取用户列表（管理员）

### 请求

```
GET /users?page=1&page_size=20&role=user&status=1
Authorization: Bearer {admin_token}
```

### 查询参数

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| page | int | 否 | 页码，默认1 |
| page_size | int | 否 | 每页数量，默认20，最大100 |
| role | string | 否 | 角色筛选：user/organization/volunteer/admin |
| status | int | 否 | 状态筛选：0-禁用，1-正常 |

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "username": "testuser",
        "real_name": "张三",
        "phone": "13800138000",
        "email": "test@example.com",
        "avatar": "",
        "gender": 1,
        "role": "user",
        "status": 1
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 20,
      "total": 100,
      "total_pages": 5
    }
  },
  "timestamp": 1733629200
}
```

### cURL 测试

```bash
curl -X GET "http://localhost:8080/api/v1/users?page=1&page_size=20&role=user" \
  -H "Authorization: Bearer ADMIN_TOKEN"
```

---

## 8. 搜索用户（管理员）

### 请求

```
GET /users/search?keyword=张三&page=1&page_size=20
Authorization: Bearer {admin_token}
```

### 查询参数

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| keyword | string | 是 | 搜索关键词（用户名/手机/邮箱/真实姓名） |
| page | int | 否 | 页码，默认1 |
| page_size | int | 否 | 每页数量，默认20 |

### cURL 测试

```bash
curl -X GET "http://localhost:8080/api/v1/users/search?keyword=test&page=1" \
  -H "Authorization: Bearer ADMIN_TOKEN"
```

---

## 9. 禁用用户（管理员）

### 请求

```
PUT /users/:id/disable
Authorization: Bearer {admin_token}
```

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "message": "用户已禁用"
  },
  "timestamp": 1733629200
}
```

### cURL 测试

```bash
curl -X PUT http://localhost:8080/api/v1/users/1/disable \
  -H "Authorization: Bearer ADMIN_TOKEN"
```

---

## 10. 启用用户（管理员）

### 请求

```
PUT /users/:id/enable
Authorization: Bearer {admin_token}
```

### cURL 测试

```bash
curl -X PUT http://localhost:8080/api/v1/users/1/enable \
  -H "Authorization: Bearer ADMIN_TOKEN"
```

---

## 🧪 完整测试流程

### 1. 准备工作

```bash
# 确保服务已启动
go run cmd/server/main.go

# 测试健康检查
curl http://localhost:8080/health
```

### 2. 注册流程测试

```bash
# Step 1: 发送验证码
curl -X POST http://localhost:8080/api/v1/users/send-code \
  -H "Content-Type: application/json" \
  -d '{"phone":"13900000001"}'

# 查看控制台日志，获取验证码（开发环境）

# Step 2: 注册用户
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "newuser",
    "password": "123456",
    "phone": "13900000001",
    "email": "newuser@test.com",
    "code": "123456"
  }'
```

### 3. 登录流程测试

```bash
# 登录获取token
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "newuser",
    "password": "123456"
  }' | jq .

# 保存返回的 token
TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

### 4. 用户信息测试

```bash
# 获取个人信息
curl -X GET http://localhost:8080/api/v1/users/profile \
  -H "Authorization: Bearer $TOKEN" | jq .

# 更新个人信息
curl -X PUT http://localhost:8080/api/v1/users/profile \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "real_name": "测试用户",
    "gender": 1,
    "address": "测试地址"
  }' | jq .
```

### 5. 管理员功能测试

```bash
# 使用管理员账户登录
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "Admin@123456"
  }' | jq .

# 保存管理员 token
ADMIN_TOKEN="..."

# 获取用户列表
curl -X GET "http://localhost:8080/api/v1/users?page=1&page_size=10" \
  -H "Authorization: Bearer $ADMIN_TOKEN" | jq .

# 搜索用户
curl -X GET "http://localhost:8080/api/v1/users/search?keyword=test" \
  -H "Authorization: Bearer $ADMIN_TOKEN" | jq .
```

---

## ❗ 错误码说明

| 错误码 | 说明 |
|--------|------|
| 200 | 成功 |
| 400 | 请求参数错误 |
| 401 | 未认证（未登录或token无效） |
| 403 | 无权限 |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |

---

## 📌 注意事项

1. **Token获取**: 登录成功后会返回token，后续需要认证的接口都要在Header中携带
2. **Token格式**: `Authorization: Bearer {token}`
3. **Token有效期**: 默认168小时（7天）
4. **验证码**: 开发环境验证码会打印在控制台，生产环境需要配置短信服务
5. **管理员账户**: 默认用户名`admin`，密码`Admin@123456`（请及时修改）

---

## 🔧 Postman 导入

可以将以上接口导入Postman进行测试，创建以下环境变量：

- `base_url`: `http://localhost:8080/api/v1`
- `token`: 登录后获取的token
- `admin_token`: 管理员token

---

**最后更新**: 2024-12-08  
**版本**: v1.0
