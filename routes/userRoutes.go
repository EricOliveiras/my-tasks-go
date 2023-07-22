package routes

import (
	"github/ericoliveiras/basic-crud-go/controllers"
	"github/ericoliveiras/basic-crud-go/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	userRoutes := router.Group("/user")

	userRoutes.POST("", controllers.CreateUser)

	userRoutes.Use(middleware.AuthMiddleware)
	{
		userRoutes.GET("/me", controllers.Me)
		userRoutes.PUT("", controllers.UpdateUser)
		userRoutes.DELETE("", controllers.DeleteUser)
	}
}
