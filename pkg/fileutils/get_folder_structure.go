package fileutils

import (
	"os"
	"path/filepath"
)

func GetFolderStructure(root string) ([]string, error) {
	var paths []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the root folder
		if path == root {
			return nil
		}

		// Get the relative path
		relativePath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		paths = append(paths, relativePath)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return paths, nil
}
