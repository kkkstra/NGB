package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	App        app        `yaml:"app"`
	Postgresql postgresql `yaml:"postgresql"`
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
