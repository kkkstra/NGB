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

func GenerateUserJwt(username string, role model.RoleType) *UserJwt {
	claims := UserClaims{
		getBasicClaim(),
		username,
		role,
	}
	tokenClaims := jwtgo.NewWithClaims(jwtgo.SigningMethodRS256, claims)
	tokenClaims.Header["kid"] = userJwtRSAKey.Kid
	return &UserJwt{&Token{tokenClaims, userJwtRSAKey.PrivateKey}}
}

func (t *UserJwt) ExpiresAt() (exp int64) {
	return t.Claims.(UserClaims).ExpiresAt
}
