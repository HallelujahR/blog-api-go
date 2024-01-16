package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecretKey = []byte("memorylinks") // 替换为你的密钥

// Claims 是一个自定义的结构体，用于JWT的payload部分
type Claims struct {
	Phonenumber string `json:"phonenumber"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT Token
func GenerateToken(phonenumber string) (string, error) {
	expirationTime := time.Now().Add(24 * 30 * time.Hour) // Token有效期为30天
	claims := &Claims{
		Phonenumber: phonenumber,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "memorylinks", // 替换为你的应用名
		},
	}

	// 使用HS256算法和密钥生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken 验证JWT Token
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorMalformed)
	}
}
