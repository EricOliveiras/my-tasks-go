package main

import (
	"github/ericoliveiras/basic-crud-go/config"
	"github/ericoliveiras/basic-crud-go/database"
	"github/ericoliveiras/basic-crud-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.ConnectDatabase()

	routes.UserRouter(router)
	routes.TaskRouter(router)

	router.Run(config.BaseURL())
}
