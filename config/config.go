package config

import (
	"gin-base-api/utils"
)

type App struct {
	Host string
	Port string
}

type DB struct {
	Host string
	Port string
	Name string
	User string
	Password string
}


type Config struct {
	App App
	DB DB
}

func NewConfig() *Config {
	return &Config{
		App{
			Host: utils.GetEnvOrDefault("APP_HOST", "localhost"),
			Port: utils.GetEnvOrDefault("APP_PORT", "8080"),
		},
		DB{
			Host: utils.GetEnvOrDefault("DB_HOST", "localhost"),
			Port: utils.GetEnvOrDefault("DB_PORT", "3306"),
			Name: utils.GetEnvOrDefault("DB_NAME", "gin_base_api"),
			User: utils.GetEnvOrDefault("DB_USER", "root"),
			Password: utils.GetEnvOrDefault("DB_PASSWORD", "root"),
		},
	}
}