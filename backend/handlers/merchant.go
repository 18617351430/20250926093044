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

type MerchantHandler struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewMerchantHandler(db *gorm.DB, cfg *config.Config) *MerchantHandler {
	return &MerchantHandler{db: db, cfg: cfg}
}

// GetProducts 获取商户商品列表
func (h *MerchantHandler) GetProducts(c *gin.Context) {
	merchantID, _ := c.Get("merchantID")

	// 获取查询参数
	name := c.Query("name")
	status := c.Query("status")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "20"))

	// 构建查询条件
	query := h.db.Model(&models.Product{}).Where("merchant_id = ?", merchantID)
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// 分页查询
	var total int64
	query.Count(&total)

	var products []models.Product
	offset := (page - 1) * size
	query.Offset(offset).Limit(size).Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": gin.H{
			"total": total,
			"page":  page,
			"size":  size,
			"list":  products,
		},
	})
}

// CreateProduct 创建商品
func (h *MerchantHandler) CreateProduct(c *gin.Context) {
	merchantID, _ := c.Get("merchantID")

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	// 设置商户ID和默认值
	product.MerchantID = merchantID.(uint)
	product.Status = 1
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	if err := h.db.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "商品创建失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "商品创建成功",
		"data": gin.H{
			"product_id": product.ID,
		},
	})
}

// GetBatches 获取商品批次列表
func (h *MerchantHandler) GetBatches(c *gin.Context) {
	merchantID, _ := c.Get("merchantID")

	// 获取查询参数
	productID := c.Query("product_id")
	batchCode := c.Query("batch_code")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "20"))

	// 构建查询条件
	query := h.db.Model(&models.ProductBatch{}).Where("merchant_id = ?", merchantID)
	if productID != "" {
		query = query.Where("product_id = ?", productID)
	}
	if batchCode != "" {
		query = query.Where("batch_code LIKE ?", "%"+batchCode+"%")
	}

	// 分页查询
	var total int64
	query.Count(&total)

	var batches []models.ProductBatch
	offset := (page - 1) * size
	query.Offset(offset).Limit(size).Find(&batches)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": gin.H{
			"total": total,
			"page":  page,
			"size":  size,
			"list":  batches,
		},
	})
}

// CreateBatch 创建批次
func (h *MerchantHandler) CreateBatch(c *gin.Context) {
	merchantID, _ := c.Get("merchantID")

	var batch models.ProductBatch
	if err := c.ShouldBindJSON(&batch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	// 验证商品是否存在且属于当前商户
	var product models.Product
	if err := h.db.Where("id = ? AND merchant_id = ?", batch.ProductID, merchantID).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "商品不存在或无权限",
		})
		return
	}

	// 设置商户ID和默认值
	batch.MerchantID = merchantID.(uint)
	batch.Status = 1
	batch.CreatedAt = time.Now()
	batch.UpdatedAt = time.Now()

	if err := h.db.Create(&batch).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "批次创建失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "批次创建成功",
		"data": gin.H{
			"batch_id": batch.ID,
		},
	})
}

// GetRules 获取防伪码规则列表
func (h *MerchantHandler) GetRules(c *gin.Context) {
	merchantID, _ := c.Get("merchantID")

	// 获取查询参数
	name := c.Query("name")
	status := c.Query("status")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "20"))

	// 构建查询条件
	query := h.db.Model(&models.SecurityCodeRule{}).Where("merchant_id = ?", merchantID)
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// 分页查询
	var total int64
	query.Count(&total)

	var rules []models.SecurityCodeRule
	offset := (page - 1) * size
	query.Offset(offset).Limit(size).Find(&rules)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": gin.H{
			"total": total,
			"page":  page,
			"size":  size,
			"list":  rules,
		},
	})
}

// CreateRule 创建防伪码规则
func (h *MerchantHandler) CreateRule(c *gin.Context) {
	merchantID, _ := c.Get("merchantID")

	var rule models.SecurityCodeRule
	if err := c.ShouldBindJSON(&rule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	// 设置商户ID和默认值
	rule.MerchantID = merchantID.(uint)
	rule.Status = 1
	rule.CreatedAt = time.Now()
	rule.UpdatedAt = time.Now()

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

// GetMerchantStatistics 获取商户统计信息
func (h *MerchantHandler) GetMerchantStatistics(c *gin.Context) {
	merchantID, _ := c.Get("merchantID")

	var stats struct {
		TotalProducts int64 `json:"total_products"`
		TotalBatches  int64 `json:"total_batches"`
		TotalCodes    int64 `json:"total_codes"`
		VerifiedCodes int64 `json:"verified_codes"`
		FakeCodes     int64 `json:"fake_codes"`
		TodayVerified int64 `json:"today_verified"`
		TodayFake     int64 `json:"today_fake"`
	}

	today := time.Now().Format("2006-01-02")

	// 统计商品数量
	h.db.Model(&models.Product{}).Where("merchant_id = ?", merchantID).Count(&stats.TotalProducts)

	// 统计批次数量
	h.db.Model(&models.ProductBatch{}).Where("merchant_id = ?", merchantID).Count(&stats.TotalBatches)

	// 统计防伪码数量
	h.db.Model(&models.VerificationRecord{}).Where("merchant_id = ?", merchantID).Count(&stats.TotalCodes)
	h.db.Model(&models.VerificationRecord{}).Where("merchant_id = ? AND result = 1", merchantID).Count(&stats.VerifiedCodes)
	h.db.Model(&models.VerificationRecord{}).Where("merchant_id = ? AND result = 0", merchantID).Count(&stats.FakeCodes)

	// 统计今日验证数据
	h.db.Model(&models.VerificationRecord{}).
		Where("merchant_id = ? AND DATE(verify_time) = ? AND result = 1", merchantID, today).
		Count(&stats.TodayVerified)
	h.db.Model(&models.VerificationRecord{}).
		Where("merchant_id = ? AND DATE(verify_time) = ? AND result = 0", merchantID, today).
		Count(&stats.TodayFake)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "统计查询成功",
		"data": stats,
	})
}
