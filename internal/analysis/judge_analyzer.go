package analysis

import (
	"context"
	"tasker_go/internal/models"
)

type JudgeAnalyzer interface {
	Analyze(ctx context.Context, task *models.Task) (score uint, text string, err error)
}
