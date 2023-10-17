package controllers

import (
	"github/ericoliveiras/basic-crud-go/internal/handlers"
	req "github/ericoliveiras/basic-crud-go/internal/requests"
	"github/ericoliveiras/basic-crud-go/internal/responses"
	s "github/ericoliveiras/basic-crud-go/internal/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type UserController struct {
	UserService *s.UserService
}

func (c *UserController) Create(ctx *gin.Context) {
	var userRequest req.CreateUserRequest
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"ERROR::": err.Error()})
		return
	}

	if err := c.UserService.Create(&userRequest); err == s.ErrUserAlreadyExists {
		ctx.JSON(http.StatusConflict, gin.H{"ERROR::": err.Error()})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"ERROR::": "Failed to create user"})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (c *UserController) Read(ctx *gin.Context) {
	var user responses.UserResponse

	userId, err := handlers.GetUsertIdFromClaims(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ERROR::": err.Error()})
		return
	}

	user, err = c.UserService.Read(userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"ERROR::": err.Error()})
		return
	}

	formatUser := responses.UserResponse{
		ID:        user.ID,
		FirstName: cases.Title(language.BrazilianPortuguese).String(user.FirstName),
		LastName:  cases.Title(language.BrazilianPortuguese).String(user.LastName),
		Email:     user.Email,
		Tasks:     user.Tasks,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"data": formatUser})
}

func (c *UserController) Update(ctx *gin.Context) {
	userId, err := handlers.GetUsertIdFromClaims(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ERROR::": err.Error()})
		return
	}

	_, err = c.UserService.Read(userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"ERROR::": err.Error()})
		return
	}

	var updateUser req.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&updateUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.UserService.Update(userId, updateUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *UserController) Delete(ctx *gin.Context) {
	userId, err := handlers.GetUsertIdFromClaims(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ERROR::": err.Error()})
		return
	}

	if err := c.UserService.Delete(userId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ERROR::": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
