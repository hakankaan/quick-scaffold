package cmd

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
)

func TestHandleCodeCommand(t *testing.T) {

	testRootCmd := &cobra.Command{}
	testRootCmd.AddCommand(codeCmd)

	testCases := []struct {
		name      string
		rootPath  string
		userInput string
		filePaths []string
		expectErr bool
	}{
		{
			name:      "Example code command with root path and file paths",
			rootPath:  "./",
			userInput: "Create a simple Go function",
			filePaths: []string{"./folder1/file1.go", "./folder2/file2.go", "./folder3"},
			expectErr: false,
		},
		{
			name:      "Example code command without root path and with file paths",
			userInput: "Create a Go function that reads a file",
			filePaths: []string{"./folder1/file1.go", "./folder2/file2.go", "./folder3"},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := executeCommand(testRootCmd, []string{"code", tc.userInput, tc.rootPath, tc.filePaths[0], tc.filePaths[1], tc.filePaths[2]}, &bytes.Buffer{})

			if (err != nil) != tc.expectErr {
				t.Errorf("HandleCodeCommand() error = %v and expectErr = %v", err, tc.expectErr)
				return
			}
		})
	}
}
