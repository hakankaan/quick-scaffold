package conversation

import (
	"fmt"

	"github.com/hakankaan/quick-scaffold/pkg/fileutils"
	"github.com/sashabaranov/go-openai"
)

func AddFileNamesAndContentsToConversationHistory(conversationHistory *[]openai.ChatCompletionMessage, codeFiles []fileutils.FileContent) {
	neededFileNamesAndContents := "I need content of the following files in the project to response you better."
	for _, file := range codeFiles {
		neededFileNamesAndContents += fmt.Sprintf("\nNeed_content_of_file: %s", file.Path)
	}
	*conversationHistory = append(*conversationHistory, openai.ChatCompletionMessage{
		Role:    "assistant",
		Content: neededFileNamesAndContents,
	})

	fileNamesAndContents := ""
	for _, file := range codeFiles {
		fileNamesAndContents += fmt.Sprintf("\nFilecontent: %s\n```\n%s\n```", file.Path, file.Content)
	}
	*conversationHistory = append(*conversationHistory, openai.ChatCompletionMessage{
		Role:    "user",
		Content: fileNamesAndContents,
	})
}
