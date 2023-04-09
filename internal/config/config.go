package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App      app      `yaml:"app"`
	Database database `yaml:"database"`
	Debug    debug    `yaml:"debug"`
	User     userConf `yaml:"user"`
}

type app struct {
	Addr string `yaml:"addr"`
}

type database struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Port     string `yaml:"port"`
}
type debug struct {
	Enable bool `yaml:"enable"`
}

type userConf struct {
	Jwt jwt `yaml:"jwt"`
}

type jwt struct {
	Expire int    `yaml:"expire"`
	Issuer string `yaml:"issuer"`
	Key    string `yaml:"key"`
}

var C *Config

func init() {
	configFile := "default.yaml"
	r, err := os.ReadFile(fmt.Sprintf("./configs/config_files/%s", configFile))
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
