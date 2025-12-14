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

## 五、社区模块 `/community`

### 5.1 获取动态列表

```http
GET /community/posts?page=1&page_size=20&type=daily
```

**参数:**
| 参数 | 类型 | 说明 |
|------|------|------|
| page | int | 页码（默认1） |
| page_size | int | 每页数量（默认20） |
| type | string | 类型筛选: story/knowledge/daily/other |

**响应:**
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "id": 1,
        "user_id": 1,
        "username": "testuser",
        "user_avatar": "",
        "title": "领养日记",
        "content": "今天带小橘回家了...",
        "images": ["https://example.com/1.jpg"],
        "video_url": "",
        "type": "story",
        "view_count": 100,
        "like_count": 10,
        "comment_count": 5,
        "is_liked": false,
        "created_at": "2025-12-14T12:00:00Z"
      }
    ],
    "total": 50,
    "page": 1,
    "page_size": 20
  }
}
```

### 5.2 获取动态详情

```http
GET /community/posts/:id
```

### 5.3 搜索动态

```http
GET /community/posts/search?keyword=领养&page=1&page_size=20
```

### 5.4 发布动态

```http
POST /community/posts
Authorization: Bearer {token}
```

**请求体:**
```json
{
  "title": "我的领养故事",
  "content": "今天带小橘回家了，很开心！",
  "images": ["https://example.com/1.jpg", "https://example.com/2.jpg"],
  "video_url": "",
  "type": "story",
  "pet_id": 1
}
```

**type 类型:**
- `story` - 领养故事
- `knowledge` - 养宠知识
- `daily` - 日常分享
- `other` - 其他

### 5.5 更新动态

```http
PUT /community/posts/:id
Authorization: Bearer {token}
```

### 5.6 删除动态

```http
DELETE /community/posts/:id
Authorization: Bearer {token}
```

### 5.7 点赞动态

```http
POST /community/posts/:id/like
Authorization: Bearer {token}
```

### 5.8 取消点赞动态

```http
DELETE /community/posts/:id/like
Authorization: Bearer {token}
```

### 5.9 获取我的动态

```http
GET /community/posts/my?page=1&page_size=20
Authorization: Bearer {token}
```

### 5.10 获取动态评论列表

```http
GET /community/posts/:id/comments?page=1&page_size=20
```

**响应:**
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "id": 1,
        "post_id": 1,
        "user_id": 2,
        "username": "user2",
        "user_avatar": "",
        "parent_id": 0,
        "content": "好可爱的猫咪！",
        "like_count": 3,
        "is_liked": false,
        "created_at": "2025-12-14T14:00:00Z",
        "replies": [
          {
            "id": 2,
            "user_id": 1,
            "username": "testuser",
            "parent_id": 1,
            "reply_to_user_id": 2,
            "reply_to_username": "user2",
            "content": "谢谢！",
            "like_count": 1,
            "created_at": "2025-12-14T14:30:00Z"
          }
        ]
      }
    ],
    "total": 10
  }
}
```

### 5.11 发表评论

```http
POST /community/comments
Authorization: Bearer {token}
```

**请求体:**
```json
{
  "post_id": 1,
  "content": "好可爱的猫咪！",
  "parent_id": 0,
  "reply_to_user_id": null
}
```

**说明:**
- `parent_id`: 0 表示一级评论，否则为回复的父评论ID
- `reply_to_user_id`: 回复指定用户时填写

### 5.12 删除评论

```http
DELETE /community/comments/:id
Authorization: Bearer {token}
```

### 5.13 点赞评论

```http
POST /community/comments/:id/like
Authorization: Bearer {token}
```

### 5.14 取消点赞评论

```http
DELETE /community/comments/:id/like
Authorization: Bearer {token}
```

---

## 六、系统接口

### 6.1 健康检查

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

### 6.2 测试接口

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
