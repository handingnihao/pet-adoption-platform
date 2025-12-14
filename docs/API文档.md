# 宠物领养平台 API 文档

## 📝 基础信息

- **基础URL**: `http://localhost:8080/api/v1`
- **Content-Type**: `application/json`
- **认证方式**: Bearer Token (JWT)

---

## 🔑 测试账号

| 用户名 | 密码 | 角色 |
|--------|------|------|
| admin | Test@123456 | 管理员 |
| testuser | Test@123456 | 普通用户 |

---

## 一、用户模块 `/users`

### 1.1 用户注册

```http
POST /users/register
```

**请求体:**
```json
{
  "username": "testuser",
  "password": "Test@123456",
  "phone": "13800138000",
  "email": "test@example.com"
}
```

**响应:**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "user_id": 1,
    "message": "注册成功"
  }
}
```

### 1.2 用户登录

```http
POST /users/login
```

**请求体:**
```json
{
  "username": "admin",
  "password": "Test@123456"
}
```

**响应:**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIs...",
    "expires_at": "2025-12-21T22:00:00Z",
    "user": {
      "id": 1,
      "username": "admin",
      "role": "admin",
      "status": 1
    }
  }
}
```

### 1.3 获取个人信息

```http
GET /users/profile
Authorization: Bearer {token}
```

### 1.4 更新个人信息

```http
PUT /users/profile
Authorization: Bearer {token}
```

**请求体:**
```json
{
  "real_name": "张三",
  "avatar": "https://example.com/avatar.jpg",
  "gender": 1,
  "address": "北京市朝阳区"
}
```

### 1.5 修改密码

```http
PUT /users/password
Authorization: Bearer {token}
```

**请求体:**
```json
{
  "old_password": "Test@123456",
  "new_password": "NewPassword@123"
}
```

### 1.6 获取用户列表（管理员）

```http
GET /users?page=1&page_size=20&role=user&status=1
Authorization: Bearer {admin_token}
```

### 1.7 搜索用户（管理员）

```http
GET /users/search?keyword=test&page=1
Authorization: Bearer {admin_token}
```

### 1.8 禁用/启用用户（管理员）

```http
PUT /users/{id}/disable
PUT /users/{id}/enable
Authorization: Bearer {admin_token}
```

---

## 二、宠物模块 `/pets`

### 2.1 获取宠物列表

```http
GET /pets?page=1&page_size=20
```

### 2.2 条件查询宠物

```http
GET /pets/query?type=dog&city=深圳市&status=1
```

**查询参数:**
| 参数 | 类型 | 说明 |
|------|------|------|
| type | string | 宠物类型(dog/cat/rabbit/bird/other) |
| breed | string | 品种 |
| gender | string | 性别(male/female) |
| city | string | 城市 |
| status | int | 状态(0待审核/1可领养/2已领养/3已下架) |

### 2.3 搜索宠物

```http
GET /pets/search?keyword=金毛&page=1
```

### 2.4 获取推荐宠物

```http
GET /pets/recommended?limit=10
```

### 2.5 获取宠物详情

```http
GET /pets/{id}
```

### 2.6 发布宠物

```http
POST /pets
Authorization: Bearer {token}
```

**请求体:**
```json
{
  "name": "小白",
  "type": "dog",
  "breed": "金毛寻回犬",
  "gender": "male",
  "age": 18,
  "size": "large",
  "color": "金黄色",
  "weight": 28.5,
  "is_vaccinated": true,
  "is_sterilized": false,
  "health_status": "健康",
  "description": "温顺可爱的金毛",
  "character": "活泼、温顺",
  "cover_photo": "https://example.com/photo.jpg",
  "province": "广东省",
  "city": "深圳市",
  "district": "南山区",
  "address": "科技园"
}
```

### 2.7 获取我的宠物

```http
GET /pets/my?page=1&page_size=20
Authorization: Bearer {token}
```

### 2.8 更新宠物信息

```http
PUT /pets/{id}
Authorization: Bearer {token}
```

### 2.9 删除宠物

```http
DELETE /pets/{id}
Authorization: Bearer {token}
```

### 2.10 下架宠物

```http
PUT /pets/{id}/offline
Authorization: Bearer {token}
```

### 2.11 审核宠物（管理员）

```http
PUT /pets/{id}/approve
PUT /pets/{id}/reject
Authorization: Bearer {admin_token}
```

**拒绝时请求体:**
```json
{
  "reason": "信息不完整，请补充照片"
}
```

### 2.12 获取待审核宠物（管理员）

```http
GET /pets/pending?page=1&page_size=20
Authorization: Bearer {admin_token}
```

### 2.13 获取宠物统计（管理员）

```http
GET /pets/statistics
Authorization: Bearer {admin_token}
```

---

## 三、领养模块 `/adoptions`

### 3.1 提交领养申请

```http
POST /adoptions/applications
Authorization: Bearer {token}
```

**请求体:**
```json
{
  "pet_id": 1,
  "applicant_name": "张三",
  "applicant_phone": "13800138000",
  "applicant_address": "北京市朝阳区",
  "housing_type": "apartment",
  "has_pet_experience": true,
  "pet_experience": "养过两只猫",
  "family_agree": true,
  "adoption_reason": "喜欢宠物，有时间照顾"
}
```

### 3.2 获取我的申请列表

```http
GET /adoptions/applications/my?page=1&page_size=20
Authorization: Bearer {token}
```

### 3.3 获取申请详情

```http
GET /adoptions/applications/{id}
Authorization: Bearer {token}
```

### 3.4 更新申请信息

```http
PUT /adoptions/applications/{id}
Authorization: Bearer {token}
```

### 3.5 取消申请

```http
DELETE /adoptions/applications/{id}
Authorization: Bearer {token}
```

### 3.6 获取我的领养记录

```http
GET /adoptions/records/my?page=1&page_size=20
Authorization: Bearer {token}
```

### 3.7 获取领养记录详情

```http
GET /adoptions/records/{id}
Authorization: Bearer {token}
```

### 3.8 审核申请（管理员）

```http
PUT /adoptions/applications/{id}/review
Authorization: Bearer {admin_token}
```

**请求体:**
```json
{
  "status": "approved",
  "comment": "审核通过，请联系机构领取宠物"
}
```

**状态值:** pending, reviewing, approved, rejected, cancelled

### 3.9 获取所有申请列表（管理员）

```http
GET /adoptions/applications?page=1&page_size=20&status=pending
Authorization: Bearer {admin_token}
```

### 3.10 获取待审核申请（管理员）

```http
GET /adoptions/applications/pending?page=1
Authorization: Bearer {admin_token}
```

### 3.11 创建领养记录（管理员）

```http
POST /adoptions/records
Authorization: Bearer {admin_token}
```

### 3.12 获取领养统计（管理员）

```http
GET /adoptions/statistics
Authorization: Bearer {admin_token}
```

---

## 四、机构模块 `/organizations`

### 4.1 获取机构列表

```http
GET /organizations?page=1&page_size=20&status=1
```

### 4.2 获取机构详情

```http
GET /organizations/{id}
```

### 4.3 申请入驻/创建机构

```http
POST /organizations
Authorization: Bearer {token}
```

**请求体:**
```json
{
  "name": "爱心宠物救助中心",
  "logo": "https://example.com/logo.jpg",
  "description": "专注于流浪动物救助和领养的非营利组织",
  "address": "北京市朝阳区建国路88号",
  "phone": "13800138001",
  "email": "love@example.com"
}
```

### 4.4 获取我创建的机构

```http
GET /organizations/my?page=1
Authorization: Bearer {token}
```

### 4.5 更新机构信息

```http
PUT /organizations/{id}
Authorization: Bearer {token}
```

### 4.6 删除机构

```http
DELETE /organizations/{id}
Authorization: Bearer {token}
```

### 4.7 审核机构（管理员）

```http
PUT /organizations/{id}/status
Authorization: Bearer {admin_token}
```

**请求体:**
```json
{
  "status": 1,
  "reject_reason": ""
}
```

**状态值:** 0-待审核, 1-已通过, 2-已拒绝

---

## 五、系统接口

### 5.1 健康检查

```http
GET /health
```

**响应:**
```json
{
  "status": "ok",
  "message": "服务运行正常"
}
```

### 5.2 测试接口

```http
GET /api/v1/ping
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

1. **Token格式**: `Authorization: Bearer {token}`
2. **Token有效期**: 默认168小时（7天）
3. **管理员功能**: 需要admin角色的Token
4. **分页参数**: page从1开始，page_size默认20，最大100

---

**最后更新**: 2025-12-14  
**版本**: v2.0
