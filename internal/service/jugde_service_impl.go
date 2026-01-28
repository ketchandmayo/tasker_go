package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"tasker_go/internal/analysis"
	"tasker_go/internal/models"
	"tasker_go/internal/repository"
)

type judgeService struct {
	taskRepo  repository.TaskRepository
	judgeRepo repository.JudgeRepository
	analyzer  analysis.JudgeAnalyzer
}

func NewJudgeService(tRepo repository.TaskRepository, jRepo repository.JudgeRepository, analyzer analysis.JudgeAnalyzer) JudgeService {
	return &judgeService{
		taskRepo:  tRepo,
		judgeRepo: jRepo,
		analyzer:  analyzer,
	}
}

func (j *judgeService) GetByTaskID(ctx context.Context, userId uint, taskId uint) (*models.Judge, error) {
	task, err := j.taskRepo.FindByID(ctx, userId, taskId)
	if err != nil {
		return nil, err
	}
	taskHash := fingerprint(task)

	existingJudge, err := j.judgeRepo.FindByTaskID(ctx, task.ID)
	if err != nil {
		return nil, ErrTaskNotFound
	}

	if existingJudge.TaskHash == taskHash {
		return existingJudge, nil
	}

	score, text, err := j.analyzer.Analyze(ctx, task)
	if err != nil {
		return nil, err
	}

	judge := models.Judge{
		TaskID:   task.ID,
		TaskHash: taskHash,
		Score:    score,
		Text:     text,
	}

	if err := j.judgeRepo.Create(ctx, &judge); err != nil {
		return nil, err
	}

	return existingJudge, nil
}

func fingerprint(t *models.Task) string {
	h := sha256.New()

	h.Write([]byte(t.Title))
	h.Write([]byte(t.Description))
	h.Write([]byte(t.Status))

	return hex.EncodeToString(h.Sum(nil))
}
