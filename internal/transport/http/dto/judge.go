package dto

type JudgeResponse struct {
	TaskID uint   `json:"task_id"`
	Score  uint   `json:"score"`
	Text   string `json:"text"`
}
