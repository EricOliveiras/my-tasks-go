package routes

import (
	"github/ericoliveiras/basic-crud-go/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser)
	router.GET("/user", controllers.ReadUsers)
	router.GET("/user/:id", controllers.ReadUser)
	router.PUT("/user/:id", controllers.UpdateUser)
	router.DELETE("/user/:id", controllers.DeleteUser)
}