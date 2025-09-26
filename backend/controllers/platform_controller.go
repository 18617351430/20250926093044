package controllers

import (
	"anti-fake-system/config"
	"anti-fake-system/handlers"
	"anti-fake-system/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PlatformController struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewPlatformController(db *gorm.DB, cfg *config.Config) *PlatformController {
	return &PlatformController{db: db, cfg: cfg}
}

func (pc *PlatformController) RegisterRoutes(r *gin.Engine) {
	platformGroup := r.Group("/api/platform")
	platformGroup.Use(middleware.PlatformAuth())

	handler := handlers.NewPlatformHandler(pc.db, pc.cfg)

	// 商户管理
	platformGroup.GET("/merchants", handler.GetMerchants)
	platformGroup.POST("/merchants", handler.CreateMerchant)
	platformGroup.PUT("/merchants/:id", handler.UpdateMerchant)
	platformGroup.GET("/merchants/:id", handler.GetMerchantDetail)

	// 统计报表
	platformGroup.GET("/statistics", handler.GetStatistics)
	platformGroup.GET("/verify-trend", handler.GetVerifyTrend)
}
