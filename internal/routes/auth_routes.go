package routes

import (
	"github/ericoliveiras/basic-crud-go/internal/controllers"
	"github/ericoliveiras/basic-crud-go/internal/repositories"
	s "github/ericoliveiras/basic-crud-go/internal/server"
	services "github/ericoliveiras/basic-crud-go/internal/services/auth"
)

func AuthRoutes(server *s.Server) {
	userRepository := repositories.NewUserRepository(server.DB)
	authService := services.NewAuthService(userRepository)
	loginController := controllers.AuthController{AuthService: authService}

	authRoutes := server.Gin.Group("/auth")

	authRoutes.POST("", loginController.Login)
}
