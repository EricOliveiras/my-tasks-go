package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
	})
}
