package requests

import "time"

type CreateTaskRequest struct {
	UserId      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Finished    bool   `json:"finished"`
}

type UpdateTaskRequest struct {
	TaskRequestId
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Finished    bool      `json:"finished"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type TaskRequestId struct {
	ID string `json:"id"`
}
