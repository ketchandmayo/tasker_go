package handler

import (
	"tasker_go/internal/service"
)

type Handler struct {
	taskService service.TaskService
}

func New(taskService service.TaskService) *Handler {
	return &Handler{
		taskService: taskService,
	}
}
