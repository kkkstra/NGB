package jwt

import (
	"byitter/src/config"
	"byitter/src/model"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
	"os"
)

type RSAKey struct {
	Kid        string
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
}

type JWK struct {
	Kid string
	E   string
	Kty string
	N   string
}

var (
	rsaKeys       = map[string]RSAKey{}
	userJwtRSAKey *RSAKey
	Jwks          []JWK
)

func InitRSAKey() {
	if config.C.Init.RsaKey {
		err := GenerateRSAKey()
		if err != nil {
			panic(err)
			return
		}
		m := model.GetModel()
		m.CreateRSAKey("user_jwt")
	}
	loadRSAkeys()
	loadJWKs()
}

func loadRSAkeys() {
	m := model.GetModel()
	keyList := m.FindRSAKey()
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

func loadJWKs() {
	for _, rsaKey := range rsaKeys {
		pubKey, err := jwk.New(rsaKey.PublicKey)
		if err != nil {
			panic(err)
			return
		}
		if _, ok := pubKey.(jwk.RSAPublicKey); !ok {
			panic(err)
			return
		}
		err = pubKey.Set(jwk.KeyIDKey, rsaKey.Kid)
		if err != nil {
			panic(err)
			return
		}
		buf, err := json.MarshalIndent(pubKey, "", "  ")
		if err != nil {
			panic(err)
			return
		}

		jwkey := &JWK{}
		err = json.Unmarshal(buf, jwkey)
		if err != nil {
			panic(err)
			return
		}
		Jwks = append(Jwks, *jwkey)
	}
}

func GenerateRSAKey() error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("./env/RSAKey/" + config.PrivateKeyFile)
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
		Type:  "RSA PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("./env/RSAKey/" + config.PublicKeyFile)
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
