package responses

import "time"

type UserResponse struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
