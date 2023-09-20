package services

import (
	r "github/ericoliveiras/basic-crud-go/internal/repositories"
)

type AuthServiceWrapper interface {
	Login(email, password string) (string, error)
}

type AuthService struct {
	UserRepository *r.UserRepository
}

func NewAuthService(authService *r.UserRepository) *AuthService {
	return &AuthService{UserRepository: authService}
}
