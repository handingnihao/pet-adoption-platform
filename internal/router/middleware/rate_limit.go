package middleware

import (
	"fmt"
	"sync"
	"time"

	"pet-adoption-platform/config"
	"pet-adoption-platform/pkg/response"

	"github.com/gin-gonic/gin"
)

// 简单的令牌桶限流器
type tokenBucket struct {
	capacity  int       // 桶容量
	tokens    int       // 当前令牌数
	rate      int       // 每秒生成令牌数
	lastToken time.Time // 上次生成令牌时间
	mu        sync.Mutex
}

func newTokenBucket(rate, capacity int) *tokenBucket {
	return &tokenBucket{
		capacity:  capacity,
		tokens:    capacity,
		rate:      rate,
		lastToken: time.Now(),
	}
}

func (tb *tokenBucket) allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	// 计算应该生成的令牌数
	elapsed := now.Sub(tb.lastToken)
	tokensToAdd := int(elapsed.Seconds() * float64(tb.rate))

	if tokensToAdd > 0 {
		tb.tokens += tokensToAdd
		if tb.tokens > tb.capacity {
			tb.tokens = tb.capacity
		}
		tb.lastToken = now
	}

	// 检查是否有令牌
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}

	return false
}

// 全局限流器
var (
	ipLimiters   = make(map[string]*tokenBucket)
	userLimiters = make(map[int64]*tokenBucket)
	limiterMutex sync.RWMutex
)

// getIPLimiter 获取IP限流器
func getIPLimiter(ip string) *tokenBucket {
	limiterMutex.Lock()
	defer limiterMutex.Unlock()

	limiter, exists := ipLimiters[ip]
	if !exists {
		limit := config.AppConfig.RateLimit.IPLimit
		limiter = newTokenBucket(limit, limit*2)
		ipLimiters[ip] = limiter
	}

	return limiter
}

// getUserLimiter 获取用户限流器
func getUserLimiter(userID int64) *tokenBucket {
	limiterMutex.Lock()
	defer limiterMutex.Unlock()

	limiter, exists := userLimiters[userID]
	if !exists {
		limit := config.AppConfig.RateLimit.UserLimit
		limiter = newTokenBucket(limit, limit*2)
		userLimiters[userID] = limiter
	}

	return limiter
}

// RateLimit IP限流中间件
func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查是否启用限流
		if !config.AppConfig.RateLimit.Enabled {
			c.Next()
			return
		}

		ip := c.ClientIP()
		limiter := getIPLimiter(ip)

		if !limiter.allow() {
			response.TooManyRequests(c, "请求过于频繁，请稍后再试")
			c.Abort()
			return
		}

		c.Next()
	}
}

// UserRateLimit 用户限流中间件（需要在Auth中间件之后使用）
func UserRateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查是否启用限流
		if !config.AppConfig.RateLimit.Enabled {
			c.Next()
			return
		}

		// 获取用户ID
		userID, exists := c.Get("user_id")
		if !exists {
			c.Next()
			return
		}

		limiter := getUserLimiter(userID.(int64))

		if !limiter.allow() {
			response.TooManyRequests(c, "操作过于频繁，请稍后再试")
			c.Abort()
			return
		}

		c.Next()
	}
}

// 定期清理过期的限流器（防止内存泄漏）
func init() {
	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			limiterMutex.Lock()
			// 清理IP限流器
			if len(ipLimiters) > 10000 {
				ipLimiters = make(map[string]*tokenBucket)
			}
			// 清理用户限流器
			if len(userLimiters) > 10000 {
				userLimiters = make(map[int64]*tokenBucket)
			}
			limiterMutex.Unlock()
			fmt.Println("限流器缓存已清理")
		}
	}()
}
