package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App      app      `yaml:"app"`
	Database database `yaml:"database"`
	Log      log      `yaml:"log"`
	User     userConf `yaml:"user"`
}

type app struct {
	Addr  string `yaml:"addr"`
	Debug bool   `yaml:"debug"`
}

type database struct {
	Sql		sql		`yaml:"sql"`
	Redis	redis	`yaml:"redis"`
}

type sql struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Port     string `yaml:"port"`
}

type redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Db       int		`yaml:"db"`
}

type log struct {
	Filepath       string `yaml:"filepath"`
	FilenamePrefix string `yaml:"filename-prefix"`
}

type userConf struct {
	Jwt jwt `yaml:"jwt"`
}

type jwt struct {
	Expire int         `yaml:"expire"`
	Issuer string      `yaml:"issuer"`
	Key    string      `yaml:"key"`
	SkipPaths [][]string `yaml:"skip-paths"`
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
