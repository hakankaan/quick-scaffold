package conversation

import (
	"strings"
	"testing"
)

func TestCreateInitialConversationHistory(t *testing.T) {
	userNeed := "I need a function that takes two integers and returns their sum."

	messages := CreateInitialConversationHistory(userNeed)

	if len(messages) == 0 {
		t.Fatal("Expected conversation history to have messages")
	}

	lastMessage := messages[len(messages)-1]
	if lastMessage.Role != "user" {
		t.Errorf("Expected last message role to be 'user', got: %v", lastMessage.Role)
	}

	if !strings.Contains(lastMessage.Content, userNeed) {
		t.Errorf("Expected last message content to contain user need: %v", userNeed)
	}
}
