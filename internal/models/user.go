package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        string `json:"id" gorm:"primaryKey;unique"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	Task      []Task `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
