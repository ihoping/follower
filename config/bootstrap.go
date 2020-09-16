package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type DatabaseConfig struct {
	Driver    string `yaml:"driver"`
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	Database  string `yaml:"database"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Charset   string `yaml:"charset"`
	Collation string `yaml:"collation"`
	Prefix    string `yaml:"prefix"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}

type AppConfig struct {
	AppName    string `yaml:"app_name"`
	AppVersion string `yaml:"app_version"`
	AppAddr    string `yaml:"app_addr"`
	Env        string `yaml:"env"`
	Database   []DatabaseConfig
	Redis      []RedisConfig
	LogHost    string `yaml:"log_host"`
	LogPort    string `yaml:"log_port"`
}

var App AppConfig

func init() {
	yamlContent, err := ioutil.ReadFile("./config/app.yaml")
	if err != nil {
		log.Fatalln("read app.yaml err")
	}

	err = yaml.Unmarshal(yamlContent, &App)
	if err != nil {
		log.Fatalln("yaml.Unmarshal err")
	}
}
