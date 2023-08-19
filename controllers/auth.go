package controllers

import (
	"github/ericoliveiras/basic-crud-go/database"
	"github/ericoliveiras/basic-crud-go/handlers"
	"github/ericoliveiras/basic-crud-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	var user models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ERROR::": err.Error()})
		return
	}

	defer c.Request.Body.Close()

	email := req.Email
	password := req.Password

	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"ERROR::": "Invalid credentials."})
		return
	}

	if err := handlers.ComparePassword(user.Password, password); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"ERROR::": "Invalid credentials."})
		return
	}

	token, err := handlers.GenerateToken(user.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"ERROR::": "Error generating token."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
