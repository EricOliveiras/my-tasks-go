package config

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Config struct {
	Auth AuthConfig
	DB   DBConfig
	HTTP HTTPConfig
	Cors gin.HandlerFunc
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		Auth: LoadAuthConfig(),
		DB:   LoadDBConfig(),
		HTTP: LoadHTTPConfig(),
		Cors: LoadCorsConfig(),
	}
}
