package middleware

import (
	"strings"

	"pet-adoption-platform/pkg/response"
	"pet-adoption-platform/pkg/utils"

	"github.com/gin-gonic/gin"
)

// Auth JWT认证中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Header获取Token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Unauthorized(c, "请先登录")
			c.Abort()
			return
		}

		// Token格式：Bearer <token>
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Unauthorized(c, "Token格式错误")
			c.Abort()
			return
		}

		// 解析Token
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			response.Unauthorized(c, "Token无效或已过期")
			c.Abort()
			return
		}

		// 将用户信息保存到上下文
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}

// AdminAuth 管理员权限验证
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			response.Forbidden(c, "权限不足")
			c.Abort()
			return
		}

		if role != "admin" {
			response.Forbidden(c, "需要管理员权限")
			c.Abort()
			return
		}

		c.Next()
	}
}

// OrganizationAuth 机构权限验证
func OrganizationAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			response.Forbidden(c, "权限不足")
			c.Abort()
			return
		}

		roleStr := role.(string)
		if roleStr != "admin" && roleStr != "organization" {
			response.Forbidden(c, "需要机构权限")
			c.Abort()
			return
		}

		c.Next()
	}
}
