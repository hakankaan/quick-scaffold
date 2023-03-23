package fileutils

import "os"

// GetFileContent returns the content of given file with relative path
func GetFileContent(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
