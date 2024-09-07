package chatgpt

import (
	"context"

	"github.com/hakankaan/quick-scaffold/pkg/conversation"
	"github.com/hakankaan/quick-scaffold/pkg/fileutils"
	"github.com/sashabaranov/go-openai"
)

func ChatWithGPT(ctx context.Context, chatGpt ChatGPTClientInterface, rootPath string, lastMessages []openai.ChatCompletionMessage, userQuestion string, filePaths []string) (*openai.ChatCompletionMessage, error) {
	// Initialize conversation
	newConversation := conversation.CreateInitialConversationHistory(userQuestion)

	// Add folder structure
	folderStructure, err := fileutils.GetFolderStructure(rootPath)
	if err != nil {
		return nil, err
	}

	// merge string slices into one string with newlines
	stringFolderStructure := ""
	for _, path := range folderStructure {
		stringFolderStructure += path + "\n"
	}

	conversation.AddProjectSummaryToConversationHistory(&newConversation, stringFolderStructure)

	// Add file contents
	filesWithContent, err := fileutils.GetFilesWithContent(filePaths...)
	if err != nil {
		return nil, err
	}

	conversation.AddFileNamesAndContentsToConversationHistory(&newConversation, filesWithContent)

	newConversation = append(newConversation, lastMessages...)

	// Send conversation to ChatGPT and get the response
	response, err := chatGpt.SendMessage(ctx, newConversation)
	if err != nil {
		return nil, err
	}

	return response, nil
}
