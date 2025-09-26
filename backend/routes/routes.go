package routes

import (
	"anti-fake-system/config"
	"anti-fake-system/controllers"
	"anti-fake-system/middleware"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, redisClient interface{}, cfg *config.Config) *gin.Engine {
	r := gin.Default()

	// 应用全局中间件
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":    "ok",
			"timestamp": time.Now().Unix(),
		})
	})

	// 初始化控制器
	platformController := controllers.NewPlatformController(db, cfg)
	merchantController := controllers.NewMerchantController(db, cfg)
	authController := controllers.NewAuthController(db, cfg)
	codeController := controllers.NewCodeController(db, cfg)
	ruleController := controllers.NewRuleController(db, cfg)

	// 注册路由
	platformController.RegisterRoutes(r)
	merchantController.RegisterRoutes(r)
	authController.RegisterRoutes(r)
	codeController.RegisterRoutes(r)
	ruleController.RegisterRoutes(r)

	return r
}
