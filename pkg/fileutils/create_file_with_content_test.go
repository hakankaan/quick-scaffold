package fileutils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateFileWithContent(t *testing.T) {
	testCases := []struct {
		name        string
		filePath    string
		content     string
		expectedErr bool
	}{
		{
			name:        "Create a file with content",
			filePath:    "test_files/test_file.txt",
			content:     "Hello, world!",
			expectedErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := CreateFileWithContent(tc.filePath, tc.content)
			if (err != nil) != tc.expectedErr {
				t.Errorf("CreateFileWithContent() error = %v, expectedErr %v", err, tc.expectedErr)
			}

			if !tc.expectedErr {
				file, err := os.Open(tc.filePath)
				if err != nil {
					t.Errorf("Failed to open created file: %v", err)
				}

				buf := make([]byte, len(tc.content))
				_, err = file.Read(buf)
				if err != nil {
					t.Errorf("Failed to read created file: %v", err)
				}

				if string(buf) != tc.content {
					t.Errorf("Content of created file is not correct: %v", string(buf))
				}

				file.Close()

				err = os.RemoveAll(filepath.Dir(tc.filePath))
				if err != nil {
					t.Errorf("Failed to remove test files: %v", err)
				}
			}
		})
	}
}
