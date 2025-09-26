package middleware

import (
	"anti-fake-system/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization头
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "未提供认证令牌",
			})
			c.Abort()
			return
		}

		// 检查Bearer格式
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "认证令牌格式错误",
			})
			c.Abort()
			return
		}

		// 验证JWT令牌
		token := parts[1]
		claims, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "认证令牌无效或已过期",
			})
			c.Abort()
			return
		}

		// 将用户信息存入上下文
		c.Set("userID", claims.UserID)
		c.Set("userType", claims.UserType)
		c.Set("merchantID", claims.MerchantID)
		c.Set("username", claims.Username)

		c.Next()
	}
}

// PlatformAdminOnly 平台管理员专用中间件
func PlatformAdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		userType, exists := c.Get("userType")
		if !exists || userType != 1 { // 1表示平台用户
			c.JSON(http.StatusForbidden, gin.H{
				"code": 403,
				"msg":  "无权限访问此资源",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// MerchantAdminOnly 商户管理员专用中间件
func MerchantAdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		userType, exists := c.Get("userType")
		merchantID, _ := c.Get("merchantID")

		if !exists || userType != 2 || merchantID == nil { // 2表示商户用户
			c.JSON(http.StatusForbidden, gin.H{
				"code": 403,
				"msg":  "无权限访问此资源",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// PlatformAuth 平台认证中间件（组合认证和平台管理员权限检查）
func PlatformAuth() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		// 先进行认证
		AuthMiddleware()(c)
		if c.IsAborted() {
			return
		}

		// 再检查平台管理员权限
		PlatformAdminOnly()(c)
	})
}

// MerchantAuth 商户认证中间件（组合认证和商户权限检查）
func MerchantAuth(jwtSecret string) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		// 先进行认证
		AuthMiddleware()(c)
		if c.IsAborted() {
			return
		}

		// 再检查商户权限
		MerchantAdminOnly()(c)
	})
}
