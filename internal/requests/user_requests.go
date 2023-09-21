package requests

import "time"

type CreateUserRequest struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BasicAuth
}

type UpdateUserRequest struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Password  string    `json:"password"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
