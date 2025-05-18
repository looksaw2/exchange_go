package config

import (
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	InfoConfig     InfoConfig     `mapstructure:"info"`
	ExchangeConfig ExchangeConfig `mapstructure:"exchange_go"`
}

type InfoConfig struct {
	Version     string `mapstructure:"version"`
	Author      string `mapstructure:"author"`
	Description string `mapstructure:"description"`
}

type ExchangeConfig struct {
	Port     int            `mapstructure:"port"`
	DbConfig DataBaseConfig `mapstructure:"database"`
}

type DataBaseConfig struct {
	Driver string `mapstructure:"driver"`
	Uri    string `mapstructure:"uri"`
}

var (
	cfg     *Config
	cfgOnce sync.Once
)

func InitConfig() *Config {
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../config")
	_, b, _, _ := runtime.Caller(0)
	rootPath := filepath.Dir(filepath.Dir(b))
	viper.AddConfigPath(filepath.Join(rootPath, "config"))
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic("Can't read config.yml")
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic("Can't parse config.yml")
	}
	cfg.ExchangeConfig.DbConfig.Uri = getDBUri()
	return &cfg
}

func getDBUri() string {
	dbUri := os.Getenv("EXCHANGEGO_DB_DSN")
	return dbUri
}
