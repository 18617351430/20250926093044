package controllers

import (
	"anti-fake-system/config"
	"anti-fake-system/handlers"
	"anti-fake-system/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MerchantController struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewMerchantController(db *gorm.DB, cfg *config.Config) *MerchantController {
	return &MerchantController{db: db, cfg: cfg}
}

func (mc *MerchantController) RegisterRoutes(r *gin.Engine) {
	merchantGroup := r.Group("/api/merchant")
	merchantGroup.Use(middleware.MerchantAuth(mc.cfg.JWT.Secret))

	handler := handlers.NewMerchantHandler(mc.db, mc.cfg)

	// 商品管理
	merchantGroup.GET("/products", handler.GetProducts)
	merchantGroup.POST("/products", handler.CreateProduct)

	// 批次管理
	merchantGroup.GET("/batches", handler.GetBatches)
	merchantGroup.POST("/batches", handler.CreateBatch)

	// 规则管理
	merchantGroup.GET("/rules", handler.GetRules)
	merchantGroup.POST("/rules", handler.CreateRule)

	// 商户统计
	merchantGroup.GET("/statistics", handler.GetMerchantStatistics)
}
