package fileutils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDeleteFile(t *testing.T) {
	testCases := []struct {
		name        string
		filePath    string
		expectedErr bool
	}{
		{
			name:        "Delete an existing file",
			filePath:    "test_files/test_file.txt",
			expectedErr: false,
		},
		{
			name:        "Try to delete a non-existent file",
			filePath:    "test_files/non_existent_file.txt",
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if !tc.expectedErr {
				err := CreateFileWithContent(tc.filePath, "File to be deleted")
				if err != nil {
					t.Fatalf("Failed to create a file for the test: %v", err)
				}
			}

			err := DeleteFile(tc.filePath)
			if (err != nil) != tc.expectedErr {
				t.Errorf("DeleteFile() error = %v, expectedErr %v", err, tc.expectedErr)
			}

			if !tc.expectedErr {
				// Clean up. check if file exists
				file, err := os.Stat(tc.filePath)
				if !os.IsNotExist(err) {
					t.Errorf("Failed to delete the file: %v", err)
				}

				// Clean up. delete the file if still exists
				if file != nil {
					err = os.Remove(tc.filePath)
					if err != nil {
						t.Errorf("Failed to delete the file: %v", err)
					}
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
