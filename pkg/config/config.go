package config

import (
	"fmt"
	"os"
	"strconv"
)

type App struct {
	Host string
	Port string
}

type DB struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

type Jwt struct {
	Secret         string
	ExpirationTime int // in seconds
}

type Config struct {
	App
	DB
	Jwt
}

var configInstance *Config

func New() *Config {
	if configInstance == nil {
		configInstance = &Config{
			App{
				Host: getEnvOrDefault("APP_HOST", "localhost"),
				Port: getEnvOrDefault("APP_PORT", "8080"),
			},
			DB{
				Host:     getEnvOrDefault("DB_HOST", "localhost"),
				Port:     getEnvOrDefault("DB_PORT", "3306"),
				Name:     getEnvOrDefault("DB_NAME", "gin_base_api"),
				User:     getEnvOrDefault("DB_USER", "root"),
				Password: getEnvOrDefault("DB_PASSWORD", "root"),
			},
			Jwt{
				Secret:         getEnvOrPanic("JWT_SECRET"),
				ExpirationTime: expirationTimeToInt(getEnvOrDefault("JWT_EXPIRATION_TIME", "86400")), // in seconds (default: 24 hours)
			},
		}
	}
	return configInstance
}

func expirationTimeToInt(str string) int {
	intValue, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error converting string to int on the JWT_EXPIRATION_TIME env variable: ", err)
		return 86400
	}
	return intValue
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("Environment variable " + key + " is not set")
	}
	return value
}
