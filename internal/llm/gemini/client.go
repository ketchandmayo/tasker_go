package gemini

import (
	"context"
	"errors"
	"os"
	"tasker_go/internal/llm"

	"google.golang.org/genai"
)

type lLMClient struct {
	model  string
	client *genai.Client
}

func NewClient(ctx context.Context, model string) (llm.LLMClient, error) {
	aiClient, err := genai.NewClient(ctx, nil)
	if err != nil {
		return nil, err
	}

	if model == "" {
		model = os.Getenv("GEMINI_MODEL")
	}

	return &lLMClient{
		client: aiClient,
		model:  model,
	}, nil
}

func (c *lLMClient) Generate(ctx context.Context, prompt string) (string, error) {

	result, err := c.client.Models.GenerateContent(
		ctx,
		c.model,
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return "", err
	}

	text := result.Text()
	if text == "" {
		return "", errors.New("empty response from gemini")
	}

	return text, nil
}
