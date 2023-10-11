package services

import "github/ericoliveiras/basic-crud-go/internal/models"

func (s *TaskService) Read(id string) (models.Task, error) {
	var task models.Task

	task, err := s.TaskRepository.GetById(id)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}
