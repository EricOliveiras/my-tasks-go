package routes

import (
	"github/ericoliveiras/basic-crud-go/internal/controllers"
	"github/ericoliveiras/basic-crud-go/internal/middleware"
	"github/ericoliveiras/basic-crud-go/internal/repositories"
	s "github/ericoliveiras/basic-crud-go/internal/server"
	services "github/ericoliveiras/basic-crud-go/internal/services/task"
)

func TaskRoutes(server *s.Server) {
	taskRespository := repositories.NewTaskRepository(server.DB)
	taskService := services.NewTaskService(taskRespository)
	taskController := controllers.TaskController{TaskService: taskService}

	taskRoutes := server.Gin.Group("/task")

	taskRoutes.Use(middleware.AuthMiddleware)
	{
		taskRoutes.POST("", taskController.Create)
		taskRoutes.GET("", taskController.ReadAll)
		taskRoutes.GET("/:id", taskController.Read)
		taskRoutes.PATCH("", taskController.Update)
		taskRoutes.DELETE("", taskController.Delete)
	}
}
