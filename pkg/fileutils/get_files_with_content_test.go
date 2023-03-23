package fileutils

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGetFilesWithContent(t *testing.T) {
	// Creating temporary files and folders
	tempFolder1 := createTempFolder(t, "test_folder_1")
	defer cleanupTempFolder(t, tempFolder1)

	tempFolder2 := createTempFolder(t, "test_folder_2")
	defer cleanupTempFolder(t, tempFolder1)

	file1Content := "This is file 1 content."
	file1 := createTempFile(t, tempFolder1, "file1.txt", file1Content)
	defer cleanupTempFile(t, file1)

	file2Content := "This is file 2 content."
	file2 := createTempFile(t, tempFolder1, "file2.txt", file2Content)
	defer cleanupTempFile(t, file2)

	file3Content := "This is file 3 content."
	file3 := createTempFile(t, tempFolder2, "file1.txt", file3Content)
	defer cleanupTempFile(t, file3)

	file4Content := "This is file 4 content."
	file4 := createTempFile(t, tempFolder2, "file2.txt", file4Content)
	defer cleanupTempFile(t, file4)

	// Call the GetFilesWithContent function
	filesWithContent, err := GetFilesWithContent(file1, tempFolder2, file2)
	if err != nil {
		t.Fatalf("GetFilesWithContent() error: %v", err)
	}

	// Check if the returned files and contents are correct
	if len(filesWithContent) != 4 {
		t.Fatalf("Expected 4 files, got: %d", len(filesWithContent))
	}

	for _, fileWithContent := range filesWithContent {
		if fileWithContent.Path == file1 {
			if fileWithContent.Content != file1Content {
				t.Errorf("Expected content for file1: %v, got: %v", file1Content, fileWithContent.Content)
			}
		} else if fileWithContent.Path == file2 {
			if fileWithContent.Content != file2Content {
				t.Errorf("Expected content for file2: %v, got: %v", file2Content, fileWithContent.Content)
			}
		} else if fileWithContent.Path == file3 {
			if fileWithContent.Content != file3Content {
				t.Errorf("Expected content for file3: %v, got: %v", file3Content, fileWithContent.Content)
			}
		} else if fileWithContent.Path == file4 {
			if fileWithContent.Content != file4Content {
				t.Errorf("Expected content for file4: %v, got: %v", file4Content, fileWithContent.Content)
			}
		} else {
			t.Errorf("Unexpected file path: %v", fileWithContent.Path)
		}
	}
}

func createTempFile(t *testing.T, folder string, fileName string, content string) string {
	tempFile, err := ioutil.TempFile(folder, fileName)
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}

	_, err = tempFile.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write content to temporary file: %v", err)
	}

	err = tempFile.Close()
	if err != nil {
		t.Fatalf("Failed to close temporary file: %v", err)
	}

	return tempFile.Name()
}

func createTempFolder(t *testing.T, folderName string) string {
	tempFolder, err := ioutil.TempDir("", folderName)
	if err != nil {
		t.Fatalf("Failed to create temporary folder: %v", err)
	}
	return tempFolder
}

func cleanupTempFile(t *testing.T, filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		t.Fatalf("Failed to clean up temporary file: %v", err)
	}
}

func cleanupTempFolder(t *testing.T, folderPath string) {
	err := os.RemoveAll(folderPath)
	if err != nil {
		t.Fatalf("Failed to clean up temporary folder: %v", err)
	}
}
