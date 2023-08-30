package services

import (
	"errors"
	"github/ericoliveiras/basic-crud-go/internal/builders"
	"github/ericoliveiras/basic-crud-go/internal/handlers"
	"github/ericoliveiras/basic-crud-go/internal/requests"

	"github.com/google/uuid"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

func (s *UserService) Create(user *requests.CreateUserRequest) error {
	_, err := s.UserRepository.GetByEmail(user.Email)
	if err == nil {
		return ErrUserAlreadyExists
	}

	hashedPassword, err := handlers.HashPassword(user.Password)
	if err != nil {
		return err
	}

	id := uuid.New().String()

	createUser := builders.NewUserBuilder().
		SetID(id).
		SetFirstname(user.FirstName).
		SetLastname(user.LastName).
		SetEmail(user.Email).
		SetPassword(hashedPassword).
		Build()

	err = s.UserRepository.Create(&createUser)
	if err != nil {
		return err
	}

	return nil
}
