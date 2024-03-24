package util

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte(JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

func GetToken(username, password string) (string, error) {

	claims := Claims{
		Username: username,
		Password: password,
		RegisteredClaims: jwt.RegisteredClaims{
			// 24 小时过期
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			// 签发标识
			Issuer: "ai-server",
			// 签发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// 生效时间
			NotBefore: jwt.NewNumericDate(time.Now()),
		}}

	// 使用HS256签名算法
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaim.SignedString(jwtSecret)

	return token, err
}

func ParserToken(tokenStr string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
