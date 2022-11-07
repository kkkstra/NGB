package jwt

import (
	"byitter/src/config"
	jwtgo "github.com/dgrijalva/jwt-go"
	"time"
)

type BasicClaims struct {
	jwtgo.StandardClaims
}

func getBasicClaim() *BasicClaims {
	return &BasicClaims{
		jwtgo.StandardClaims{
			Issuer:    "",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Duration(config.C.UserJWT.Expire) * time.Hour).Unix(),
		},
	}
}
