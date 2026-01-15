package service

import (
	"context"
	"tasker_go/internal/models"
)

type TaskService interface {
	Create(ctx context.Context, userID uint, task *models.Task) (uint, error)
	List(ctx context.Context, userID uint) ([]models.Task, error)
	Get(ctx context.Context, userID uint, id uint) (*models.Task, error)
	Update(ctx context.Context, userID uint, task *models.Task) error
	Delete(ctx context.Context, userID uint, id uint) error
}
