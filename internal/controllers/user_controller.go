package controllers

import (
	req "github/ericoliveiras/basic-crud-go/internal/requests"
	s "github/ericoliveiras/basic-crud-go/internal/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *s.UserService
}

func (c *UserController) Create(ctx *gin.Context) {
	var userRequest req.CreateUserRequest
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.UserService.Create(&userRequest); err == s.ErrUserAlreadyExists {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *UserController) Read(ctx *gin.Context) {
	
}
