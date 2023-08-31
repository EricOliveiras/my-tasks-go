package services

import "github/ericoliveiras/basic-crud-go/internal/models"

func (s *UserService) Read(id string) (models.User, error) {
	user, err := s.UserRepository.GetById(id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}