package routes

import (
	"github/ericoliveiras/basic-crud-go/controllers"
	"github/ericoliveiras/basic-crud-go/middleware"

	"github.com/gin-gonic/gin"
)

func TaskRouter(router *gin.Engine) {
	taskRoutes := router.Group("/task")

	taskRoutes.Use(middleware.AuthMiddleware)
	{
		taskRoutes.POST("", controllers.CreateTask)
		taskRoutes.GET("", controllers.ReadTasks)
		taskRoutes.PUT("", controllers.UpdateTask)
		taskRoutes.DELETE("", controllers.DeleteTask)
	}
}
