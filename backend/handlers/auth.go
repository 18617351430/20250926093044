package handlers

import (
	"anti-fake-system/config"
	"anti-fake-system/models"
	"anti-fake-system/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewAuthHandler(db *gorm.DB, cfg *config.Config) *AuthHandler {
	return &AuthHandler{db: db, cfg: cfg}
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token      string    `json:"token"`
	UserID     uint      `json:"user_id"`
	Username   string    `json:"username"`
	UserType   int       `json:"user_type"`
	MerchantID *uint     `json:"merchant_id,omitempty"`
	ExpireAt   time.Time `json:"expire_at"`
}

// Login 用户登录
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	// 查询用户
	var user models.User
	if err := h.db.Where("username = ? AND status = 1", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "用户名或密码错误",
		})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "用户名或密码错误",
		})
		return
	}

	// 生成JWT令牌
	token, err := utils.GenerateToken(user.ID, user.Username, user.UserType, user.MerchantID, h.cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "令牌生成失败",
		})
		return
	}

	// 更新最后登录时间
	now := time.Now()
	h.db.Model(&user).Update("last_login_at", now)

	// 返回登录结果
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": LoginResponse{
			Token:      token,
			UserID:     user.ID,
			Username:   user.Username,
			UserType:   user.UserType,
			MerchantID: user.MerchantID,
			ExpireAt:   now.Add(time.Hour * time.Duration(h.cfg.JWT.Expire)),
		},
	})
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Phone      string `json:"phone"`
	UserType   int    `json:"user_type" binding:"required"` // 1-平台用户, 2-商户用户
	MerchantID *uint  `json:"merchant_id"`                  // 商户用户必填
}

// Register 用户注册
func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	// 检查用户名是否已存在
	var existingUser models.User
	if err := h.db.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"code": 409,
			"msg":  "用户名已存在",
		})
		return
	}

	// 商户用户必须指定商户ID
	if req.UserType == 2 && req.MerchantID == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "商户用户必须指定商户ID",
		})
		return
	}

	// 检查商户是否存在（如果是商户用户）
	if req.UserType == 2 && req.MerchantID != nil {
		var merchant models.Merchant
		if err := h.db.First(&merchant, *req.MerchantID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "商户不存在",
			})
			return
		}
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "密码加密失败",
		})
		return
	}

	// 创建用户
	user := models.User{
		Username:   req.Username,
		Password:   string(hashedPassword),
		Email:      req.Email,
		Phone:      req.Phone,
		UserType:   req.UserType,
		MerchantID: req.MerchantID,
		Status:     1,
	}

	if err := h.db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "用户创建失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
		"data": gin.H{
			"user_id": user.ID,
		},
	})
}

// PlatformLogin 平台用户登录
func (h *AuthHandler) PlatformLogin(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	// 查询平台用户
	var user models.User
	if err := h.db.Where("username = ? AND user_type = 1 AND status = 1", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "用户名或密码错误",
		})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "用户名或密码错误",
		})
		return
	}

	// 生成JWT令牌
	token, err := utils.GenerateToken(user.ID, user.Username, user.UserType, user.MerchantID, h.cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "令牌生成失败",
		})
		return
	}

	// 更新最后登录时间
	now := time.Now()
	h.db.Model(&user).Update("last_login_at", now)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": LoginResponse{
			Token:      token,
			UserID:     user.ID,
			Username:   user.Username,
			UserType:   user.UserType,
			MerchantID: user.MerchantID,
			ExpireAt:   now.Add(time.Hour * time.Duration(h.cfg.JWT.Expire)),
		},
	})
}

// MerchantLogin 商户用户登录
func (h *AuthHandler) MerchantLogin(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	// 查询商户用户
	var user models.User
	if err := h.db.Where("username = ? AND user_type = 2 AND status = 1", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "用户名或密码错误",
		})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "用户名或密码错误",
		})
		return
	}

	// 生成JWT令牌
	token, err := utils.GenerateToken(user.ID, user.Username, user.UserType, user.MerchantID, h.cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "令牌生成失败",
		})
		return
	}

	// 更新最后登录时间
	now := time.Now()
	h.db.Model(&user).Update("last_login_at", now)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": LoginResponse{
			Token:      token,
			UserID:     user.ID,
			Username:   user.Username,
			UserType:   user.UserType,
			MerchantID: user.MerchantID,
			ExpireAt:   now.Add(time.Hour * time.Duration(h.cfg.JWT.Expire)),
		},
	})
}

// RefreshToken 刷新令牌
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "未认证",
		})
		return
	}

	// 查询用户信息
	var user models.User
	if err := h.db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "用户不存在",
		})
		return
	}

	// 生成新的JWT令牌
	token, err := utils.GenerateToken(user.ID, user.Username, user.UserType, user.MerchantID, h.cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "令牌生成失败",
		})
		return
	}

	now := time.Now()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "令牌刷新成功",
		"data": LoginResponse{
			Token:      token,
			UserID:     user.ID,
			Username:   user.Username,
			UserType:   user.UserType,
			MerchantID: user.MerchantID,
			ExpireAt:   now.Add(time.Hour * time.Duration(h.cfg.JWT.Expire)),
		},
	})
}
