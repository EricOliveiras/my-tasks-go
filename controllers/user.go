package controllers

import (
	"github/ericoliveiras/basic-crud-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateUserInput struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ERROR::": err.Error()})
		return
	}

	var existingUser models.User
	result := models.DB.Where("email = ?", input.Email).First(&existingUser)
	if result.Error == nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"ERROR::": "Email already exists."})
		return
	}

	input.ID = uuid.New().String()

	user := models.User{
		ID:        input.ID,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
	}

	models.DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func ReadUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}
