package routes

import (
	"github/ericoliveiras/basic-crud-go/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.Engine) {
	router.POST("/login", controllers.Login)
}