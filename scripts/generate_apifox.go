package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	api := map[string]interface{}{
		"openapi": "3.0.0",
		"info": map[string]string{
			"title":       "宠物领养平台API",
			"version":     "1.0.0",
			"description": "宠物领养平台后端API接口文档\n\n测试账号：\n- 管理员: admin / Test@123456\n- 普通用户: testuser1 / Test@123456",
		},
		"servers": []map[string]string{
			{"url": "http://localhost:8080", "description": "本地开发环境"},
			{"url": "http://115.190.125.177:8080", "description": "测试服务器"},
		},
		"tags": []map[string]string{
			{"name": "系统", "description": "系统健康检查"},
			{"name": "用户认证", "description": "用户登录、注册"},
			{"name": "用户信息", "description": "个人信息管理"},
			{"name": "用户管理", "description": "管理员功能"},
		},
		"components": map[string]interface{}{
			"securitySchemes": map[string]interface{}{
				"bearerAuth": map[string]string{
					"type":         "http",
					"scheme":       "bearer",
					"bearerFormat": "JWT",
				},
			},
		},
		"paths": map[string]interface{}{
			"/health": map[string]interface{}{
				"get": map[string]interface{}{
					"tags":        []string{"系统"},
					"summary":     "健康检查",
					"description": "检查服务器运行状态",
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "成功",
						},
					},
				},
			},
			"/api/v1/users/login": map[string]interface{}{
				"post": map[string]interface{}{
					"tags":        []string{"用户认证"},
					"summary":     "用户登录",
					"description": "用户名密码登录",
					"requestBody": map[string]interface{}{
						"required": true,
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"type": "object",
									"properties": map[string]interface{}{
										"username": map[string]string{"type": "string", "example": "admin"},
										"password": map[string]string{"type": "string", "example": "Test@123456"},
									},
									"required": []string{"username", "password"},
								},
							},
						},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "登录成功",
						},
					},
				},
			},
			"/api/v1/users/register": map[string]interface{}{
				"post": map[string]interface{}{
					"tags":        []string{"用户认证"},
					"summary":     "用户注册",
					"description": "新用户注册",
					"requestBody": map[string]interface{}{
						"required": true,
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"type": "object",
									"properties": map[string]interface{}{
										"username": map[string]string{"type": "string"},
										"password": map[string]string{"type": "string"},
										"phone":    map[string]string{"type": "string"},
										"email":    map[string]string{"type": "string"},
										"code":     map[string]string{"type": "string", "description": "验证码"},
									},
								},
							},
						},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "注册成功",
						},
					},
				},
			},
			"/api/v1/users/send-code": map[string]interface{}{
				"post": map[string]interface{}{
					"tags":        []string{"用户认证"},
					"summary":     "发送验证码",
					"description": "发送手机验证码",
					"requestBody": map[string]interface{}{
						"required": true,
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"type": "object",
									"properties": map[string]interface{}{
										"phone": map[string]string{"type": "string", "example": "13900000001"},
									},
								},
							},
						},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "发送成功",
						},
					},
				},
			},
			"/api/v1/users/profile": map[string]interface{}{
				"get": map[string]interface{}{
					"tags":        []string{"用户信息"},
					"summary":     "获取个人信息",
					"description": "获取当前登录用户的个人信息",
					"security":    []map[string][]string{{"bearerAuth": {}}},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "成功",
						},
					},
				},
				"put": map[string]interface{}{
					"tags":        []string{"用户信息"},
					"summary":     "更新个人信息",
					"description": "更新当前登录用户的个人信息",
					"security":    []map[string][]string{{"bearerAuth": {}}},
					"requestBody": map[string]interface{}{
						"required": true,
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"type": "object",
									"properties": map[string]interface{}{
										"nickname":  map[string]string{"type": "string"},
										"real_name": map[string]string{"type": "string"},
										"gender":    map[string]interface{}{"type": "integer", "enum": []int{0, 1, 2}},
										"avatar":    map[string]string{"type": "string"},
									},
								},
							},
						},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "更新成功",
						},
					},
				},
			},
			"/api/v1/users/password": map[string]interface{}{
				"put": map[string]interface{}{
					"tags":        []string{"用户信息"},
					"summary":     "修改密码",
					"description": "修改当前用户密码",
					"security":    []map[string][]string{{"bearerAuth": {}}},
					"requestBody": map[string]interface{}{
						"required": true,
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"type": "object",
									"properties": map[string]interface{}{
										"old_password": map[string]string{"type": "string"},
										"new_password": map[string]string{"type": "string"},
									},
								},
							},
						},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "修改成功",
						},
					},
				},
			},
			"/api/v1/users": map[string]interface{}{
				"get": map[string]interface{}{
					"tags":        []string{"用户管理"},
					"summary":     "获取用户列表",
					"description": "管理员获取用户列表（分页）",
					"security":    []map[string][]string{{"bearerAuth": {}}},
					"parameters": []map[string]interface{}{
						{"name": "page", "in": "query", "schema": map[string]string{"type": "integer"}, "description": "页码"},
						{"name": "page_size", "in": "query", "schema": map[string]string{"type": "integer"}, "description": "每页数量"},
						{"name": "role", "in": "query", "schema": map[string]string{"type": "string"}, "description": "角色筛选"},
						{"name": "status", "in": "query", "schema": map[string]string{"type": "integer"}, "description": "状态筛选"},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "成功",
						},
					},
				},
			},
			"/api/v1/users/search": map[string]interface{}{
				"get": map[string]interface{}{
					"tags":        []string{"用户管理"},
					"summary":     "搜索用户",
					"description": "管理员搜索用户",
					"security":    []map[string][]string{{"bearerAuth": {}}},
					"parameters": []map[string]interface{}{
						{"name": "keyword", "in": "query", "required": true, "schema": map[string]string{"type": "string"}, "description": "搜索关键词"},
						{"name": "page", "in": "query", "schema": map[string]string{"type": "integer"}},
						{"name": "page_size", "in": "query", "schema": map[string]string{"type": "integer"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "成功",
						},
					},
				},
			},
			"/api/v1/users/{id}": map[string]interface{}{
				"get": map[string]interface{}{
					"tags":        []string{"用户管理"},
					"summary":     "获取用户详情",
					"description": "管理员获取指定用户详情",
					"security":    []map[string][]string{{"bearerAuth": {}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]string{"type": "integer"}, "description": "用户ID"},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "成功",
						},
					},
				},
			},
			"/api/v1/users/{id}/disable": map[string]interface{}{
				"put": map[string]interface{}{
					"tags":        []string{"用户管理"},
					"summary":     "禁用用户",
					"description": "管理员禁用指定用户",
					"security":    []map[string][]string{{"bearerAuth": {}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]string{"type": "integer"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "禁用成功",
						},
					},
				},
			},
			"/api/v1/users/{id}/enable": map[string]interface{}{
				"put": map[string]interface{}{
					"tags":        []string{"用户管理"},
					"summary":     "启用用户",
					"description": "管理员启用指定用户",
					"security":    []map[string][]string{{"bearerAuth": {}}},
					"parameters": []map[string]interface{}{
						{"name": "id", "in": "path", "required": true, "schema": map[string]string{"type": "integer"}},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "启用成功",
						},
					},
				},
			},
		},
	}

	jsonData, err := json.MarshalIndent(api, "", "  ")
	if err != nil {
		fmt.Printf("生成JSON失败: %v\n", err)
		return
	}

	err = os.WriteFile("docs/apifox-api.json", jsonData, 0644)
	if err != nil {
		fmt.Printf("写入文件失败: %v\n", err)
		return
	}

	fmt.Println("✅ API文档生成成功: docs/apifox-api.json")
	fmt.Println("📝 可以直接导入Apifox使用")
}
