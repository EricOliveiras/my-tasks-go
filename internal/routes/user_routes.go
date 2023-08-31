package routes

import (
	"github/ericoliveiras/basic-crud-go/internal/controllers"
	"github/ericoliveiras/basic-crud-go/internal/repositories"
	s "github/ericoliveiras/basic-crud-go/internal/server"
	services "github/ericoliveiras/basic-crud-go/internal/services/user"
)

func UserRoutes(server *s.Server) {
	userRepository := repositories.NewUserRepository(server.DB)
	userService := services.NewUserService(userRepository)
	userController := controllers.UserController{UserService: userService}

	userRoutes := server.Gin.Group("/user")

	userRoutes.POST("", userController.Create)
}
