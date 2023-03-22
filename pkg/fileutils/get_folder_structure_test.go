package fileutils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetFolderStructure(t *testing.T) {
	tempDir := t.TempDir()

	// Create a test folder structure
	err := os.Mkdir(filepath.Join(tempDir, "folder1"), 0755)
	if err != nil {
		t.Fatalf("Failed to create folder1: %v", err)
	}
	err = os.Mkdir(filepath.Join(tempDir, "folder2"), 0755)
	if err != nil {
		t.Fatalf("Failed to create folder2: %v", err)
	}
	err = os.Mkdir(filepath.Join(tempDir, "folder1", "subfolder"), 0755)
	if err != nil {
		t.Fatalf("Failed to create subfolder: %v", err)
	}
	err = os.WriteFile(filepath.Join(tempDir, "file1.txt"), []byte("file1 content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create file1.txt: %v", err)
	}
	err = os.WriteFile(filepath.Join(tempDir, "file2.txt"), []byte("file2 content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create file2.txt: %v", err)
	}
	err = os.WriteFile(filepath.Join(tempDir, "folder1", "subfolder", "file3.txt"), []byte("file3 content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create file3.txt: %v", err)
	}

	structure, err := GetFolderStructure(tempDir)
	if err != nil {
		t.Fatalf("GetFolderStructure() error: %v", err)
	}

	expectedStructure := []string{
		"file1.txt",
		"file2.txt",
		"folder1",
		"folder1/subfolder",
		"folder1/subfolder/file3.txt",
		"folder2",
	}

	if len(structure) != len(expectedStructure) {
		t.Fatalf("Expected %d items in the folder structure, got %d", len(expectedStructure), len(structure))
	}

	for i, path := range expectedStructure {
		if structure[i] != path {
			t.Errorf("Expected path: %s, got: %s", path, structure[i])
		}
	}
}
