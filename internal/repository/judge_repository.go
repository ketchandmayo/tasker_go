package repository

import (
	"context"
	"tasker_go/internal/models"
)

type JudgeRepository interface {
	FindByTaskID(ctx context.Context, taskID uint) (*models.Judge, error)
	Create(ctx context.Context, judge *models.Judge) error
}
