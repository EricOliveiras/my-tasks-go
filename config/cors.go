package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"https://my-tasks-project.vercel.app"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
	})
}
