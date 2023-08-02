package routes

import (
	"github/ericoliveiras/basic-crud-go/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.Engine) {
	authRoutes := router.Group("/auth")

	authRoutes.POST("/login", controllers.Login)
}
