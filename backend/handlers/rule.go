package handlers

import (
	"anti-fake-system/config"
	"anti-fake-system/models"
	"anti-fake-system/services"
	"anti-fake-system/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RuleHandler struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewRuleHandler(db *gorm.DB, cfg *config.Config) *RuleHandler {
	return &RuleHandler{db: db, cfg: cfg}
}

// CreateRuleRequest 创建规则请求
type CreateRuleRequest struct {
	Name        string              `json:"name" binding:"required"`
	Description string              `json:"description"`
	RuleConfig  services.RuleConfig `json:"rule_config" binding:"required"`
	Status      int                 `json:"status"`
}

// CreateRule 创建防伪码规则
func (h *RuleHandler) CreateRule(c *gin.Context) {
	merchantID, _ := c.Get("merchantID")

	var req CreateRuleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	// 加密规则配置
	encryptionKey := []byte(h.cfg.JWT.Secret)
	generator := services.NewCodeGenerator(&req.RuleConfig, encryptionKey)

	encryptedConfig, err := generator.EncryptRuleConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "规则配置加密失败",
		})
		return
	}

	// 创建规则记录
	rule := models.SecurityCodeRule{
		Name:       req.Name,
		RuleConfig: encryptedConfig,
		MerchantID: merchantID.(uint),
		Status:     req.Status,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := h.db.Create(&rule).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "规则创建失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "规则创建成功",
		"data": gin.H{
			"rule_id": rule.ID,
		},
	})
}

// UpdateRule 更新防伪码规则
func (h *RuleHandler) UpdateRule(c *gin.Context) {
	merchantID, _ := c.Get("merchantID")
	ruleID := c.Param("id")

	// 检查规则是否存在且属于当前商户
	var rule models.SecurityCodeRule
	if err := h.db.Where("id = ? AND merchant_id = ?", ruleID, merchantID).First(&rule).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "规则不存在或无权限",
		})
		return
	}

	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	// 如果更新了规则配置，需要重新加密
	if ruleConfig, exists := updateData["rule_config"]; exists {
		encryptionKey := []byte(h.cfg.JWT.Secret)

		// 将规则配置转换为JSON
		configJSON, err := json.Marshal(ruleConfig)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "规则配置序列化失败",
			})
			return
		}

		// 加密配置
		encryptedConfig, err := utils.Encrypt(configJSON, encryptionKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "规则配置加密失败",
			})
			return
		}

		updateData["rule_config"] = encryptedConfig
	}

	updateData["updated_at"] = time.Now()

	if err := h.db.Model(&rule).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "规则更新失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "规则更新成功",
	})
}

// GetRuleDetail 获取规则详情
func (h *RuleHandler) GetRuleDetail(c *gin.Context) {
	merchantID, _ := c.Get("merchantID")
	ruleID := c.Param("id")

	var rule models.SecurityCodeRule
	if err := h.db.Where("id = ? AND merchant_id = ?", ruleID, merchantID).First(&rule).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "规则不存在或无权限",
		})
		return
	}

	// 解密规则配置
	encryptionKey := []byte(h.cfg.JWT.Secret)
	generator := services.NewCodeGenerator(nil, encryptionKey)

	var ruleConfig services.RuleConfig
	if err := generator.DecryptRuleConfig(rule.RuleConfig); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "规则配置解密失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "规则详情查询成功",
		"data": gin.H{
			"rule":        rule,
			"rule_config": ruleConfig,
		},
	})
}

// TestRule 测试规则生成
func (h *RuleHandler) TestRule(c *gin.Context) {
	merchantID, _ := c.Get("merchantID")
	ruleID := c.Param("id")

	// 获取测试参数
	testCount, _ := strconv.Atoi(c.DefaultQuery("count", "5"))
	startSeq, _ := strconv.Atoi(c.DefaultQuery("start_seq", "1"))

	// 查询规则
	var rule models.SecurityCodeRule
	if err := h.db.Where("id = ? AND merchant_id = ?", ruleID, merchantID).First(&rule).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "规则不存在或无权限",
		})
		return
	}

	// 解密规则配置
	encryptionKey := []byte(h.cfg.JWT.Secret)
	generator := services.NewCodeGenerator(nil, encryptionKey)

	var ruleConfig services.RuleConfig
	if err := generator.DecryptRuleConfig(rule.RuleConfig); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "规则配置解密失败",
		})
		return
	}

	// 创建生成器并测试生成
	generator = services.NewCodeGenerator(&ruleConfig, encryptionKey)
	testCodes, err := generator.GenerateBatch(startSeq, testCount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "测试生成失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "测试生成成功",
		"data": gin.H{
			"test_codes":  testCodes,
			"rule_config": ruleConfig,
		},
	})
}

// ValidateRuleConfig 验证规则配置
func (h *RuleHandler) ValidateRuleConfig(c *gin.Context) {
	var ruleConfig services.RuleConfig
	if err := c.ShouldBindJSON(&ruleConfig); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "规则配置参数错误",
		})
		return
	}

	// 验证规则配置的完整性
	if ruleConfig.MerchantCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "商户标识码不能为空",
		})
		return
	}

	if ruleConfig.BatchCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "批次标识不能为空",
		})
		return
	}

	if ruleConfig.Sequence == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "序号配置不能为空",
		})
		return
	}

	if ruleConfig.TotalLength <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "总长度必须大于0",
		})
		return
	}

	// 测试生成一个防伪码验证配置是否有效
	encryptionKey := []byte(h.cfg.JWT.Secret)
	generator := services.NewCodeGenerator(&ruleConfig, encryptionKey)

	testCode, err := generator.GenerateSingle(1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "规则配置无效: " + err.Error(),
		})
		return
	}

	// 验证生成的防伪码格式
	if !generator.ValidateCode(testCode) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "生成的防伪码格式不符合规则",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "规则配置验证通过",
		"data": gin.H{
			"test_code":   testCode,
			"code_length": len(testCode),
		},
	})
}
