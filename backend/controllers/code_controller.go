package controllers

import (
	"anti-fake-system/config"
	"anti-fake-system/handlers"
	"anti-fake-system/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CodeController struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewCodeController(db *gorm.DB, cfg *config.Config) *CodeController {
	return &CodeController{db: db, cfg: cfg}
}

func (cc *CodeController) RegisterRoutes(r *gin.Engine) {
	// 商户端防伪码管理
	merchantGroup := r.Group("/api/merchant")
	merchantGroup.Use(middleware.MerchantAuth(cc.cfg.JWT.Secret))

	merchantHandler := handlers.NewCodeHandler(cc.db, cc.cfg)

	merchantGroup.POST("/codes/generate", merchantHandler.GenerateCodes)
	merchantGroup.GET("/codes", merchantHandler.GetCodes)
	merchantGroup.GET("/codes/:id", merchantHandler.GetCodeDetail)

	// 公共验证接口
	publicGroup := r.Group("/api/public")

	verifyHandler := handlers.NewVerifyHandler(cc.db, cc.cfg)

	publicGroup.POST("/verify", verifyHandler.VerifyCode)
	publicGroup.GET("/verify/records", verifyHandler.GetVerifyRecords)
	publicGroup.GET("/verify/statistics", verifyHandler.GetVerifyStatistics)
}
