package conversation

import "github.com/sashabaranov/go-openai"

func AddProjectSummaryToConversationHistory(conversationHistory *[]openai.ChatCompletionMessage, projectSummary string) {
	*conversationHistory = append(*conversationHistory, openai.ChatCompletionMessage{
		Role:    "assistant",
		Content: "I need full project structure to understand the project better. Can you provide it?",
	})
	*conversationHistory = append(*conversationHistory, openai.ChatCompletionMessage{
		Role:    "user",
		Content: projectSummary,
	})
}
