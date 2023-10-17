package services

import (
	"github/ericoliveiras/basic-crud-go/internal/models"
	r "github/ericoliveiras/basic-crud-go/internal/repositories"
)

type UserServiceWrapper interface {
	Create(user *models.User) error
}

type UserService struct {
	UserRepository *r.UserRepository
	TaskRepository *r.TaskRepository
}

func NewUserService(userRepository *r.UserRepository, taskRepository *r.TaskRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
		TaskRepository: taskRepository,
	}
}
