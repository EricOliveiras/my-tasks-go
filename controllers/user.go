package controllers

import (
	"github/ericoliveiras/basic-crud-go/database"
	"github/ericoliveiras/basic-crud-go/handlers"
	"github/ericoliveiras/basic-crud-go/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateUserInput struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  *string `json:"password"`
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ERROR::": err.Error()})
		return
	}

	defer c.Request.Body.Close()

	var existingUser models.User
	if err := database.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"ERROR::": "Email already exists."})
		return
	}

	input.ID = uuid.New().String()

	hashPassword, err := handlers.HashPassword(input.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ERROR::": err.Error()})
		return
	}

	user := models.User{
		ID:        input.ID,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  hashPassword,
	}

	database.DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func Me(c *gin.Context) {
	var user models.User

	userId, err := handlers.GetUsertIdFromClaims(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ERROR::": err.Error()})
		return
	}

	if err := database.DB.Preload("Task").Where("id = ?", userId).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"ERROR::": "User not found or not exists."})
		return
	}

	user = models.User{
		ID:        user.ID,
		FirstName: strings.Title(user.FirstName),
		LastName:  strings.Title(user.LastName),
		Email:     user.Email,
		Task:      user.Task,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	var user models.User

	userId, err := handlers.GetUsertIdFromClaims(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ERROR::": err.Error()})
		return
	}

	if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"ERROR::": "User not found or not exists."})
		return
	}

	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ERROR::": err.Error()})
		return
	}

	defer c.Request.Body.Close()

	updateUser := models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	if input.Password != nil {
		newPass, err := handlers.HashPassword(*input.Password)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ERROR::": err.Error()})
			return
		}	

		updateUser.Password = newPass
	}

	database.DB.Model(&user).Updates(&updateUser)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	var user models.User

	userId, err := handlers.GetUsertIdFromClaims(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ERROR::": err.Error()})
		return
	}

	if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"ERROR::": "User not found or not exists."})
		return
	}

	if err := database.DB.Where("user_id = ?", userId).Delete(&models.Task{}).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ERROR::": "Failed to delete tasks."})
		return
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ERROR::": "Failed to delete user."})
		return
	}

	c.Status(http.StatusOK)
}
