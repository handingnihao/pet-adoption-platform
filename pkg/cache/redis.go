package cache

import (
	"context"
	"fmt"
	"time"

	"pet-adoption-platform/config"

	"github.com/redis/go-redis/v9"
)

// RedisClient 全局Redis客户端
var RedisClient *redis.Client

// RDB Redis客户端别名（用于兼容）
var RDB *redis.Client

// InitRedis 初始化Redis连接
func InitRedis() error {
	cfg := config.AppConfig.Redis

	// 创建Redis客户端
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         cfg.GetRedisAddr(),
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	// 设置别名
	RDB = RedisClient

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := RedisClient.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("Redis连接失败: %w", err)
	}

	return nil
}

// GetRedis 获取Redis客户端
func GetRedis() *redis.Client {
	return RedisClient
}

// Set 设置缓存
func Set(key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	return RedisClient.Set(ctx, key, value, expiration).Err()
}

// Get 获取缓存
func Get(key string) (string, error) {
	ctx := context.Background()
	return RedisClient.Get(ctx, key).Result()
}

// Del 删除缓存
func Del(keys ...string) error {
	ctx := context.Background()
	return RedisClient.Del(ctx, keys...).Err()
}

// Exists 检查key是否存在
func Exists(keys ...string) (int64, error) {
	ctx := context.Background()
	return RedisClient.Exists(ctx, keys...).Result()
}

// Expire 设置过期时间
func Expire(key string, expiration time.Duration) error {
	ctx := context.Background()
	return RedisClient.Expire(ctx, key, expiration).Err()
}

// Close 关闭Redis连接
func Close() error {
	if RedisClient != nil {
		return RedisClient.Close()
	}
	return nil
}
