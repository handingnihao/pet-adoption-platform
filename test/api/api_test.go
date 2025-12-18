package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pet-adoption-platform/internal/router"
	"pet-adoption-platform/pkg/database"
	"pet-adoption-platform/pkg/cache"
	"pet-adoption-platform/pkg/response"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	testRouter *gin.Engine
	testToken  string
	adminToken string
)

// setupTestRouter 初始化测试路由
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	
	// 初始化数据库和缓存（使用测试配置）
	if err := database.InitMySQL(); err != nil {
		panic("初始化数据库失败: " + err.Error())
	}
	
	if err := cache.InitRedis(); err != nil {
		panic("初始化Redis失败: " + err.Error())
	}
	
	db := database.GetDB()
	rdb := cache.GetRedis()
	
	return router.InitRouter(db, rdb)
}

// TestMain 测试入口
func TestMain(m *testing.M) {
	testRouter = setupTestRouter()
	m.Run()
}

// ========== 用户模块测试 ==========

// TestUserRegister 测试用户注册
func TestUserRegister(t *testing.T) {
	tests := []struct {
		name       string
		payload    map[string]interface{}
		wantStatus int
		wantCode   int
	}{
		{
			name: "正常注册",
			payload: map[string]interface{}{
				"username": "testuser001",
				"password": "Test@123456",
				"phone":    "13800138001",
				"email":    "test001@example.com",
			},
			wantStatus: http.StatusOK,
			wantCode:   200,
		},
		{
			name: "用户名已存在",
			payload: map[string]interface{}{
				"username": "admin",
				"password": "Test@123456",
				"phone":    "13800138002",
				"email":    "test002@example.com",
			},
			wantStatus: http.StatusBadRequest,
			wantCode:   400,
		},
		{
			name: "密码强度不够",
			payload: map[string]interface{}{
				"username": "testuser003",
				"password": "123456",
				"phone":    "13800138003",
				"email":    "test003@example.com",
			},
			wantStatus: http.StatusBadRequest,
			wantCode:   400,
		},
		{
			name: "手机号格式错误",
			payload: map[string]interface{}{
				"username": "testuser004",
				"password": "Test@123456",
				"phone":    "12345",
				"email":    "test004@example.com",
			},
			wantStatus: http.StatusBadRequest,
			wantCode:   400,
		},
		{
			name: "缺少必填参数",
			payload: map[string]interface{}{
				"password": "Test@123456",
				"phone":    "13800138005",
				"email":    "test005@example.com",
			},
			wantStatus: http.StatusBadRequest,
			wantCode:   400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest("POST", "/api/v1/users/register", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			
			w := httptest.NewRecorder()
			testRouter.ServeHTTP(w, req)
			
			assert.Equal(t, tt.wantStatus, w.Code)
			
			var resp response.Response
			json.Unmarshal(w.Body.Bytes(), &resp)
			assert.Equal(t, tt.wantCode, resp.Code)
		})
	}
}

// TestUserLogin 测试用户登录
func TestUserLogin(t *testing.T) {
	tests := []struct {
		name       string
		payload    map[string]interface{}
		wantStatus int
		wantCode   int
	}{
		{
			name: "正常登录",
			payload: map[string]interface{}{
				"username": "testuser",
				"password": "Test@123456",
			},
			wantStatus: http.StatusOK,
			wantCode:   200,
		},
		{
			name: "用户不存在",
			payload: map[string]interface{}{
				"username": "nonexistent",
				"password": "Test@123456",
			},
			wantStatus: http.StatusBadRequest,
			wantCode:   400,
		},
		{
			name: "密码错误",
			payload: map[string]interface{}{
				"username": "testuser",
				"password": "WrongPassword",
			},
			wantStatus: http.StatusBadRequest,
			wantCode:   400,
		},
		{
			name: "使用手机号登录",
			payload: map[string]interface{}{
				"username": "13800138000",
				"password": "Test@123456",
			},
			wantStatus: http.StatusOK,
			wantCode:   200,
		},
		{
			name: "使用邮箱登录",
			payload: map[string]interface{}{
				"username": "test@example.com",
				"password": "Test@123456",
			},
			wantStatus: http.StatusOK,
			wantCode:   200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest("POST", "/api/v1/users/login", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			
			w := httptest.NewRecorder()
			testRouter.ServeHTTP(w, req)
			
			assert.Equal(t, tt.wantStatus, w.Code)
			
			var resp response.Response
			json.Unmarshal(w.Body.Bytes(), &resp)
			assert.Equal(t, tt.wantCode, resp.Code)
			
			// 保存token用于后续测试
			if tt.wantCode == 200 && tt.name == "正常登录" {
				data := resp.Data.(map[string]interface{})
				testToken = data["token"].(string)
			}
		})
	}
}

// TestGetProfile 测试获取个人信息
func TestGetProfile(t *testing.T) {
	tests := []struct {
		name       string
		token      string
		wantStatus int
		wantCode   int
	}{
		{
			name:       "正常获取",
			token:      testToken,
			wantStatus: http.StatusOK,
			wantCode:   200,
		},
		{
			name:       "未登录",
			token:      "",
			wantStatus: http.StatusUnauthorized,
			wantCode:   401,
		},
		{
			name:       "Token无效",
			token:      "invalid_token",
			wantStatus: http.StatusUnauthorized,
			wantCode:   401,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/api/v1/users/profile", nil)
			if tt.token != "" {
				req.Header.Set("Authorization", "Bearer "+tt.token)
			}
			
			w := httptest.NewRecorder()
			testRouter.ServeHTTP(w, req)
			
			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

// TestUpdateProfile 测试更新个人信息
func TestUpdateProfile(t *testing.T) {
	tests := []struct {
		name       string
		token      string
		payload    map[string]interface{}
		wantStatus int
		wantCode   int
	}{
		{
			name:  "正常更新",
			token: testToken,
			payload: map[string]interface{}{
				"real_name": "张三",
				"gender":    1,
				"address":   "北京市朝阳区",
			},
			wantStatus: http.StatusOK,
			wantCode:   200,
		},
		{
			name:  "部分更新",
			token: testToken,
			payload: map[string]interface{}{
				"real_name": "李四",
			},
			wantStatus: http.StatusOK,
			wantCode:   200,
		},
		{
			name:  "未登录",
			token: "",
			payload: map[string]interface{}{
				"real_name": "王五",
			},
			wantStatus: http.StatusUnauthorized,
			wantCode:   401,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest("PUT", "/api/v1/users/profile", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			if tt.token != "" {
				req.Header.Set("Authorization", "Bearer "+tt.token)
			}
			
			w := httptest.NewRecorder()
			testRouter.ServeHTTP(w, req)
			
			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

// ========== 宠物模块测试 ==========

// TestPetList 测试宠物列表
func TestPetList(t *testing.T) {
	tests := []struct {
		name       string
		query      string
		wantStatus int
		wantCode   int
	}{
		{
			name:       "获取第一页",
			query:      "?page=1&page_size=10",
			wantStatus: http.StatusOK,
			wantCode:   200,
		},
		{
			name:       "无分页参数",
			query:      "",
			wantStatus: http.StatusOK,
			wantCode:   200,
		},
		{
			name:       "页码为0",
			query:      "?page=0&page_size=10",
			wantStatus: http.StatusOK,
			wantCode:   200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/api/v1/pets"+tt.query, nil)
			
			w := httptest.NewRecorder()
			testRouter.ServeHTTP(w, req)
			
			assert.Equal(t, tt.wantStatus, w.Code)
			
			var resp response.Response
			json.Unmarshal(w.Body.Bytes(), &resp)
			assert.Equal(t, tt.wantCode, resp.Code)
		})
	}
}

// TestPetDetail 测试宠物详情
func TestPetDetail(t *testing.T) {
	tests := []struct {
		name       string
		petID      string
		wantStatus int
	}{
		{
			name:       "正常获取",
			petID:      "1",
			wantStatus: http.StatusOK,
		},
		{
			name:       "宠物不存在",
			petID:      "99999",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "ID格式错误",
			petID:      "abc",
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/api/v1/pets/"+tt.petID, nil)
			
			w := httptest.NewRecorder()
			testRouter.ServeHTTP(w, req)
			
			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

// TestCreatePet 测试发布宠物
func TestCreatePet(t *testing.T) {
	tests := []struct {
		name       string
		token      string
		payload    map[string]interface{}
		wantStatus int
	}{
		{
			name:  "正常发布",
			token: testToken,
			payload: map[string]interface{}{
				"name":          "小白",
				"type":          "dog",
				"breed":         "金毛",
				"gender":        "male",
				"age":           12,
				"size":          "large",
				"description":   "性格温顺",
				"cover_photo":   "http://example.com/photo.jpg",
				"province":      "北京市",
				"city":          "北京市",
			},
			wantStatus: http.StatusCreated,
		},
		{
			name:  "未登录",
			token: "",
			payload: map[string]interface{}{
				"name": "小黑",
				"type": "dog",
			},
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:  "缺少必填字段",
			token: testToken,
			payload: map[string]interface{}{
				"type": "dog",
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest("POST", "/api/v1/pets", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			if tt.token != "" {
				req.Header.Set("Authorization", "Bearer "+tt.token)
			}
			
			w := httptest.NewRecorder()
			testRouter.ServeHTTP(w, req)
			
			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

// TestSearchPets 测试搜索宠物
func TestSearchPets(t *testing.T) {
	tests := []struct {
		name       string
		query      string
		wantStatus int
		wantCode   int
	}{
		{
			name:       "搜索名称",
			query:      "?keyword=小白",
			wantStatus: http.StatusOK,
			wantCode:   200,
		},
		{
			name:       "搜索品种",
			query:      "?keyword=金毛",
			wantStatus: http.StatusOK,
			wantCode:   200,
		},
		{
			name:       "空关键词",
			query:      "?keyword=",
			wantStatus: http.StatusOK,
			wantCode:   200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/api/v1/pets/search"+tt.query, nil)
			
			w := httptest.NewRecorder()
			testRouter.ServeHTTP(w, req)
			
			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

// ========== 领养模块测试 ==========

// TestCreateAdoptionApplication 测试提交领养申请
func TestCreateAdoptionApplication(t *testing.T) {
	tests := []struct {
		name       string
		token      string
		payload    map[string]interface{}
		wantStatus int
	}{
		{
			name:  "正常提交",
			token: testToken,
			payload: map[string]interface{}{
				"pet_id":           1,
				"organization_id":  1,
				"applicant_name":   "张三",
				"applicant_phone":  "13800138000",
				"applicant_address": "北京市朝阳区",
				"housing_type":     "house",
				"housing_area":     100,
				"family_members":   3,
				"family_agree":     true,
				"adoption_reason":  "喜欢宠物",
				"how_to_care":      "每天遛狗",
				"emergency_plan":   "送医院",
			},
			wantStatus: http.StatusCreated,
		},
		{
			name:  "未登录",
			token: "",
			payload: map[string]interface{}{
				"pet_id": 1,
			},
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest("POST", "/api/v1/adoptions/applications", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			if tt.token != "" {
				req.Header.Set("Authorization", "Bearer "+tt.token)
			}
			
			w := httptest.NewRecorder()
			testRouter.ServeHTTP(w, req)
			
			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

// TestGetMyApplications 测试获取我的申请
func TestGetMyApplications(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/v1/adoptions/applications/my", nil)
	req.Header.Set("Authorization", "Bearer "+testToken)
	
	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
}

// ========== 社区模块测试 ==========

// TestCreatePost 测试发布动态
func TestCreatePost(t *testing.T) {
	tests := []struct {
		name       string
		token      string
		payload    map[string]interface{}
		wantStatus int
	}{
		{
			name:  "正常发布",
			token: testToken,
			payload: map[string]interface{}{
				"title":    "我家的小狗找到新家了",
				"content":  "经过一个月的等待，终于找到了合适的家庭",
				"category": "领养故事",
			},
			wantStatus: http.StatusCreated,
		},
		{
			name:  "标题为空",
			token: testToken,
			payload: map[string]interface{}{
				"title":   "",
				"content": "内容",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name:  "未登录",
			token: "",
			payload: map[string]interface{}{
				"title":   "标题",
				"content": "内容",
			},
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest("POST", "/api/v1/community/posts", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			if tt.token != "" {
				req.Header.Set("Authorization", "Bearer "+tt.token)
			}
			
			w := httptest.NewRecorder()
			testRouter.ServeHTTP(w, req)
			
			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

// TestLikePost 测试点赞动态
func TestLikePost(t *testing.T) {
	tests := []struct {
		name       string
		token      string
		postID     string
		wantStatus int
	}{
		{
			name:       "正常点赞",
			token:      testToken,
			postID:     "1",
			wantStatus: http.StatusOK,
		},
		{
			name:       "未登录",
			token:      "",
			postID:     "1",
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/api/v1/community/posts/"+tt.postID+"/like", nil)
			if tt.token != "" {
				req.Header.Set("Authorization", "Bearer "+tt.token)
			}
			
			w := httptest.NewRecorder()
			testRouter.ServeHTTP(w, req)
			
			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

// ========== 捐赠模块测试 ==========

// TestCreateDonation 测试发起捐赠
func TestCreateDonation(t *testing.T) {
	tests := []struct {
		name       string
		token      string
		payload    map[string]interface{}
		wantStatus int
	}{
		{
			name:  "正常捐赠",
			token: testToken,
			payload: map[string]interface{}{
				"organization_id": 1,
				"amount":          100.00,
				"payment_method":  "alipay",
				"donor_name":      "张三",
				"donor_phone":     "13800138000",
				"message":         "希望能帮助更多流浪动物",
				"is_anonymous":    false,
			},
			wantStatus: http.StatusCreated,
		},
		{
			name:  "金额无效",
			token: testToken,
			payload: map[string]interface{}{
				"organization_id": 1,
				"amount":          -10,
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name:  "未登录",
			token: "",
			payload: map[string]interface{}{
				"organization_id": 1,
				"amount":          100,
			},
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest("POST", "/api/v1/donations", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			if tt.token != "" {
				req.Header.Set("Authorization", "Bearer "+tt.token)
			}
			
			w := httptest.NewRecorder()
			testRouter.ServeHTTP(w, req)
			
			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

// ========== 通用测试 ==========

// TestHealthCheck 测试健康检查
func TestHealthCheck(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
}

// TestPing 测试Ping接口
func TestPing(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/v1/ping", nil)
	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	
	var resp response.Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, 200, resp.Code)
}
