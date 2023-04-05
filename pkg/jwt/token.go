package jwt

import (
	"errors"
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
	Str() string
	ParseToken() (*jwtgo.Token, error)
}

type TokenStr string

func (t *Token) GenerateTokenStr() (TokenStr, error) {
	token, err := t.SignedString(t.key)
	return TokenStr(token), err
}

//func GetToken(username string, role model.RoleType) (TokenStr, error) {
//	t := GenerateUserJwt(username, role)
//	str, err := t.GenerateTokenStr()
//	if err != nil {
//		return "", err
//	}
//	return str, nil
//}

func (t *TokenStr) ParseToken(claims jwtgo.Claims) (*jwtgo.Token, error) {
	tokenClaims, err := jwtgo.ParseWithClaims(t.Str(), claims, keyFunc)
	if err != nil {
		return nil, err
	}
	if !tokenClaims.Valid {
		return nil, errors.New("Invalid token. ")
	}
	return tokenClaims, nil
}

func (t *TokenStr) Str() string {
	return string(*t)
}
