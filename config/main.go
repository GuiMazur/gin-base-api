package config

import (
	"gin-base-api/utils"
)

type App struct {
	Host string
	Port string
}

type Config struct {
	App App
}

func NewConfig() *Config {
	return &Config{
		App{
			Host: utils.GetEnvOrDefault("APP_HOST", "localhosto"),
			Port: utils.GetEnvOrDefault("APP_PORT", "8086"),
		},
	}
}