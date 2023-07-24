package routes

import (
	"github/ericoliveiras/basic-crud-go/controllers"
	"github/ericoliveiras/basic-crud-go/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.Engine) {
	authRoutes := router.Group("/auth")

	authRoutes.POST("/login", controllers.Login)

	authRoutes.Use(middleware.AuthMiddleware)
	{
		authRoutes.POST("/logout", controllers.Logout)
	}
}