package gorm

import (
	"context"
	"tasker_go/internal/models"
	"tasker_go/internal/repository"

	"gorm.io/gorm"
)

type judgeRepository struct {
	db *gorm.DB
}

func NewJudgeRepository(db *gorm.DB) repository.JudgeRepository {
	return &judgeRepository{db: db}
}
func (j judgeRepository) FindByTaskID(ctx context.Context, taskID uint) (*models.Judge, error) {
	var judge models.Judge

	err := j.db.WithContext(ctx).
		Where("task_id = ?", taskID).
		First(&judge).
		Error

	if err != nil {
		return nil, err
	}

	return &judge, nil
}

func (j judgeRepository) Create(ctx context.Context, judge *models.Judge) error {
	return j.db.WithContext(ctx).
		Create(judge).
		Error
}

func (j judgeRepository) Update(ctx context.Context, judge *models.Judge) error {
	return j.db.WithContext(ctx).
		Save(judge).
		Error
}
