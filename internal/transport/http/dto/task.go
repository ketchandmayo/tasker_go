package dto

import "time"

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required,max=100"`
	Description string `json:"description"`
}

type TaskListItemResponse struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type TaskDetailResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type PatchTaskRequest struct {
	Title       *string `json:"title" validate:"omitempty,max=100"`
	Description *string `json:"description"`
	Status      *string `json:"status" validate:"omitempty,oneof=new done archived"`
}
