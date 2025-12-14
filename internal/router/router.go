package router

import (
	"net/http"

	"pet-adoption-platform/internal/controller"
	"pet-adoption-platform/internal/router/middleware"
	"pet-adoption-platform/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// InitRouter 初始化路由
func InitRouter(db *gorm.DB, rdb *redis.Client) *gin.Engine {
	r := gin.New()

	// 使用中间件
	r.Use(middleware.Logger())    // 日志中间件
	r.Use(middleware.CORS())      // 跨域中间件
	r.Use(middleware.RateLimit()) // 限流中间件
	r.Use(gin.Recovery())         // 恢复中间件

	// 创建控制器实例
	userCtrl := controller.NewUserController(db, rdb)
	petCtrl := controller.NewPetController(db, rdb)
	adoptionCtrl := controller.NewAdoptionController(db)
	orgCtrl := controller.NewOrganizationController(db)
	communityCtrl := controller.NewCommunityController(db)

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "服务运行正常",
		})
	})

	// API v1版本
	v1 := r.Group("/api/v1")
	{
		// 测试接口
		v1.GET("/ping", func(c *gin.Context) {
			response.Success(c, gin.H{
				"message": "pong",
			})
		})

		// 用户相关路由（公开接口，无需登录）
		users := v1.Group("/users")
		{
			// 公开接口
			users.POST("/register", userCtrl.Register) // 用户注册
			users.POST("/login", userCtrl.Login)       // 用户登录

			// 需要登录的接口
			usersAuth := users.Group("")
			usersAuth.Use(middleware.Auth())
			{
				usersAuth.GET("/profile", userCtrl.GetProfile)      // 获取个人信息
				usersAuth.PUT("/profile", userCtrl.UpdateProfile)   // 更新个人信息
				usersAuth.PUT("/password", userCtrl.UpdatePassword) // 修改密码
			}

			// 管理员接口
			usersAdmin := users.Group("")
			usersAdmin.Use(middleware.Auth(), middleware.AdminAuth())
			{
				usersAdmin.GET("", userCtrl.GetUserList)             // 获取用户列表
				usersAdmin.GET("/search", userCtrl.SearchUsers)      // 搜索用户
				usersAdmin.GET("/:id", userCtrl.GetUserByID)         // 获取用户详情
				usersAdmin.PUT("/:id/disable", userCtrl.DisableUser) // 禁用用户
				usersAdmin.PUT("/:id/enable", userCtrl.EnableUser)   // 启用用户
			}
		}

		// 宠物相关路由
		pets := v1.Group("/pets")
		{
			// 公开接口（无需登录）
			pets.GET("", petCtrl.ListPets)                       // 宠物列表
			pets.GET("/query", petCtrl.QueryPets)                // 条件查询
			pets.GET("/search", petCtrl.SearchPets)              // 搜索宠物
			pets.GET("/recommended", petCtrl.GetRecommendedPets) // 推荐宠物
			pets.GET("/:id", petCtrl.GetPetDetail)               // 宠物详情

			// 需要登录的接口
			petsAuth := pets.Group("")
			petsAuth.Use(middleware.Auth())
			{
				petsAuth.POST("", petCtrl.CreatePet)             // 发布宠物
				petsAuth.GET("/my", petCtrl.GetMyPets)           // 我的宠物
				petsAuth.PUT("/:id", petCtrl.UpdatePet)          // 更新宠物
				petsAuth.DELETE("/:id", petCtrl.DeletePet)       // 删除宠物
				petsAuth.PUT("/:id/offline", petCtrl.OfflinePet) // 下架宠物
			}

			// 管理员接口
			petsAdmin := pets.Group("")
			petsAdmin.Use(middleware.Auth(), middleware.AdminAuth())
			{
				petsAdmin.GET("/pending", petCtrl.GetPendingPets)   // 待审核列表
				petsAdmin.PUT("/:id/approve", petCtrl.ApprovePet)   // 审核通过
				petsAdmin.PUT("/:id/reject", petCtrl.RejectPet)     // 拒绝
				petsAdmin.GET("/statistics", petCtrl.GetStatistics) // 统计信息
			}
		}

		// 领养相关路由
		adoptions := v1.Group("/adoptions")
		{
			// 需要登录的接口
			adoptionsAuth := adoptions.Group("")
			adoptionsAuth.Use(middleware.Auth())
			{
				// 申请相关
				adoptionsAuth.POST("/applications", adoptionCtrl.CreateApplication)       // 提交领养申请
				adoptionsAuth.GET("/applications/my", adoptionCtrl.GetMyApplications)     // 我的申请列表
				adoptionsAuth.GET("/applications/:id", adoptionCtrl.GetApplicationByID)   // 申请详情
				adoptionsAuth.PUT("/applications/:id", adoptionCtrl.UpdateApplication)    // 更新申请
				adoptionsAuth.DELETE("/applications/:id", adoptionCtrl.CancelApplication) // 取消申请

				// 领养记录相关
				adoptionsAuth.GET("/records/my", adoptionCtrl.GetMyAdoptions)   // 我的领养记录
				adoptionsAuth.GET("/records/:id", adoptionCtrl.GetAdoptionByID) // 领养记录详情
			}

			// 管理员接口
			adoptionsAdmin := adoptions.Group("")
			adoptionsAdmin.Use(middleware.Auth(), middleware.AdminAuth())
			{
				// 申请管理
				adoptionsAdmin.GET("/applications", adoptionCtrl.ListApplications)               // 所有申请列表
				adoptionsAdmin.GET("/applications/pending", adoptionCtrl.GetPendingApplications) // 待审核申请
				adoptionsAdmin.GET("/applications/query", adoptionCtrl.QueryApplications)        // 条件查询申请
				adoptionsAdmin.PUT("/applications/:id/review", adoptionCtrl.ReviewApplication)   // 审核申请

				// 领养记录管理
				adoptionsAdmin.POST("/records", adoptionCtrl.CreateAdoption)                 // 创建领养记录
				adoptionsAdmin.GET("/records", adoptionCtrl.ListAdoptions)                   // 所有领养记录
				adoptionsAdmin.PUT("/records/:id/status", adoptionCtrl.UpdateAdoptionStatus) // 更新状态

				// 统计信息
				adoptionsAdmin.GET("/statistics", adoptionCtrl.GetStatistics) // 统计信息
			}
		}

		// 机构相关路由
		organizations := v1.Group("/organizations")
		{
			// 公开接口（无需登录）
			organizations.GET("", orgCtrl.ListOrganizations)       // 机构列表
			organizations.GET("/:id", orgCtrl.GetOrganization)     // 机构详情

			// 需要登录的接口
			orgsAuth := organizations.Group("")
			orgsAuth.Use(middleware.Auth())
			{
				orgsAuth.POST("", orgCtrl.CreateOrganization)      // 申请入驻/创建机构
				orgsAuth.GET("/my", orgCtrl.GetMyOrganizations)    // 我创建的机构
				orgsAuth.PUT("/:id", orgCtrl.UpdateOrganization)   // 更新机构信息
				orgsAuth.DELETE("/:id", orgCtrl.DeleteOrganization) // 删除机构
			}

			// 管理员接口
			orgsAdmin := organizations.Group("")
			orgsAdmin.Use(middleware.Auth(), middleware.AdminAuth())
			{
				orgsAdmin.PUT("/:id/status", orgCtrl.UpdateOrganizationStatus) // 审核机构（通过/拒绝）
			}
		}

		// 社区相关路由
		community := v1.Group("/community")
		{
			// 公开接口（无需登录）
			community.GET("/posts", communityCtrl.ListPosts)              // 动态列表
			community.GET("/posts/search", communityCtrl.SearchPosts)     // 搜索动态
			community.GET("/posts/:id", communityCtrl.GetPost)            // 动态详情
			community.GET("/posts/:id/comments", communityCtrl.GetPostComments) // 动态评论列表

			// 需要登录的接口
			communityAuth := community.Group("")
			communityAuth.Use(middleware.Auth())
			{
				// 动态相关
				communityAuth.POST("/posts", communityCtrl.CreatePost)           // 发布动态
				communityAuth.GET("/posts/my", communityCtrl.GetMyPosts)         // 我的动态
				communityAuth.PUT("/posts/:id", communityCtrl.UpdatePost)        // 更新动态
				communityAuth.DELETE("/posts/:id", communityCtrl.DeletePost)     // 删除动态
				communityAuth.POST("/posts/:id/like", communityCtrl.LikePost)    // 点赞动态
				communityAuth.DELETE("/posts/:id/like", communityCtrl.UnlikePost) // 取消点赞

				// 评论相关
				communityAuth.POST("/comments", communityCtrl.CreateComment)           // 发表评论
				communityAuth.DELETE("/comments/:id", communityCtrl.DeleteComment)     // 删除评论
				communityAuth.POST("/comments/:id/like", communityCtrl.LikeComment)    // 点赞评论
				communityAuth.DELETE("/comments/:id/like", communityCtrl.UnlikeComment) // 取消点赞评论
			}
		}

		// 捐赠相关路由（未实现）
		// donations := v1.Group("/donations")
		// donations.Use(middleware.Auth())
		// {
		// 	// donations.POST("", controller.CreateDonation)      // 创建捐赠
		// 	// donations.GET("", controller.GetDonationList)      // 捐赠列表
		// 	// donations.POST("/:id/pay", controller.PayDonation) // 支付捐赠
		// }

		// 文件上传路由（未实现）
		// files := v1.Group("/files")
		// files.Use(middleware.Auth())
		// {
		// 	// files.POST("/image", controller.UploadImage)  // 上传图片
		// 	// files.POST("/video", controller.UploadVideo)  // 上传视频
		// }

		// 管理后台路由（未实现）
		// admin := v1.Group("/admin")
		// admin.Use(middleware.Auth(), middleware.AdminAuth())
		// {
		// 	// admin.GET("/dashboard", controller.GetDashboard) // 仪表盘
		// 	// admin.GET("/stats", controller.GetStats)         // 统计数据
		// }
	}

	return r
}
