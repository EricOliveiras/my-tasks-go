package models

import (
	"time"
)

type Task struct {
	ID          string    `json:"id" gorm:"primaryKey;unique"`
	UserId      string    `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Finished    bool      `json:"finished"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
