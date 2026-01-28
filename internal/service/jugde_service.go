package service

import (
	"context"
	"tasker_go/internal/models"
)

type JudgeService interface {
	GetByTaskID(ctx context.Context, userId uint, taskId uint) (*models.Judge, error)
}
