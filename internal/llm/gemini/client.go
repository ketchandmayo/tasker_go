package gemini

import (
	"context"
	"errors"
	"net"
	"net/http"
	"net/url"
	"os"
	"tasker_go/internal/llm"

	"golang.org/x/net/proxy"
	"google.golang.org/genai"
)

type lLMClient struct {
	model  string
	client *genai.Client
}

func NewClient(ctx context.Context, model string) (llm.LLMClient, error) {
	proxyURL := os.Getenv("GEMINI_SOCKS5_PROXY")
	clientConfig, err := geminiClientConfigFromProxy(proxyURL)
	if err != nil {
		return nil, err
	}

	aiClient, err := genai.NewClient(ctx, clientConfig)
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

func geminiClientConfigFromProxy(proxyURL string) (*genai.ClientConfig, error) {
	if proxyURL == "" {
		return nil, nil
	}

	parsedURL, err := url.Parse(proxyURL)
	if err != nil {
		return nil, err
	}

	dialer, err := proxy.FromURL(parsedURL, proxy.Direct)
	if err != nil {
		return nil, err
	}

	transport := &http.Transport{}
	if contextDialer, ok := dialer.(proxy.ContextDialer); ok {
		transport.DialContext = contextDialer.DialContext
	} else {
		transport.DialContext = func(_ context.Context, network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		}
	}

	return &genai.ClientConfig{
		HTTPClient: &http.Client{
			Transport: transport,
		},
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
