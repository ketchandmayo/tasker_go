package dto

import (
	"strings"
	"tasker_go/internal/models"
)

func (r PatchTaskRequest) Apply(t *models.Task) {
	if r.Title != nil {
		t.Title = strings.TrimSpace(*r.Title)
	}

	if r.Description != nil {
		t.Description = *r.Description
	}

	if r.Status != nil {
		t.Status = *r.Status
	}
}

func (r CreateTaskRequest) ToModel(userID uint) models.Task {
	return models.Task{
		Title:       strings.TrimSpace(r.Title),
		Description: r.Description,
		UserId:      userID,
		Status:      "new",
	}
}

func TaskToDetailResponse(model models.Task) TaskDetailResponse {
	return TaskDetailResponse{
		ID:          model.ID,
		Title:       model.Title,
		Description: model.Description,
		Status:      model.Status,
		CreatedAt:   model.CreatedAt,
	}
}

func TaskToListItem(model models.Task) TaskListItemResponse {
	return TaskListItemResponse{
		ID:     model.ID,
		Title:  model.Title,
		Status: model.Status,
	}
}
