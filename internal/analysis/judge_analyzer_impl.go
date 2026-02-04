package analysis

import (
	"context"
	"fmt"
	"strings"
	"tasker_go/internal/llm"
	"tasker_go/internal/models"
)

type judgeAnalyzer struct {
	llm llm.LLMClient
}

func NewJudgeAnalyzer(llm llm.LLMClient) JudgeAnalyzer {
	return &judgeAnalyzer{
		llm: llm,
	}
}

//curl --proxy socks5://ketch:ez1QYU2pM4@jdhlgkfjdhasdkas.polarrp.ru:55964 "https://generativelanguage.googleapis.com/v1beta/models/gemma-3-12b-it:generateContent"   -H "x-goog-api-key: AIzaSyBo7BKdOm6d-1GVwKEGWKEObGGzgIgEY1w"   -H 'Content-Type: application/json'   -X POST   -d '{
//"contents": [
//{
//"parts": [
//{
//"text": "Ты араб, продающий арбузы на рынке, ответь на вопрос `Привет, как у тебя дела?`"
//}
//]
//}
//]
//}'

func (j judgeAnalyzer) Analyze(ctx context.Context, task *models.Task) (score uint, text string, err error) {
	score = ScoreTask(task)
	title := strings.TrimSpace(task.Title)
	desc := strings.TrimSpace(task.Description)

	prompt := buildJudgePrompt(title, desc, task.Status, score)

	if j.llm != nil {
		text, err = j.llm.Generate(ctx, prompt)
		if err != nil {
			return 0, "", err
		}
	} else {
		text = fallbackJudgeComment(score)
	}

	return score, text, nil
}

func ScoreTask(task *models.Task) uint {
	title := strings.TrimSpace(task.Title)
	desc := strings.TrimSpace(task.Description)

	score := uint(5)

	if title == "" {
		score = 1
	}

	if len(desc) < 20 && score > 1 {
		score--
	}

	if task.Status == "done" {
		score++
	}

	if score < 1 {
		score = 1
	}
	if score > 10 {
		score = 10
	}

	return score
}

func PreliminaryJudge(task *models.Task) (score uint, text string) {
	score = ScoreTask(task)
	text = fallbackJudgeComment(score)
	return score, text
}

func buildJudgePrompt(title, desc, status string, score uint) string {
	return fmt.Sprintf(
		`Ты мудрый и ироничный арабский судья задач.
				Заголовок: "%s"
				Описание: "%s"
				Статус: %s
				Предварительная оценка: %d/10
				
				Дай короткий (не затягивай) комментарий с лёгким восточным юмором.`,
		title,
		desc,
		status,
		score,
	)
}

func fallbackJudgeComment(score uint) string {
	switch {
	case score <= 3:
		return "Брат, в этой задаче больше тумана, чем смысла."
	case score <= 6:
		return "Неплохо, но мудрость любит ясность."
	case score <= 8:
		return "Хорошая задача, путь понятен."
	default:
		return "Великолепно! Даже старейшины одобрили бы."
	}
}
