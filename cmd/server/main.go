package main

import (
	"fmt"
	"log"
	"pet-adoption-platform/config"
	"pet-adoption-platform/internal/router"
	"pet-adoption-platform/pkg/cache"
	"pet-adoption-platform/pkg/database"
	"pet-adoption-platform/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 加载配置
	if err := config.Init(); err != nil {
		log.Fatalf("配置加载失败: %v", err)
	}

	// 2. 初始化日志
	if err := logger.Init(); err != nil {
		log.Fatalf("日志初始化失败: %v", err)
	}
	defer logger.Sync()

	// 3. 初始化数据库
	if err := database.InitMySQL(); err != nil {
		logger.Error("数据库连接失败: " + err.Error())
		log.Fatalf("数据库连接失败: %v", err)
	}
	logger.Info("数据库连接成功")

	// 4. 初始化Redis（可选，失败不影响启动）
	if err := cache.InitRedis(); err != nil {
		logger.Warn("Redis连接失败（部分功能可能受限）: " + err.Error())
		// 暂时允许应用启动，Redis功能（验证码等）将不可用
	} else {
		logger.Info("Redis连接成功")
	}

	// 5. 设置运行模式
	if config.AppConfig.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 6. 初始化路由（传入数据库和缓存实例）
	r := router.InitRouter(database.DB, cache.RDB)

	// 7. 启动服务器
	addr := fmt.Sprintf("%s:%d", config.AppConfig.Server.Host, config.AppConfig.Server.Port)
	logger.Info(fmt.Sprintf("服务器启动成功，监听地址: %s", addr))
	
	if err := r.Run(addr); err != nil {
		logger.Error("服务器启动失败: " + err.Error())
		log.Fatalf("服务器启动失败: %v", err)
	}
}
