package main

import (
	"github/ericoliveiras/basic-crud-go/controllers"
	"github/ericoliveiras/basic-crud-go/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.POST("/user", controllers.CreateUser)
	router.GET("/user", controllers.ReadUsers)

	router.Run("localhost:8080")
}