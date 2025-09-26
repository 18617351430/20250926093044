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

type PlatformHandler struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewPlatformHandler(db *gorm.DB, cfg *config.Config) *PlatformHandler {
	return &PlatformHandler{db: db, cfg: cfg}
}

// GetMerchants 获取商户列表
func (h *PlatformHandler) GetMerchants(c *gin.Context) {
	// 获取查询参数
	name := c.Query("name")
	status := c.Query("status")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "20"))

	// 构建查询条件
	query := h.db.Model(&models.Merchant{})
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// 分页查询
	var total int64
	query.Count(&total)

	var merchants []models.Merchant
	offset := (page - 1) * size
	query.Offset(offset).Limit(size).Find(&merchants)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": gin.H{
			"total": total,
			"page":  page,
			"size":  size,
			"list":  merchants,
		},
	})
}

// CreateMerchant 创建商户
func (h *PlatformHandler) CreateMerchant(c *gin.Context) {
	var merchant models.Merchant
	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	// 设置默认值
	merchant.Status = 1
	merchant.CreatedAt = time.Now()
	merchant.UpdatedAt = time.Now()

	if err := h.db.Create(&merchant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "商户创建失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "商户创建成功",
		"data": gin.H{
			"merchant_id": merchant.ID,
		},
	})
}

// UpdateMerchant 更新商户
func (h *PlatformHandler) UpdateMerchant(c *gin.Context) {
	id := c.Param("id")

	var merchant models.Merchant
	if err := h.db.First(&merchant, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "商户不存在",
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

	updateData["updated_at"] = time.Now()

	if err := h.db.Model(&merchant).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "商户更新失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "商户更新成功",
	})
}

// GetStatistics 获取平台统计信息
func (h *PlatformHandler) GetStatistics(c *gin.Context) {
	var stats struct {
		TotalMerchants  int64 `json:"total_merchants"`
		ActiveMerchants int64 `json:"active_merchants"`
		TotalProducts   int64 `json:"total_products"`
		TotalBatches    int64 `json:"total_batches"`
		TotalCodes      int64 `json:"total_codes"`
		VerifiedCodes   int64 `json:"verified_codes"`
		FakeCodes       int64 `json:"fake_codes"`
	}

	// 统计商户数量
	h.db.Model(&models.Merchant{}).Count(&stats.TotalMerchants)
	h.db.Model(&models.Merchant{}).Where("status = 1").Count(&stats.ActiveMerchants)

	// 统计商品数量
	h.db.Model(&models.Product{}).Count(&stats.TotalProducts)

	// 统计批次数量
	h.db.Model(&models.ProductBatch{}).Count(&stats.TotalBatches)

	// 统计防伪码数量（简化处理）
	h.db.Model(&models.VerificationRecord{}).Count(&stats.TotalCodes)
	h.db.Model(&models.VerificationRecord{}).Where("result = 1").Count(&stats.VerifiedCodes)
	h.db.Model(&models.VerificationRecord{}).Where("result = 0").Count(&stats.FakeCodes)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "统计查询成功",
		"data": stats,
	})
}

// GetVerifyTrend 获取验证趋势数据
func (h *PlatformHandler) GetVerifyTrend(c *gin.Context) {
	days, _ := strconv.Atoi(c.DefaultQuery("days", "7"))

	// 获取最近N天的验证趋势数据
	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -days)

	var trends []struct {
		Date    string `json:"date"`
		Total   int64  `json:"total"`
		Genuine int64  `json:"genuine"`
		Fake    int64  `json:"fake"`
	}

	// 按日期统计验证记录
	rows, err := h.db.Model(&models.VerificationRecord{}).
		Select("DATE(verify_time) as date, COUNT(*) as total, "+
			"SUM(CASE WHEN result = 1 THEN 1 ELSE 0 END) as genuine, "+
			"SUM(CASE WHEN result = 0 THEN 1 ELSE 0 END) as fake").
		Where("verify_time BETWEEN ? AND ?", startDate, endDate).
		Group("DATE(verify_time)").
		Order("date").
		Rows()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "趋势数据查询失败",
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var trend struct {
			Date    string `json:"date"`
			Total   int64  `json:"total"`
			Genuine int64  `json:"genuine"`
			Fake    int64  `json:"fake"`
		}
		rows.Scan(&trend.Date, &trend.Total, &trend.Genuine, &trend.Fake)
		trends = append(trends, trend)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "趋势数据查询成功",
		"data": trends,
	})
}

// GetMerchantDetail 获取商户详情
func (h *PlatformHandler) GetMerchantDetail(c *gin.Context) {
	id := c.Param("id")

	var merchant models.Merchant
	if err := h.db.First(&merchant, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "商户不存在",
		})
		return
	}

	// 获取商户统计信息
	var stats struct {
		TotalProducts int64 `json:"total_products"`
		TotalBatches  int64 `json:"total_batches"`
		TotalCodes    int64 `json:"total_codes"`
		VerifiedCodes int64 `json:"verified_codes"`
	}

	h.db.Model(&models.Product{}).Where("merchant_id = ?", id).Count(&stats.TotalProducts)
	h.db.Model(&models.ProductBatch{}).Where("merchant_id = ?", id).Count(&stats.TotalBatches)
	h.db.Model(&models.VerificationRecord{}).Where("merchant_id = ?", id).Count(&stats.TotalCodes)
	h.db.Model(&models.VerificationRecord{}).Where("merchant_id = ? AND result = 1", id).Count(&stats.VerifiedCodes)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "商户详情查询成功",
		"data": gin.H{
			"merchant": merchant,
			"stats":    stats,
		},
	})
}
