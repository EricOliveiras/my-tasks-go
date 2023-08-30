package services

import (
	"github/ericoliveiras/basic-crud-go/internal/models"
	r "github/ericoliveiras/basic-crud-go/internal/repositories"
)

type ServiceWrapper interface {
	Create(user *models.User) error
}

type UserService struct {
	UserRepository *r.UserRepository
}

func NewUserService(userRepository *r.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}
