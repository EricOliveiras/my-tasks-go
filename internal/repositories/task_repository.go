package repositories

import (
	"github/ericoliveiras/basic-crud-go/internal/models"
	"github/ericoliveiras/basic-crud-go/internal/requests"

	"gorm.io/gorm"
)

type ITaskInterface interface {
	Create(task *models.Task) error
	GetById(id string) (models.Task, error)
	GetByUserId(userId string) []models.Task
	Update(task *models.Task, updateTask *requests.UpdateTaskRequest)
	Delete(id string)
}

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (t *TaskRepository) Create(task *models.Task) error {
	return t.DB.Create(&task).Error
}

func (t *TaskRepository) GetById(id string) (models.Task, error) {
	var task models.Task
	if err := t.DB.Where("id = ?", id).First(&task).Error; err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (t *TaskRepository) GetByUserId(userId string) []models.Task {
	var tasks []models.Task
	t.DB.Find(&tasks).Where("user_id = ?", userId)

	return tasks
}

func (t *TaskRepository) Update(task *models.Task, updateTask *requests.UpdateTaskRequest) error {
	if err := t.DB.Model(&task).Updates(&updateTask).Error; err != nil {
		return err
	}

	return nil
}

func (t *TaskRepository) Delete(id string) error {
	var task models.Task
	if err := t.DB.Where("id = ?", id).Delete(&task).Error; err != nil {
		return err
	}

	return nil
}
