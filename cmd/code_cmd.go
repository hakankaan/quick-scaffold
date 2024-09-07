package cmd

// example use of this command is 'quick-scaffold code -r ./ "user input to ask chat gpt" ./folder1/file1.go ./folder2/file2.go ./folder3'.
// -r flag is for rootPath. this is optional. default to ./
// the text inside quotes is for user requirement from chat gpt.
// and the next folder and file paths are for getting contents of those files.

import (
	"bufio"
	"context"

	"github.com/hakankaan/quick-scaffold/pkg/chatgpt"
	"github.com/hakankaan/quick-scaffold/pkg/config"
	"github.com/hakankaan/quick-scaffold/pkg/conversation"
	"github.com/hakankaan/quick-scaffold/pkg/fileutils"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var codeCmd = &cobra.Command{
	Use:   "code [USER_INPUT] [FILE_PATHS]",
	Short: "Generate code from user input and file paths",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rootPath, _ := cmd.Flags().GetString("rootPath")
		userInput := args[0]
		filePaths := args[1:]
		// use chatgpt.ChatWithGPT to get response from chat gpt.
		// use conversation.ParseResponse to parse the response.
		// if parsed response does not contain any file operation then just print the response.
		// and ask for user input to continue.
		// iterate same process until user input is "exit". or if response contains file operation.
		// if response contains file operation then iterate over file operations.

		// if file operation is "create" then create file with given path and content.
		// if file operation is "modify" then mofify content to file with given path.
		// if file operation is "delete" then delete file with given path.
		// if file operation is "Need_file_content" then get the content of file with given path. And add to conversation

		// if user input is "exit" then exit the program.

		apiKey, err := config.GetAPIKey()
		if err != nil {
			cmd.PrintErrln("Failed to get api key:", err)
			return
		}

		chatGptClient, err := chatgpt.New(apiKey)
		if err != nil {
			cmd.PrintErrln("Failed to create chat gpt client:", err)
			return
		}

		var currentConversation []openai.ChatCompletionMessage

		for {
			response, err := chatgpt.ChatWithGPT(context.Background(), chatGptClient, rootPath, currentConversation, userInput, filePaths)
			if err != nil {
				cmd.PrintErrln("Failed to get response from chat gpt:", err)
				return
			}

			fileOperations, err := conversation.ParseResponse(response.Content)
			if err != nil {
				cmd.PrintErrln("Failed to parse response:", err)
				return
			}

			if len(fileOperations) == 0 {
				cmd.Print(response.Content)
				cmd.Print("What is your answer? (type 'exit' to exit) ")

				scanner := bufio.NewScanner(reader)
				if scanner.Scan() {
					userInput = scanner.Text()
				}

				if userInput == "exit" {
					break
				}

				currentConversation = append(currentConversation, openai.ChatCompletionMessage{
					Role:    "user",
					Content: userInput,
				})
			} else {
				handleFileActions(cmd, fileOperations, currentConversation)
				break
			}

			if len(fileOperations) == 0 {
				cmd.Print(response.Content)
				cmd.Print("What is your anser? (type 'exit' to exit)")
			}

		}

	},
}

func init() {
	rootCmd.AddCommand(codeCmd)
	codeCmd.Flags().StringP("rootPath", "r", "./", "Root path for the project")
}

func handleFileActions(cmd *cobra.Command, fileOperations []conversation.FileOperation, currentConversation []openai.ChatCompletionMessage) {
	for _, fileOperation := range fileOperations {
		if fileOperation.Type == "create" {
			err := fileutils.CreateFileWithContent(fileOperation.FileName, fileOperation.Content)
			if err != nil {
				cmd.PrintErrln("Failed to create file:", err)
				return
			}
		}

		if fileOperation.Type == "modify" {
			err := fileutils.ModifyFileWithContent(fileOperation.FileName, fileOperation.Content)
			if err != nil {
				cmd.PrintErrln("Failed to modify file:", err)
				return
			}
		}

		if fileOperation.Type == "delete" {
			err := fileutils.DeleteFile(fileOperation.FileName)
			if err != nil {
				cmd.PrintErrln("Failed to delete file:", err)
				return
			}
		}

		if fileOperation.Type == "Need_file_content" {
			fileContent, err := fileutils.GetFileContent(fileOperation.FileName)
			if err != nil {
				cmd.PrintErrln("Failed to get file content:", err)
				return
			}
			currentConversation = append(currentConversation, openai.ChatCompletionMessage{
				Role:    "user",
				Content: fileContent,
			})
		}
	}
}
