package models

import "time"

type User struct {
	ID        string    `json:"id" gorm:"primaryKey;unique"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
