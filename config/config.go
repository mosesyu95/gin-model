package config

import (
	"encoding/json"
	"github.com/toolkits/pkg/file"
	"log"
)

var (
	ConfigFile string
	Config     *GlobalConfig
)

func GetConfigFile() string {
	return ConfigFile
}

func Init(cfg string) {
	if cfg == "" {
		log.Fatalln("config file not specified: use -c $filename")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file specified not found:", cfg)
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file", cfg, "error:", err.Error())
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file", cfg, "error:", err.Error())
	}

	Config = &c
	log.Println("InitConfig ok, file", cfg)

}

type GlobalConfig struct {
	Port  int         `json:"port"`
	Log   LogConfig   `json:"log"`
	DB    DBConfig    `json:"db"`
	Redis RedisConfig `json:"redis"`
}

type LogConfig struct {
	Path  string `json:"path"`
	Level string `json:"level"`
}

type DBConfig struct {
	Debug   bool   `json:"debug"`
	Dsn     string `json:"dsn"`
	MaxIdle int    `json:"maxIdle"`
	MaxConn int    `json:"maxConn"`
}

type RedisConfig struct {
	Enable   bool     `json:"enable"`
	Model    string   `json:"model"`
	Master   string   `json:"master"`
	Addrs    []string `json:"addrs"`
	Password string   `json:"password"`
}
