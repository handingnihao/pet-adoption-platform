package utils

import (
	"errors"
	"pet-adoption-platform/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims JWT Claims结构
type Claims struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT Token
func GenerateToken(userID int64, role string) (string, time.Time, error) {
	expiresAt := time.Now().Add(time.Duration(config.AppConfig.JWT.ExpireHours) * time.Hour)

	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.AppConfig.JWT.Secret))
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expiresAt, nil
}

// ParseToken 解析JWT Token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// RefreshToken 刷新Token（生成新的Token）
func RefreshToken(tokenString string) (string, time.Time, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return "", time.Time{}, err
	}

	// 生成新的Token
	return GenerateToken(claims.UserID, claims.Role)
}
