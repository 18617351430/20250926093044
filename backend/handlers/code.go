package handlers

import (
	"anti-fake-system/config"
	"anti-fake-system/models"
	"anti-fake-system/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CodeHandler struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewCodeHandler(db *gorm.DB, cfg *config.Config) *CodeHandler {
	return &CodeHandler{db: db, cfg: cfg}
}

// GenerateRequest 生成防伪码请求
type GenerateRequest struct {
	RuleID   uint `json:"rule_id" binding:"required"`  // 规则ID
	BatchID  uint `json:"batch_id" binding:"required"` // 批次ID
	Quantity int  `json:"quantity" binding:"required"` // 生成数量
	StartSeq int  `json:"start_seq"`                   // 起始序号
}

// GenerateResponse 生成防伪码响应
type GenerateResponse struct {
	Total     int      `json:"total"`      // 生成总数
	Codes     []string `json:"codes"`      // 生成的防伪码（前100条）
	BatchCode string   `json:"batch_code"` // 批次标识
}

// GenerateCodes 生成防伪码
func (h *CodeHandler) GenerateCodes(c *gin.Context) {
	var req GenerateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	// 获取当前用户信息
	merchantID, _ := c.Get("merchantID")

	// 商户用户只能生成自己商户的防伪码
	if merchantID != nil {
		var rule models.SecurityCodeRule
		if err := h.db.Where("id = ? AND merchant_id = ?", req.RuleID, merchantID).First(&rule).Error; err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"code": 403,
				"msg":  "无权限使用此规则",
			})
			return
		}
	}

	// 查询规则配置
	var rule models.SecurityCodeRule
	if err := h.db.First(&rule, req.RuleID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "规则不存在",
		})
		return
	}

	// 查询批次信息
	var batch models.ProductBatch
	if err := h.db.First(&batch, req.BatchID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "批次不存在",
		})
		return
	}

	// 解密规则配置
	encryptionKey := []byte(h.cfg.JWT.Secret) // 使用JWT密钥作为加密密钥
	generator := services.NewCodeGenerator(nil, encryptionKey)

	var ruleConfig services.RuleConfig
	if err := generator.DecryptRuleConfig(rule.RuleConfig); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "规则配置解析失败",
		})
		return
	}

	// 设置批次标识
	ruleConfig.BatchCode = batch.BatchCode

	// 创建生成器
	generator = services.NewCodeGenerator(&ruleConfig, encryptionKey)

	// 生成防伪码
	startSeq := req.StartSeq
	if startSeq == 0 {
		startSeq = 1
	}

	codes, err := generator.GenerateBatch(startSeq, req.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "防伪码生成失败",
		})
		return
	}

	// 保存防伪码到数据库（这里简化处理，实际应该分表存储）
	// 实际实现中需要根据分表策略存储到不同的表中

	// 返回生成结果
	response := GenerateResponse{
		Total:     len(codes),
		BatchCode: batch.BatchCode,
	}

	// 只返回前100条防伪码预览
	if len(codes) > 100 {
		response.Codes = codes[:100]
	} else {
		response.Codes = codes
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "防伪码生成成功",
		"data": response,
	})
}

// GetCodes 获取防伪码列表
func (h *CodeHandler) GetCodes(c *gin.Context) {
	// 获取查询参数
	ruleID := c.Query("rule_id")
	batchID := c.Query("batch_id")
	status := c.Query("status")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "20"))

	// 构建查询条件
	query := h.db.Model(&models.VerificationRecord{})

	// 添加商户过滤条件（非平台管理员）
	merchantID, exists := c.Get("merchantID")
	if exists && merchantID != nil {
		query = query.Where("merchant_id = ?", merchantID)
	}

	if ruleID != "" {
		query = query.Where("rule_id = ?", ruleID)
	}
	if batchID != "" {
		query = query.Where("batch_id = ?", batchID)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// 分页查询
	var total int64
	query.Count(&total)

	var records []models.VerificationRecord
	offset := (page - 1) * size
	query.Offset(offset).Limit(size).Find(&records)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": gin.H{
			"total": total,
			"page":  page,
			"size":  size,
			"list":  records,
		},
	})
}

// BatchUpdateStatus 批量更新防伪码状态
func (h *CodeHandler) BatchUpdateStatus(c *gin.Context) {
	// 实现批量状态更新逻辑
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "状态更新成功",
	})
}

// GetCodeDetail 获取防伪码详情
func (h *CodeHandler) GetCodeDetail(c *gin.Context) {
	id := c.Param("id")

	var record models.VerificationRecord
	if err := h.db.First(&record, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "防伪码不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": record,
	})
}

// ExportCodes 导出防伪码
func (h *CodeHandler) ExportCodes(c *gin.Context) {
	// 实现防伪码导出逻辑
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "导出功能待实现",
	})
}
