package fileutils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetFileContent(t *testing.T) {
	testCases := []struct {
		name            string
		filePath        string
		expectedContent string
		expectedErr     bool
	}{
		{
			name:            "Get file content",
			filePath:        "test_files/test_file.txt",
			expectedContent: "This is a test file",
			expectedErr:     false,
		},
		{
			name:            "Try to get content of a non-existent file",
			filePath:        "test_files/non_existent_file.txt",
			expectedContent: "",
			expectedErr:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// create the file
			if !tc.expectedErr {
				err := CreateFileWithContent(tc.filePath, tc.expectedContent)
				if err != nil {
					t.Fatalf("Failed to create a file for the test: %v", err)
				}
			}

			content, err := GetFileContent(tc.filePath)
			if (err != nil) != tc.expectedErr {
				t.Errorf("GetFileContent() error = %v, expectedErr %v", err, tc.expectedErr)
			}

			if content != tc.expectedContent {
				t.Errorf("GetFileContent() content = %v, expectedContent %v", content, tc.expectedContent)
			}

			// check if file exists before trying to delete
			_, err = os.Stat(tc.filePath)
			if !os.IsNotExist(err) {

				// Clean up. delete the file
				err = DeleteFile(tc.filePath)
				if err != nil {
					t.Errorf("Failed to delete the file: %v", err)
				}

				// Clean up. delete the directory
				err = os.Remove(filepath.Dir(tc.filePath))
				if err != nil {
					t.Errorf("Failed to delete the directory: %v", err)
				}
			}

		})
	}
}
