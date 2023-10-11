package services

import "github/ericoliveiras/basic-crud-go/internal/models"

func (s *TaskService) ReadAll(userId string) []models.Task {
	var tasks []models.Task

	tasks = s.TaskRepository.GetByUserId(userId)

	return tasks
}
