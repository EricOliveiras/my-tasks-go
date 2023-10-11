package services

import (
	"github/ericoliveiras/basic-crud-go/internal/models"
	"github/ericoliveiras/basic-crud-go/internal/requests"
)

func (s *TaskService) Update(id string, updateTask requests.UpdateTaskRequest) error {
	var task models.Task

	task, err := s.TaskRepository.GetById(id)
	if err != nil {
		return err
	}

	if err := s.TaskRepository.Update(&task, &updateTask); err != nil {
		return err
	}

	return nil
}
