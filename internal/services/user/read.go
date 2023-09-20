package services

import (
	"github/ericoliveiras/basic-crud-go/internal/responses"
)

func (s *UserService) Read(id string) (responses.UserResponse, error) {
	var userResponse responses.UserResponse
	user, err := s.UserRepository.GetById(id)
	if err != nil {
		return responses.UserResponse{}, err
	}

	userResponse = responses.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userResponse, nil
}
