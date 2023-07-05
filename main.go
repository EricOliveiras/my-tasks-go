package main

import (
	"github/ericoliveiras/basic-crud-go/config"
	"github/ericoliveiras/basic-crud-go/controllers"
	"github/ericoliveiras/basic-crud-go/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.ConnectDatabase()

	router.POST("/user", controllers.CreateUser)
	router.GET("/user", controllers.ReadUsers)
	router.GET("/user/:id", controllers.ReadUser)

	router.Run(config.BaseURL())
}
