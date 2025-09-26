package controllers

import (
	"anti-fake-system/config"
	"anti-fake-system/handlers"
	"anti-fake-system/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RuleController struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewRuleController(db *gorm.DB, cfg *config.Config) *RuleController {
	return &RuleController{db: db, cfg: cfg}
}

func (rc *RuleController) RegisterRoutes(r *gin.Engine) {
	merchantGroup := r.Group("/api/merchant")
	merchantGroup.Use(middleware.MerchantAuth(rc.cfg.JWT.Secret))

	handler := handlers.NewRuleHandler(rc.db, rc.cfg)

	// 规则管理
	merchantGroup.POST("/rules", handler.CreateRule)
	merchantGroup.PUT("/rules/:id", handler.UpdateRule)
	merchantGroup.GET("/rules/:id", handler.GetRuleDetail)
	merchantGroup.POST("/rules/:id/test", handler.TestRule)
	merchantGroup.POST("/rules/validate", handler.ValidateRuleConfig)
}
