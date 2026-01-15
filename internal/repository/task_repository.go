package repository

import (
	"context"
	"tasker_go/internal/models"
)

type TaskRepository interface {
	Create(ctx context.Context, task *models.Task) error
	FindByUser(ctx context.Context, userID uint) ([]models.Task, error)
	FindByID(ctx context.Context, userID uint, id uint) (*models.Task, error)
	Update(ctx context.Context, task *models.Task) error
	Delete(ctx context.Context, task *models.Task) error
}
