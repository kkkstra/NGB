package jwt

import (
	"byitter/src/model"
	jwtgo "github.com/dgrijalva/jwt-go"
)

type Token struct {
	*jwtgo.Token
	key interface{}
}

type TokenInterface interface {
	GenerateTokenStr() (TokenStr, error)
}

type TokenStrInterface interface {
	ParseToken() (*jwtgo.Token, error)
}

type TokenStr string

func (t *Token) GenerateTokenStr() (TokenStr, error) {
	token, err := t.SignedString(t.key)
	return TokenStr(token), err
}

func GetToken(username string, role model.RoleType) (TokenStr, error) {
	t := GenerateUserJwt(username, role)
	str, err := t.GenerateTokenStr()
	if err != nil {
		return "", err
	}
	return str, nil
}

//func (t *TokenStr) ParseToken() (*jwtgo.Token, error) {
//tokenClaims, err := jwtgo.ParseWithClaims(token, &UserClaims{}, func(token *jwtgo.Token) (interface{}, error) {
//	return jwtSecret, nil
//})
//
//if tokenClaims != nil {
//	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
//		return claims, nil
//	}
//}
//
//return nil, err
//}
