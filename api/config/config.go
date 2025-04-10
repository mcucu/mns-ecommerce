package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type IConfig interface {
	Get(key string) any
	GetBool(key string) bool
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetString(key string) string
}

type Config struct{}

func NewConfig() IConfig {
	loadDotEnv()
	return &Config{}
}

func loadDotEnv() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		os.Setenv("APP_ENV", "development")
	}

	if appEnv != "heroku" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func (c *Config) Get(key string) interface{} {
	loadDotEnv()
	return os.Getenv(key)
}

func (c *Config) GetBool(key string) bool {
	return os.Getenv(key) == "true"
}

func (c *Config) GetFloat64(key string) float64 {
	val := os.Getenv(key)
	if val == "" {
		return 0
	}
	result, _ := strconv.ParseFloat(val, 64)
	return result
}

func (c *Config) GetInt(key string) int {
	val := os.Getenv(key)
	if val == "" {
		return 0
	}
	result, _ := strconv.Atoi(val)
	return result
}

func (c *Config) GetInt64(key string) int64 {
	val := os.Getenv(key)
	if val == "" {
		return 0
	}
	result, _ := strconv.ParseInt(val, 10, 64)
	return result
}

func (c *Config) GetString(key string) string {
	val := os.Getenv(key)
	if val == "" {
		return ""
	}
	return val
}
