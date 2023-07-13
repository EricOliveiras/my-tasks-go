package routes

import (
	"github/ericoliveiras/basic-crud-go/controllers"

	"github.com/gin-gonic/gin"
)

func TaskRouter(router *gin.Engine) {
	router.POST("/task", controllers.CreateTask)
	router.GET("/task", controllers.ReadTasks)
	router.GET("/task/:id", controllers.ReadTask)
	router.PUT("/task/:id", controllers.UpdateTask)
	router.DELETE("/task/:id", controllers.DeleteTask)
}
