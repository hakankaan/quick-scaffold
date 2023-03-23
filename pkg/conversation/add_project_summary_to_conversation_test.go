package conversation

import (
	"testing"

	"github.com/sashabaranov/go-openai"
)

func TestAddProjectSummaryToConversationHistory(t *testing.T) {
	conversationHistory := []openai.ChatCompletionMessage{
		{
			Role:    "user",
			Content: "This is a message.",
		},
	}

	projectSummary := "This is the project structure:\n- main.go\n- README.md"

	AddProjectSummaryToConversationHistory(&conversationHistory, projectSummary)

	if len(conversationHistory) != 3 {
		t.Fatalf("Expected conversation history length to be 3, got: %d", len(conversationHistory))
	}

	assistantMessage := conversationHistory[1]
	if assistantMessage.Role != "assistant" || assistantMessage.Content != "I need full project structure to understand the project better. Can you provide it?" {
		t.Errorf("Unexpected assistant message: %+v", assistantMessage)
	}

	userMessage := conversationHistory[2]
	if userMessage.Role != "user" || userMessage.Content != projectSummary {
		t.Errorf("Unexpected user message: %+v", userMessage)
	}
}
