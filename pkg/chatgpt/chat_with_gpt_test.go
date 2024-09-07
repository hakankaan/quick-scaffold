package chatgpt_test

import (
	"context"
	"testing"

	"github.com/hakankaan/quick-scaffold/pkg/chatgpt"
	"github.com/sashabaranov/go-openai"
)

type MockChatGPTClient struct{}

func (m MockChatGPTClient) SendMessage(ctx context.Context, messages []openai.ChatCompletionMessage) (*openai.ChatCompletionMessage, error) {
	return &openai.ChatCompletionMessage{
		Role:    "assistant",
		Content: "Test response",
	}, nil
}

func TestChatWithGPT(t *testing.T) {
	ctx := context.Background()
	mockChatGPTClient := MockChatGPTClient{}
	lastMessages := []openai.ChatCompletionMessage{}
	userQuestion := "How to create a Go function?"
	rootPath := "./"

	filePaths := []string{}

	res, err := chatgpt.ChatWithGPT(ctx, mockChatGPTClient, rootPath, lastMessages, userQuestion, filePaths)
	if err != nil {
		t.Errorf("ChatWithGPT() error = %v", err)
	}

	if res == nil {
		t.Fatal("ChatWithGPT() response is nil")
		return
	}

	if res.Role != "assistant" {
		t.Errorf("ChatWithGPT response role is not assistant")
	}

	if res.Content != "Test response" {
		t.Errorf("ChatWithGPT response content is not Test response")
	}

}
