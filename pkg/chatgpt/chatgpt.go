package chatgpt

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type ChatGPTClientInterface interface {
	SendMessage(ctx context.Context, messages []openai.ChatCompletionMessage) (*openai.ChatCompletionMessage, error)
}

type ChatGPTClient struct {
	client *openai.Client
}

func New(apiKey string) (ChatGPTClientInterface, error) {
	client := openai.NewClient(apiKey)

	return &ChatGPTClient{
		client: client,
	}, nil
}

func (c *ChatGPTClient) SendMessage(ctx context.Context, messages []openai.ChatCompletionMessage) (*openai.ChatCompletionMessage, error) {
	request := openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo0301,
		Messages: messages,
	}

	resp, err := c.client.CreateChatCompletion(ctx, request)
	if err != nil {
		return nil, err
	}

	return &resp.Choices[0].Message, nil
}
