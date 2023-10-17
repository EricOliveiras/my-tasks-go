package services

import (
	"errors"
	"github/ericoliveiras/basic-crud-go/internal/builders"
	"github/ericoliveiras/basic-crud-go/internal/models"
	"github/ericoliveiras/basic-crud-go/internal/requests"

	"github.com/google/uuid"
)

func (s *TaskService) Create(task *requests.CreateTaskRequest) (models.Task, error) {
	id := uuid.New().String()
	
	createTask := builders.NewTaskBuilder().
		SetID(id).
		SetUserId(task.UserId).
		SetTitle(task.Title).
		SetDescription(task.Description).
		SetFinished(task.Finished).
		Build()

	if err := s.TaskRepository.Create(&createTask); err != nil {
		return models.Task{}, errors.New("Error to create a new task.")
	}

	return createTask, nil
}
