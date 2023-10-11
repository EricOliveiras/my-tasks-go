package controllers

import (
	"github/ericoliveiras/basic-crud-go/internal/requests"
	s "github/ericoliveiras/basic-crud-go/internal/services/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *s.AuthService
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req requests.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ERROR::": err.Error()})
		return
	}

	defer ctx.Request.Body.Close()

	email := req.Email
	password := req.Password

	token, err := c.AuthService.Login(email, password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"ERROR::": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
