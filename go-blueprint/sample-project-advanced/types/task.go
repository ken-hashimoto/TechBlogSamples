package types

import "github.com/google/uuid"

type Task struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   string    `json:"created_at"`
	IsCompleted bool      `json:"is_completed"`
}

type CreateTaskRequest struct {
	TaskName    string `json:"name"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}
