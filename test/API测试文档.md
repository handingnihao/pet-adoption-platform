# 爱心宠物领养平台 - API接口测试文档

## 1. 测试环境

### 1.1 测试地址
- 本地环境: `http://localhost:8080`
- 测试环境: `http://test.pet-adoption.com`
- 生产环境: `http://api.pet-adoption.com`

### 1.2 测试工具
- Postman / Apifox
- curl 命令行
- Go httptest

### 1.3 测试账号

| 角色 | 用户名 | 密码 | 说明 |
|------|--------|------|------|
| 管理员 | admin | Test@123456 | 管理员账号 |
| 普通用户 | testuser | Test@123456 | 普通用户账号 |
| 机构用户 | testorg | Test@123456 | 机构用户账号 |

---

## 2. 用户模块接口测试

### 2.1 用户注册

**接口**: `POST /api/v1/users/register`

**请求示例**:
```bash
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "newuser",
    "password": "Test@123456",
    "phone": "13800138000",
    "email": "newuser@example.com"
  }'
```

**测试用例**:

| 用例ID | 测试场景 | 请求参数 | 期望结果 | 状态码 |
|--------|---------|---------|---------|--------|
| TC001 | 正常注册 | 完整有效参数 | 注册成功 | 200 |
| TC002 | 用户名已存在 | username: "admin" | 提示用户名已存在 | 400 |
| TC003 | 手机号已注册 | phone: "13800138001" | 提示手机号已注册 | 400 |
| TC004 | 邮箱已注册 | email: "admin@example.com" | 提示邮箱已注册 | 400 |
| TC005 | 密码强度不够 | password: "123456" | 提示密码强度不够 | 400 |
| TC006 | 手机号格式错误 | phone: "12345" | 提示手机号格式错误 | 400 |
| TC007 | 邮箱格式错误 | email: "invalid-email" | 提示邮箱格式错误 | 400 |
| TC008 | 缺少必填参数 | 缺少username | 提示参数错误 | 400 |

**响应示例**:
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "user_id": 123,
    "message": "注册成功"
  },
  "timestamp": 1702800000
}
```

---

### 2.2 用户登录

**接口**: `POST /api/v1/users/login`

**请求示例**:
```bash
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "Test@123456"
  }'
```

**测试用例**:

| 用例ID | 测试场景 | 请求参数 | 期望结果 | 状态码 |
|--------|---------|---------|---------|--------|
| TC101 | 正常登录 | 正确的用户名密码 | 返回token | 200 |
| TC102 | 用户不存在 | username: "nonexistent" | 提示用户不存在 | 400 |
| TC103 | 密码错误 | password: "wrongpassword" | 提示密码错误 | 400 |
| TC104 | 用户被禁用 | 被禁用的用户 | 提示用户已被禁用 | 403 |
| TC105 | 使用手机号登录 | username: "13800138000" | 登录成功 | 200 |
| TC106 | 使用邮箱登录 | username: "test@example.com" | 登录成功 | 200 |

**响应示例**:
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_at": "2024-12-18T10:00:00Z",
    "user": {
      "id": 1,
      "username": "testuser",
      "email": "test@example.com",
      "role": "user"
    }
  },
  "timestamp": 1702800000
}
```

---

### 2.3 获取个人信息

**接口**: `GET /api/v1/users/profile`

**请求示例**:
```bash
curl -X GET http://localhost:8080/api/v1/users/profile \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**测试用例**:

| 用例ID | 测试场景 | 请求头 | 期望结果 | 状态码 |
|--------|---------|--------|---------|--------|
| TC201 | 正常获取 | 有效token | 返回用户信息 | 200 |
| TC202 | 未登录 | 无token | 提示请先登录 | 401 |
| TC203 | token过期 | 过期的token | 提示token过期 | 401 |
| TC204 | token无效 | 无效的token | 提示token无效 | 401 |

---

### 2.4 更新个人信息

**接口**: `PUT /api/v1/users/profile`

**请求示例**:
```bash
curl -X PUT http://localhost:8080/api/v1/users/profile \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "real_name": "张三",
    "avatar": "http://example.com/avatar.jpg",
    "gender": 1,
    "address": "北京市朝阳区"
  }'
```

**测试用例**:

| 用例ID | 测试场景 | 请求参数 | 期望结果 | 状态码 |
|--------|---------|---------|---------|--------|
| TC301 | 正常更新 | 有效参数 | 更新成功 | 200 |
| TC302 | 未登录 | 无token | 提示请先登录 | 401 |
| TC303 | 部分更新 | 只更新real_name | 更新成功 | 200 |

---

### 2.5 修改密码

**接口**: `PUT /api/v1/users/password`

**请求示例**:
```bash
curl -X PUT http://localhost:8080/api/v1/users/password \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "old_password": "Test@123456",
    "new_password": "NewTest@123456"
  }'
```

**测试用例**:

| 用例ID | 测试场景 | 请求参数 | 期望结果 | 状态码 |
|--------|---------|---------|---------|--------|
| TC401 | 正常修改 | 正确的旧密码 | 修改成功 | 200 |
| TC402 | 旧密码错误 | 错误的旧密码 | 提示原密码错误 | 400 |
| TC403 | 新密码强度不够 | 弱密码 | 提示密码强度不够 | 400 |
| TC404 | 未登录 | 无token | 提示请先登录 | 401 |

---

## 3. 宠物模块接口测试

### 3.1 宠物列表

**接口**: `GET /api/v1/pets`

**请求示例**:
```bash
curl -X GET "http://localhost:8080/api/v1/pets?page=1&page_size=10"
```

**测试用例**:

| 用例ID | 测试场景 | 请求参数 | 期望结果 | 状态码 |
|--------|---------|---------|---------|--------|
| TC501 | 获取第一页 | page=1, page_size=10 | 返回10条数据 | 200 |
| TC502 | 获取空页 | page=999 | 返回空数组 | 200 |
| TC503 | 无分页参数 | 无参数 | 使用默认分页 | 200 |

---

### 3.2 宠物详情

**接口**: `GET /api/v1/pets/:id`

**请求示例**:
```bash
curl -X GET http://localhost:8080/api/v1/pets/1
```

**测试用例**:

| 用例ID | 测试场景 | 路径参数 | 期望结果 | 状态码 |
|--------|---------|---------|---------|--------|
| TC601 | 正常获取 | id=1 | 返回宠物详情 | 200 |
| TC602 | 宠物不存在 | id=99999 | 提示宠物不存在 | 404 |
| TC603 | ID格式错误 | id=abc | 提示参数错误 | 400 |

---

### 3.3 发布宠物

**接口**: `POST /api/v1/pets`

**请求示例**:
```bash
curl -X POST http://localhost:8080/api/v1/pets \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "小白",
    "type": "dog",
    "breed": "金毛",
    "gender": "male",
    "age": 12,
    "size": "large",
    "color": "金色",
    "weight": 30.5,
    "is_vaccinated": true,
    "is_sterilized": false,
    "health_status": "健康",
    "description": "性格温顺，喜欢和人玩耍",
    "character": "活泼、友好",
    "photos": ["http://example.com/photo1.jpg"],
    "cover_photo": "http://example.com/cover.jpg",
    "province": "北京市",
    "city": "北京市",
    "district": "朝阳区",
    "address": "朝阳公园附近"
  }'
```

**测试用例**:

| 用例ID | 测试场景 | 请求参数 | 期望结果 | 状态码 |
|--------|---------|---------|---------|--------|
| TC701 | 正常发布 | 完整有效参数 | 发布成功 | 201 |
| TC702 | 未登录 | 无token | 提示请先登录 | 401 |
| TC703 | 缺少必填字段 | 缺少name | 提示参数错误 | 400 |
| TC704 | 类型错误 | type: "invalid" | 提示类型错误 | 400 |

---

### 3.4 搜索宠物

**接口**: `GET /api/v1/pets/search`

**请求示例**:
```bash
curl -X GET "http://localhost:8080/api/v1/pets/search?keyword=金毛&page=1&page_size=10"
```

**测试用例**:

| 用例ID | 测试场景 | 请求参数 | 期望结果 | 状态码 |
|--------|---------|---------|---------|--------|
| TC801 | 搜索名称 | keyword=小白 | 返回匹配结果 | 200 |
| TC802 | 搜索品种 | keyword=金毛 | 返回匹配结果 | 200 |
| TC803 | 无结果 | keyword=不存在的宠物 | 返回空数组 | 200 |
| TC804 | 空关键词 | keyword= | 返回所有宠物 | 200 |

---

## 4. 领养模块接口测试

### 4.1 提交领养申请

**接口**: `POST /api/v1/adoptions/applications`

**请求示例**:
```bash
curl -X POST http://localhost:8080/api/v1/adoptions/applications \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "pet_id": 1,
    "organization_id": 1,
    "applicant_name": "张三",
    "applicant_phone": "13800138000",
    "applicant_address": "北京市朝阳区",
    "housing_type": "house",
    "housing_area": 100,
    "has_yard": true,
    "family_members": 3,
    "has_children": false,
    "family_agree": true,
    "adoption_reason": "喜欢宠物，想给它一个温暖的家",
    "how_to_care": "每天遛狗，定期体检，提供良好的生活环境",
    "emergency_plan": "如遇紧急情况，会立即送往宠物医院"
  }'
```

**测试用例**:

| 用例ID | 测试场景 | 请求参数 | 期望结果 | 状态码 |
|--------|---------|---------|---------|--------|
| TC901 | 正常提交 | 完整有效参数 | 提交成功 | 201 |
| TC902 | 未登录 | 无token | 提示请先登录 | 401 |
| TC903 | 宠物不存在 | pet_id: 99999 | 提示宠物不存在 | 404 |
| TC904 | 宠物已被领养 | 已领养的宠物 | 提示宠物已被领养 | 400 |
| TC905 | 重复申请 | 已申请过的宠物 | 提示已申请过 | 400 |

---

### 4.2 我的申请列表

**接口**: `GET /api/v1/adoptions/applications/my`

**请求示例**:
```bash
curl -X GET "http://localhost:8080/api/v1/adoptions/applications/my?page=1&page_size=10" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**测试用例**:

| 用例ID | 测试场景 | 请求参数 | 期望结果 | 状态码 |
|--------|---------|---------|---------|--------|
| TC1001 | 正常获取 | page=1 | 返回申请列表 | 200 |
| TC1002 | 未登录 | 无token | 提示请先登录 | 401 |
| TC1003 | 无申请记录 | 新用户 | 返回空数组 | 200 |

---

### 4.3 审核申请（管理员）

**接口**: `PUT /api/v1/adoptions/applications/:id/review`

**请求示例**:
```bash
curl -X PUT http://localhost:8080/api/v1/adoptions/applications/1/review \
  -H "Authorization: Bearer ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "action": "approve",
    "comment": "申请材料齐全，符合领养条件"
  }'
```

**测试用例**:

| 用例ID | 测试场景 | 请求参数 | 期望结果 | 状态码 |
|--------|---------|---------|---------|--------|
| TC1101 | 审核通过 | action: "approve" | 审核成功 | 200 |
| TC1102 | 审核拒绝 | action: "reject" | 审核成功 | 200 |
| TC1103 | 安排面试 | action: "interview" | 安排成功 | 200 |
| TC1104 | 非管理员 | 普通用户token | 提示权限不足 | 403 |
| TC1105 | 申请不存在 | id: 99999 | 提示申请不存在 | 404 |

---

## 5. 机构模块接口测试

### 5.1 机构注册

**接口**: `POST /api/v1/organizations/register`

**请求示例**:
```bash
curl -X POST http://localhost:8080/api/v1/organizations/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "爱心救助站",
    "type": "救助站",
    "logo": "http://example.com/logo.jpg",
    "description": "专注流浪动物救助",
    "province": "北京市",
    "city": "北京市",
    "address": "朝阳区某某街道",
    "contact_name": "李四",
    "contact_phone": "13800138001",
    "contact_email": "contact@example.com",
    "phone": "010-12345678",
    "email": "org@example.com",
    "credential_urls": ["http://example.com/cert1.jpg"]
  }'
```

**测试用例**:

| 用例ID | 测试场景 | 请求参数 | 期望结果 | 状态码 |
|--------|---------|---------|---------|--------|
| TC1201 | 正常注册 | 完整有效参数 | 注册成功 | 201 |
| TC1202 | 机构名已存在 | 重复的机构名 | 提示机构名已存在 | 400 |
| TC1203 | 缺少资质证明 | 无credential_urls | 提示上传资质证明 | 400 |

---

### 5.2 机构列表

**接口**: `GET /api/v1/organizations`

**请求示例**:
```bash
curl -X GET "http://localhost:8080/api/v1/organizations?page=1&page_size=10"
```

**测试用例**:

| 用例ID | 测试场景 | 请求参数 | 期望结果 | 状态码 |
|--------|---------|---------|---------|--------|
| TC1301 | 获取列表 | page=1 | 返回机构列表 | 200 |
| TC1302 | 筛选已认证 | status=1 | 返回已认证机构 | 200 |

---

## 6. 社区模块接口测试

### 6.1 发布动态

**接口**: `POST /api/v1/community/posts`

**请求示例**:
```bash
curl -X POST http://localhost:8080/api/v1/community/posts \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "我家的小狗找到新家了",
    "content": "经过一个月的等待，终于为小狗找到了合适的领养家庭",
    "images": ["http://example.com/img1.jpg"],
    "category": "领养故事"
  }'
```

**测试用例**:

| 用例ID | 测试场景 | 请求参数 | 期望结果 | 状态码 |
|--------|---------|---------|---------|--------|
| TC1401 | 正常发布 | 完整参数 | 发布成功 | 201 |
| TC1402 | 未登录 | 无token | 提示请先登录 | 401 |
| TC1403 | 标题为空 | title: "" | 提示标题不能为空 | 400 |
| TC1404 | 内容过长 | 超长内容 | 提示内容过长 | 400 |

---

### 6.2 动态列表

**接口**: `GET /api/v1/community/posts`

**请求示例**:
```bash
curl -X GET "http://localhost:8080/api/v1/community/posts?page=1&page_size=10"
```

---

### 6.3 点赞动态

**接口**: `POST /api/v1/community/posts/:id/like`

**请求示例**:
```bash
curl -X POST http://localhost:8080/api/v1/community/posts/1/like \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**测试用例**:

| 用例ID | 测试场景 | 路径参数 | 期望结果 | 状态码 |
|--------|---------|---------|---------|--------|
| TC1501 | 正常点赞 | id=1 | 点赞成功 | 200 |
| TC1502 | 重复点赞 | 已点赞的动态 | 提示已点赞 | 400 |
| TC1503 | 未登录 | 无token | 提示请先登录 | 401 |

---

## 7. 捐赠模块接口测试

### 7.1 发起捐赠

**接口**: `POST /api/v1/donations`

**请求示例**:
```bash
curl -X POST http://localhost:8080/api/v1/donations \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "organization_id": 1,
    "amount": 100.00,
    "payment_method": "alipay",
    "donor_name": "张三",
    "donor_phone": "13800138000",
    "message": "希望能帮助更多流浪动物",
    "is_anonymous": false
  }'
```

**测试用例**:

| 用例ID | 测试场景 | 请求参数 | 期望结果 | 状态码 |
|--------|---------|---------|---------|--------|
| TC1601 | 正常捐赠 | 有效参数 | 创建成功 | 201 |
| TC1602 | 金额无效 | amount: -10 | 提示金额无效 | 400 |
| TC1603 | 机构不存在 | organization_id: 99999 | 提示机构不存在 | 404 |
| TC1604 | 未登录 | 无token | 提示请先登录 | 401 |

---

### 7.2 我的捐赠记录

**接口**: `GET /api/v1/donations/my`

**请求示例**:
```bash
curl -X GET "http://localhost:8080/api/v1/donations/my?page=1&page_size=10" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

---

## 8. 通用测试场景

### 8.1 认证测试

| 场景 | 测试内容 | 期望结果 |
|------|---------|---------|
| 无Token | 访问需要认证的接口 | 401 Unauthorized |
| Token过期 | 使用过期的Token | 401 Token过期 |
| Token无效 | 使用伪造的Token | 401 Token无效 |
| 权限不足 | 普通用户访问管理员接口 | 403 Forbidden |

### 8.2 参数验证测试

| 场景 | 测试内容 | 期望结果 |
|------|---------|---------|
| 必填参数缺失 | 不传必填参数 | 400 参数错误 |
| 参数类型错误 | 传入错误类型 | 400 参数类型错误 |
| 参数格式错误 | 格式不符合要求 | 400 参数格式错误 |
| 参数长度超限 | 超过最大长度 | 400 参数长度超限 |

### 8.3 分页测试

| 场景 | 测试内容 | 期望结果 |
|------|---------|---------|
| 默认分页 | 不传分页参数 | 使用默认值 |
| 页码为0 | page=0 | 自动转为1 |
| 页码为负数 | page=-1 | 自动转为1 |
| 页大小超限 | page_size=1000 | 限制为最大值 |

---

## 9. 性能测试指标

| 接口 | QPS要求 | 响应时间(P95) | 响应时间(P99) |
|------|---------|---------------|---------------|
| 用户登录 | 1000+ | <100ms | <200ms |
| 宠物列表 | 2000+ | <50ms | <100ms |
| 宠物详情 | 3000+ | <30ms | <50ms |
| 发布宠物 | 500+ | <200ms | <500ms |
| 提交申请 | 500+ | <200ms | <500ms |

---

## 10. 测试报告模板

### 测试执行记录

| 日期 | 测试人员 | 测试环境 | 总用例数 | 通过数 | 失败数 | 通过率 |
|------|---------|---------|---------|--------|--------|--------|
| 2024-12-17 | 测试员 | 测试环境 | 100 | 95 | 5 | 95% |

### 缺陷统计

| 缺陷ID | 接口 | 严重程度 | 状态 | 描述 |
|--------|------|---------|------|------|
| BUG001 | 用户注册 | 高 | 已修复 | 手机号验证不严格 |
| BUG002 | 宠物搜索 | 中 | 待修复 | 搜索结果排序错误 |

---

**文档版本**: v1.0  
**创建日期**: 2024-12-17  
**最后更新**: 2024-12-17
