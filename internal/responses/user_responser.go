package responses

import (
	"github/ericoliveiras/basic-crud-go/internal/models"
	"time"
)

type UserResponse struct {
	ID        string        `json:"id"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Email     string        `json:"email"`
	Tasks     []models.Task `json:"tasks"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
