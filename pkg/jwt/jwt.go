package jwt

import (
	"NGB/internal/config"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte(config.C.User.Jwt.Key)

type Claims struct {
	Role int `json:"role"`
	jwt.StandardClaims
}

func GenerateJWTToken(username string, role int, id string) *jwt.Token {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(config.C.User.Jwt.Expire) * time.Hour)

	claims := Claims{
		role,
		jwt.StandardClaims{
			Issuer:    config.C.User.Jwt.Issuer,
			IssuedAt:  nowTime.Unix(),
			ExpiresAt: expireTime.Unix(),
			Subject:   id,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims
}

func GetJWTTokenString(tokenClaims *jwt.Token) (string, error) {
	token, err := tokenClaims.SignedString(jwtKey)
	return token, err
}

func GetExpiresAt(token *jwt.Token) int64 {
	return token.Claims.(Claims).ExpiresAt
}

func ParseJWTToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, keyFunc)

	if tokenClaims == nil {
		return nil, err
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	}

	return nil, err
}

var keyFunc = func(token *jwt.Token) (interface{}, error) {
	return jwtKey, nil
}
