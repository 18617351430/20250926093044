package handlers

import (
	"anti-fake-system/config"
	"anti-fake-system/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type VerifyHandler struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewVerifyHandler(db *gorm.DB, cfg *config.Config) *VerifyHandler {
	return &VerifyHandler{db: db, cfg: cfg}
}

// VerifyRequest 验证请求
type VerifyRequest struct {
	Code string `json:"code" binding:"required"` // 防伪码
}

// VerifyResponse 验证响应
type VerifyResponse struct {
	IsGenuine       bool       `json:"is_genuine"`                  // 是否真品
	ProductName     string     `json:"product_name"`                // 商品名称
	BatchCode       string     `json:"batch_code"`                  // 批次标识
	ProductionDate  time.Time  `json:"production_date"`             // 生产日期
	FirstVerifyTime *time.Time `json:"first_verify_time,omitempty"` // 首次验证时间
	VerifyCount     int        `json:"verify_count"`                // 验证次数
	MerchantName    string     `json:"merchant_name"`               // 商户名称
	Message         string     `json:"message"`                     // 验证结果消息
}

// VerifyCode 验证防伪码
func (h *VerifyHandler) VerifyCode(c *gin.Context) {
	var req VerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	// 查询防伪码记录（这里简化处理，实际应该根据分表策略查询）
	var record models.VerificationRecord
	if err := h.db.Where("security_code = ?", req.Code).First(&record).Error; err != nil {
		// 防伪码不存在
		response := VerifyResponse{
			IsGenuine: false,
			Message:   "防伪码不存在，请确认输入是否正确",
		}

		// 记录验证失败记录
		h.recordVerification(req.Code, 0, 0, false, c)

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "验证完成",
			"data": response,
		})
		return
	}

	// 查询商品批次信息
	var batch models.ProductBatch
	if err := h.db.First(&batch, record.MerchantID).Error; err != nil {
		response := VerifyResponse{
			IsGenuine: false,
			Message:   "商品信息异常",
		}
		h.recordVerification(req.Code, record.MerchantID, 0, false, c)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "验证完成",
			"data": response,
		})
		return
	}

	// 查询商品信息
	var product models.Product
	if err := h.db.First(&product, batch.ProductID).Error; err != nil {
		response := VerifyResponse{
			IsGenuine: false,
			Message:   "商品信息异常",
		}
		h.recordVerification(req.Code, record.MerchantID, batch.ID, false, c)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "验证完成",
			"data": response,
		})
		return
	}

	// 查询商户信息
	var merchant models.Merchant
	if err := h.db.First(&merchant, record.MerchantID).Error; err != nil {
		response := VerifyResponse{
			IsGenuine: false,
			Message:   "商户信息异常",
		}
		h.recordVerification(req.Code, record.MerchantID, batch.ID, false, c)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "验证完成",
			"data": response,
		})
		return
	}

	// 查询验证历史统计
	var verifyCount int64
	h.db.Model(&models.VerificationRecord{}).
		Where("security_code = ?", req.Code).
		Count(&verifyCount)

	// 查询首次验证时间
	var firstRecord models.VerificationRecord
	h.db.Where("security_code = ?", req.Code).
		Order("verify_time asc").
		First(&firstRecord)

	// 构建响应
	response := VerifyResponse{
		IsGenuine:      record.Result == 1,
		ProductName:    product.Name,
		BatchCode:      batch.BatchCode,
		ProductionDate: batch.ProductionDate,
		VerifyCount:    int(verifyCount),
		MerchantName:   merchant.Name,
	}

	if firstRecord.ID != 0 {
		response.FirstVerifyTime = &firstRecord.VerifyTime
	}

	// 根据验证结果设置消息
	if record.Result == 1 {
		if verifyCount == 1 {
			response.Message = "恭喜！这是正品，首次验证成功"
		} else {
			response.Message = "这是正品，但已被验证过" + strconv.FormatInt(verifyCount-1, 10) + "次"
		}
	} else if record.Result == 0 {
		response.Message = "警告！此防伪码为伪品"
	} else {
		response.Message = "此防伪码已作废"
	}

	// 记录验证记录
	h.recordVerification(req.Code, record.MerchantID, batch.ID, record.Result == 1, c)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "验证完成",
		"data": response,
	})
}

// GetVerifyHistory 获取验证历史
func (h *VerifyHandler) GetVerifyHistory(c *gin.Context) {
	code := c.Param("code")

	var records []models.VerificationRecord
	h.db.Where("security_code = ?", code).
		Order("verify_time desc").
		Limit(10).
		Find(&records)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": records,
	})
}

// recordVerification 记录验证记录
func (h *VerifyHandler) recordVerification(code string, merchantID uint, batchID uint, isGenuine bool, c *gin.Context) {
	result := 0
	if isGenuine {
		result = 1
	}

	record := models.VerificationRecord{
		SecurityCode: code,
		MerchantID:   merchantID,
		VerifyTime:   time.Now(),
		IPAddress:    c.ClientIP(),
		Result:       result,
		UserAgent:    c.GetHeader("User-Agent"),
	}

	h.db.Create(&record)
}

// GetVerifyRecords 获取验证记录
func (h *VerifyHandler) GetVerifyRecords(c *gin.Context) {
	var records []models.VerificationRecord
	h.db.Order("verify_time desc").Limit(100).Find(&records)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": records,
	})
}

// GetVerifyStatistics 获取验证统计
func (h *VerifyHandler) GetVerifyStatistics(c *gin.Context) {
	var totalCount int64
	var genuineCount int64

	h.db.Model(&models.VerificationRecord{}).Count(&totalCount)
	h.db.Model(&models.VerificationRecord{}).Where("result = 1").Count(&genuineCount)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": gin.H{
			"total_count":   totalCount,
			"genuine_count": genuineCount,
			"fake_count":    totalCount - genuineCount,
		},
	})
}
