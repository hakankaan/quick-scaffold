package chatgpt

import (
	"context"
	"testing"

	"github.com/sashabaranov/go-openai"
)

type MockChatGPTClient struct{}

func (m *MockChatGPTClient) SendMessage(ctx context.Context, messages []openai.ChatCompletionMessage) (*openai.ChatCompletionMessage, error) {
	return &openai.ChatCompletionMessage{
		Role:    "assistant",
		Content: "Paris is the capital of France.",
	}, nil
}

func TestSendMessage(t *testing.T) {
	client := &MockChatGPTClient{}

	ctx := context.Background()
	messages := []openai.ChatCompletionMessage{
		{
			Role:    "system",
			Content: "You are a helpful assistant.",
		},
		{
			Role:    "user",
			Content: "What is the capital of France?",
		},
	}

	resp, err := client.SendMessage(ctx, messages)
	if err != nil {
		t.Fatalf("SendMessage() error: %v", err)
	}

	if resp.Role != "assistant" {
		t.Errorf("Expected role to be 'assistant', got: %v", resp.Role)
	}

	expectedContent := "Paris is the capital of France."
	if resp.Content != expectedContent {
		t.Errorf("Expected content: %v, got: %v", expectedContent, resp.Content)
	}
}
