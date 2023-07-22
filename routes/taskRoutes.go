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
		taskRoutes.GET("/:id", controllers.ReadTask)
		taskRoutes.PUT("/:id", controllers.UpdateTask)
		taskRoutes.DELETE("/:id", controllers.DeleteTask)
	}
}
