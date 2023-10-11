package repositories

import (
	"github/ericoliveiras/basic-crud-go/internal/models"
	"github/ericoliveiras/basic-crud-go/internal/requests"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user *models.User) error
	GetByEmail(email string) bool
	GetById(id string) (models.User, error)
	Update(user *models.User, updateUser *requests.UpdateUserRequest) error
	Delete(id string) error
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (u *UserRepository) Create(user *models.User) error {
	return u.DB.Create(&user).Error
}

func (u *UserRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *UserRepository) GetById(id string) (models.User, error) {
	var user models.User
	if err := u.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *UserRepository) Update(user *models.User, updateUser *requests.UpdateUserRequest) error {
	if err := u.DB.Model(&user).Updates(&updateUser).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) Delete(id string) error {
	var user models.User
	if err := u.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
