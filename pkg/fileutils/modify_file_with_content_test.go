package fileutils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestModifyFileWithContent(t *testing.T) {
	testCases := []struct {
		name        string
		filePath    string
		content     string
		expectedErr bool
	}{
		{
			name:        "Modify a file with new content",
			filePath:    "test_files/test_file.txt",
			content:     "New content",
			expectedErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a file with initial content
			err := CreateFileWithContent(tc.filePath, "Initial content")
			if err != nil {
				t.Fatalf("Failed to create a file for the test: %v", err)
			}

			err = ModifyFileWithContent(tc.filePath, tc.content)
			if (err != nil) != tc.expectedErr {
				t.Errorf("ModifyFileWithContent() error = %v, expectedErr %v", err, tc.expectedErr)
			}

			if !tc.expectedErr {
				data, err := ioutil.ReadFile(tc.filePath)
				if err != nil {
					t.Errorf("Failed to read the modified file: %v", err)
				}

				if string(data) != tc.content {
					t.Errorf("Expected content %q, got %q", tc.content, string(data))
				}

				// Clean up
				err = os.Remove(tc.filePath)
				if err != nil {
					t.Errorf("Failed to remove the test file: %v", err)
				}

				err = os.Remove(filepath.Dir(tc.filePath))
				if err != nil {
					t.Errorf("Failed to remove the test directory: %v", err)
				}
			}
		})
	}
}
