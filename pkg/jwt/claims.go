package jwt

import (
	"NGB/internal/config"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
)

type BasicClaims struct {
	jwtgo.StandardClaims
}

type ClaimsInterface interface {
}

func getBasicClaim() *BasicClaims {
	return &BasicClaims{
		jwtgo.StandardClaims{
			Issuer:    config.C.User.UserJWT.Issuer,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Duration(config.C.User.UserJWT.Expire) * time.Hour).Unix(),
		},
	}
}
