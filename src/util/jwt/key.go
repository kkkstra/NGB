package jwt

import (
	"byitter/src/config"
	"byitter/src/model"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	jwtgo "github.com/dgrijalva/jwt-go"
	"os"
)

type RSAKey struct {
	Kid        string
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
}

var (
	rsaKeyList    = []string{"user_jwt"}
	rsaKeys       = map[string]RSAKey{}
	userJwtRSAKey *RSAKey
	//Jwks          []JWK
)

func InitRSAKey() {
	if config.C.Init.RsaKey {
		for _, keyType := range rsaKeyList {
			err := GenerateRSAKey(keyType)
			if err != nil {
				panic(err)
				return
			}
			m := model.GetModel(&model.RSAKeyModel{})
			m.(*model.RSAKeyModel).CreateRSAKey(keyType)
		}
	}
	loadRSAkeys()
	//loadJWKs()
	//fmt.Println(userJwtRSAKey)
}

func loadRSAkeys() {
	m := model.GetModel(&model.RSAKeyModel{})
	keyList, err := m.(*model.RSAKeyModel).FindRSAKey()
	if err != nil {
		panic(err)
		return
	}
	for _, keyPem := range keyList {
		publicKey, err := jwtgo.ParseRSAPublicKeyFromPEM([]byte(keyPem.PublicKey))
		if err != nil {
			panic(err)
			return
		}
		privateKey, err := jwtgo.ParseRSAPrivateKeyFromPEM([]byte(keyPem.PrivateKey))
		if err != nil {
			panic(err)
			return
		}
		keyRsa := RSAKey{
			Kid:        keyPem.Kid,
			PublicKey:  publicKey,
			PrivateKey: privateKey,
		}
		rsaKeys[keyPem.Kid] = keyRsa
		switch keyPem.Type {
		case "user_jwt":
			userJwtRSAKey = &keyRsa
		}
	}
}

func GenerateRSAKey(keyType string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("./env/RSAKey/" + keyType + "-" + config.PrivateKeyFile)
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	file.Close()

	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("./env/RSAKey/" + keyType + "-" + config.PublicKeyFile)
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}

var keyFunc = func(t *jwtgo.Token) (interface{}, error) {
	switch t.Method.Alg() {
	case jwtgo.SigningMethodRS256.Alg():
		kid, ok := t.Header["kid"].(string)
		if !ok {
			return nil, errors.New("Failed to find kid")
		}
		key, ok := rsaKeys[kid]
		if !ok {
			return nil, errors.New("Failed to find the key. ")
		}
		return key.PublicKey, nil
	}
	return nil, errors.New("Failed to find the key. ")
}
