package utils

import (
	"anti-fake-system/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// CustomClaims 自定义JWT声明
type CustomClaims struct {
	UserID     uint   `json:"user_id"`
	Username   string `json:"username"`
	UserType   int    `json:"user_type"` // 1-平台用户, 2-商户用户
	MerchantID *uint  `json:"merchant_id,omitempty"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT令牌
func GenerateToken(userID uint, username string, userType int, merchantID *uint, cfg *config.Config) (string, error) {
	claims := CustomClaims{
		UserID:     userID,
		Username:   username,
		UserType:   userType,
		MerchantID: merchantID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(cfg.JWT.Expire))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "anti-fake-system",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT.Secret))
}

// ParseToken 解析JWT令牌
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Load().JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrInvalidKey
}

// RefreshToken 刷新令牌
func RefreshToken(tokenString string, cfg *config.Config) (string, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return "", err
	}

	// 检查令牌是否即将过期（剩余时间小于一半）
	remaining := time.Until(claims.ExpiresAt.Time)
	if remaining > time.Duration(cfg.JWT.Expire/2)*time.Hour {
		return tokenString, nil // 不需要刷新
	}

	// 生成新令牌
	return GenerateToken(claims.UserID, claims.Username, claims.UserType, claims.MerchantID, cfg)
}
