package jwt

import (
	"byitter/src/config"
	jwtgo "github.com/dgrijalva/jwt-go"
	"time"
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
