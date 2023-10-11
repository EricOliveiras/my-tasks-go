package services

import (
	"github/ericoliveiras/basic-crud-go/internal/models"
	r "github/ericoliveiras/basic-crud-go/internal/repositories"
)

type TaskServiceWrapper interface {
	Create(task *models.Task) error
}

type TaskService struct {
	TaskRepository *r.TaskRepository
}

func NewTaskService(taskService *r.TaskRepository) *TaskService {
	return &TaskService{TaskRepository: taskService}
}
