package application

import (
	"github/ericoliveiras/basic-crud-go/internal/config"
	"github/ericoliveiras/basic-crud-go/internal/routes"
	"github/ericoliveiras/basic-crud-go/internal/server"
	"log"
)

func Start(cfg *config.Config) {
	app := server.NewServer(cfg)

	app.Gin.Use(config.LoadCorsConfig())

	routes.UserRoutes(app)
	routes.AuthRoutes(app)
	routes.TaskRoutes(app)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
