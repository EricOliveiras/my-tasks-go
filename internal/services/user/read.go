package services

import (
	"github/ericoliveiras/basic-crud-go/internal/models"
	"github/ericoliveiras/basic-crud-go/internal/responses"
)

func (s *UserService) Read(id string) (responses.UserResponse, error) {
	var userResponse responses.UserResponse
	user, err := s.UserRepository.GetById(id)
	if err != nil {
		return responses.UserResponse{}, err
	}

	var tasks []models.Task
	tasks = s.TaskRepository.GetByUserId(id)

	userResponse = responses.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Tasks:     tasks,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userResponse, nil
}
