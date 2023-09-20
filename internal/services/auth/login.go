package services

import (
	"errors"
	"github/ericoliveiras/basic-crud-go/internal/handlers"
)

var InvalidCredentialError = errors.New("Invalid credentials.")

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.UserRepository.GetByEmail(email)
	if err != nil {
		return "", InvalidCredentialError
	}

	if err := handlers.ComparePassword(user.Password, password); err != nil {
		return "", InvalidCredentialError
	}

	token, err := handlers.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
