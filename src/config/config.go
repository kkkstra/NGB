package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	App        app        `yaml:"app"`
	Postgresql postgresql `yaml:"postgresql"`
	Init       initEnv    `yaml:"init-env"`
	Debug      debug      `yaml:"debug"`
	User       userConf   `yaml:"user"`
}

type app struct {
	Addr string `yaml:"addr"`
}

type postgresql struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Port     string `yaml:"port"`
}

type initEnv struct {
	Admin  bool `yaml:"admin"`
	RsaKey bool `yaml:"rsa-key"`
}

type debug struct {
	Enable bool `yaml:"enable"`
}

type userConf struct {
	UserJWT userJWT `yaml:"user-jwt"`
	Admin   admin   `yaml:"admin"`
}

type userJWT struct {
	Expire int    `yaml:"expire"`
	Issuer string `yaml:"issuer"`
}

type admin struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Email    string `yaml:"email"`
}

const (
	PublicKeyFile  = "rsa-public-key.pem"
	PrivateKeyFile = "rsa-private-key.pem"
)

var C *Config

func init() {
	configFile := "config.yaml"
	r, err := os.ReadFile(fmt.Sprintf("./env/config/%s", configFile))
	if err != nil {
		panic(err)
		return
	}
	config := &Config{}
	err = yaml.Unmarshal(r, config)
	if err != nil {
		panic(err)
		return
	}
	C = config
}

func ReadRSAKeyFromFile(keyType string) (publicKey, privateKey string) {
	publicKeyByte, err := os.ReadFile(fmt.Sprintf("./env/RSAKey/%s", keyType+"-"+PublicKeyFile))
	if err != nil {
		panic(err)
		return
	}
	privateKeyByte, err := os.ReadFile(fmt.Sprintf("./env/RSAKey/%s", keyType+"-"+PrivateKeyFile))
	if err != nil {
		panic(err)
		return
	}
	return string(publicKeyByte), string(privateKeyByte)
}
