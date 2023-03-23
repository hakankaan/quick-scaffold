package fileutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// FileContent represents a file with its content
type FileContent struct {
	Path    string
	Content string
}

func GetFilesWithContent(paths ...string) ([]FileContent, error) {
	var filesWithContent []FileContent

	for _, path := range paths {
		fileInfo, err := os.Stat(path)
		if err != nil {
			return nil, fmt.Errorf("os.Stat() error: %v", err)
		}

		if fileInfo.IsDir() {
			err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() {
					content, err := ioutil.ReadFile(filePath)
					if err != nil {
						return fmt.Errorf("ioutil.ReadFile() error: %v", err)
					}
					filesWithContent = append(filesWithContent, FileContent{
						Path:    filePath,
						Content: string(content),
					})
				}
				return nil
			})
			if err != nil {
				return nil, fmt.Errorf("filepath.Walk() error: %v", err)
			}
		} else {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return nil, err
			}
			filesWithContent = append(filesWithContent, FileContent{
				Path:    path,
				Content: string(content),
			})
		}
	}

	return filesWithContent, nil
}
