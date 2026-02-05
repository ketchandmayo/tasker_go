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

func (j judgeAnalyzer) PreliminaryJudge(task *models.Task) (score uint, text string) {
	score = j.ScoreTask(task)
	text = fallbackJudgeComment(score)
	return score, text
}

func (j judgeAnalyzer) Analyze(ctx context.Context, task *models.Task) (score uint, text string, err error) {
	score = j.ScoreTask(task)
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

func (j judgeAnalyzer) ScoreTask(task *models.Task) uint {
	title := strings.TrimSpace(task.Title)
	desc := strings.TrimSpace(task.Description)

	score := rateByLength(desc)
	if task.Status == "done" {
		score++
	}
	if title == "" {
		score = 1
	}

	if score < 1 {
		score = 1
	}
	if score > 10 {
		score = 10
	}

	return score
}

func rateByLength(text string) uint {
	length := len([]rune(text))

	if length > 100 {
		return 9
	}

	score := (length * 9) / 100
	return uint(score)
}

func buildJudgePrompt(title, desc, status string, score uint) string {
	return fmt.Sprintf(
		`Ты мудрый и ироничный арабский судья задач.
				Твоя задача на оценку;
				Заголовок: "%s"
				Описание: "%s"
				Статус: %s
				Твоя оценка: %d/10
				
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
