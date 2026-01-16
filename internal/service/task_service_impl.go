package service

import (
	"context"
	"tasker_go/internal/transport/http/dto"

	"tasker_go/internal/models"
	"tasker_go/internal/repository"
)

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) Create(ctx context.Context, userID uint, task *models.Task) (uint, error) {
	task.UserId = userID
	if err := s.repo.Create(ctx, task); err != nil {
		return 0, err
	}
	return task.ID, nil
}

func (s *taskService) List(ctx context.Context, userID uint) ([]models.Task, error) {
	return s.repo.FindByUser(ctx, userID)
}

func (s *taskService) Get(ctx context.Context, userID uint, id uint) (*models.Task, error) {
	task, err := s.repo.FindByID(ctx, userID, id)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, ErrTaskNotFound
	}
	return task, nil
}

func (s *taskService) Update(ctx context.Context, userID uint, taskID uint, req *dto.PatchTaskRequest) (*models.Task, error) {
	existing, err := s.repo.FindByID(ctx, userID, taskID)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, ErrTaskNotFound
	}

	req.Apply(existing)
	err = s.repo.Update(ctx, existing)
	if err != nil {
		return nil, err
	}

	return existing, nil
}

func (s *taskService) Delete(ctx context.Context, userID uint, id uint) error {
	task, err := s.repo.FindByID(ctx, userID, id)
	if err != nil {
		return err
	}
	if task == nil {
		return ErrTaskNotFound
	}
	return s.repo.Delete(ctx, task)
}
