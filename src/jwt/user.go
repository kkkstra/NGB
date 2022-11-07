package jwt

import (
	"byitter/src/model"
	jwtgo "github.com/dgrijalva/jwt-go"
)

type UserJwt struct {
	*Token
}

type UserClaims struct {
	*BasicClaims
	Username string         `json:"user"`
	Role     model.RoleType `json:"rol"`
}

func GenerateUserJwtStr(username string, role model.RoleType) TokenInterface {
	claims := UserClaims{
		getBasicClaim(),
		username,
		role,
	}
	tokenClaims := jwtgo.NewWithClaims(jwtgo.SigningMethodRS256, claims)
	return &UserJwt{&Token{tokenClaims, userJwtRSAKey.PrivateKey}}
}
