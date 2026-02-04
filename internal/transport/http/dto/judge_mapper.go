package dto

import "tasker_go/internal/models"

func JudgeToResponse(judge *models.Judge, preliminary bool) JudgeResponse {
	return JudgeResponse{
		TaskID:      judge.TaskID,
		Score:       judge.Score,
		Text:        judge.Text,
		Preliminary: preliminary,
	}
}
