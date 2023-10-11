package services

import (
	"github/ericoliveiras/basic-crud-go/internal/handlers"
	"github/ericoliveiras/basic-crud-go/internal/models"
	"github/ericoliveiras/basic-crud-go/internal/requests"
)

func (s *UserService) Update(id string, updateUser requests.UpdateUserRequest) error {
	var user models.User

	user, err := s.UserRepository.GetById(id)
	if err != nil {
		return err
	}

	if updateUser.Password != "" {
		updateUser.Password, err = handlers.HashPassword(updateUser.Password)
		if err != nil {
			return err
		}
	}

	if err := s.UserRepository.Update(&user, &updateUser); err != nil {
		return err
	}

	return nil
}
