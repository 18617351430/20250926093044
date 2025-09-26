package controllers

import (
	"anti-fake-system/config"
	"anti-fake-system/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewAuthController(db *gorm.DB, cfg *config.Config) *AuthController {
	return &AuthController{db: db, cfg: cfg}
}

func (ac *AuthController) RegisterRoutes(r *gin.Engine) {
	authGroup := r.Group("/api/auth")

	handler := handlers.NewAuthHandler(ac.db, ac.cfg)

	// 平台登录
	authGroup.POST("/platform/login", handler.PlatformLogin)

	// 商户登录
	authGroup.POST("/merchant/login", handler.MerchantLogin)

	// 刷新token
	authGroup.POST("/refresh", handler.RefreshToken)
}
