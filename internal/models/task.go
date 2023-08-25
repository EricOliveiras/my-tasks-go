package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          string `json:"id" gorm:"primaryKey;unique"`
	UserId      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Finished    bool   `json:"finished"`
}
