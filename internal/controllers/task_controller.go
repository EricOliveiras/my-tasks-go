package controllers

import (
	"github/ericoliveiras/basic-crud-go/internal/handlers"
	"github/ericoliveiras/basic-crud-go/internal/models"
	"github/ericoliveiras/basic-crud-go/internal/requests"
	s "github/ericoliveiras/basic-crud-go/internal/services/task"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskService *s.TaskService
}

func (c *TaskController) Create(ctx *gin.Context) {
	userId, err := handlers.GetUsertIdFromClaims(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ERROR::": err.Error()})
		return
	}

	var newTask requests.CreateTaskRequest
	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ERROR::": err.Error()})
		return
	}

	newTask.UserId = userId
	newTask.Finished = false

	var task models.Task
	task, err = c.TaskService.Create(&newTask)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ERROR::": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": task})
}

func (c *TaskController) Read(ctx *gin.Context) {
	var task models.Task

	_, err := handlers.GetUsertIdFromClaims(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ERROR::": err.Error()})
		return
	}

	id := ctx.Param("id")

	task, err = c.TaskService.Read(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"ERROR::": "Task not found or not exists."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data:": task})
}

func (c *TaskController) ReadAll(ctx *gin.Context) {
	var tasks []models.Task

	userId, err := handlers.GetUsertIdFromClaims(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ERROR::": err.Error()})
		return
	}

	tasks = c.TaskService.ReadAll(userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"ERROR::": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data:": tasks})
}

func (c *TaskController) Update(ctx *gin.Context) {
	var updateTask requests.UpdateTaskRequest

	if err := ctx.ShouldBindJSON(&updateTask); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ERROR::": err.Error()})
		return
	}

	if err := c.TaskService.Update(updateTask.ID, updateTask); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ERROR::": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *TaskController) Delete(ctx *gin.Context) {
	var req requests.TaskRequestId

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ERROR::": err.Error()})
		return
	}

	if err := c.TaskService.Delete(req.ID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"ERROR::": "Task not found or not exists."})
		return
	}

	ctx.Status(http.StatusOK)
}
