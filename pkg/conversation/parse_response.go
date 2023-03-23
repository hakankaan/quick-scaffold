package conversation

import (
	"errors"
	"strings"
)

type FileOperation struct {
	Type     string
	FileName string
	Content  string
}

func ParseResponse(response string) ([]FileOperation, error) {
	var fileOperations []FileOperation
	lines := strings.Split(response, "\n")

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if strings.HasPrefix(line, "create_file:") || strings.HasPrefix(line, "modify_file:") || strings.HasPrefix(line, "delete_file:") || strings.HasPrefix(line, "Need_content_of_file:") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) != 2 {
				return nil, errors.New("invalid file operation format")
			}

			opType := parts[0]
			fileName := strings.TrimSpace(parts[1])

			content := ""

			if opType == "create_file" || opType == "modify_file" {
				if i+2 >= len(lines) || !strings.HasPrefix(lines[i+1], "```") {
					return nil, errors.New("invalid file operation format")
				}

				if i+2 >= len(lines) || !strings.HasPrefix(lines[i+1], "```") {
					return nil, errors.New("invalid file operation format")
				}

				contentStartIndex := i + 2
				contentEndIndex := contentStartIndex
				for j := contentStartIndex; j < len(lines); j++ {
					if strings.HasPrefix(lines[j], "```") {
						contentEndIndex = j
						break
					}
				}

				if contentEndIndex == contentStartIndex {
					return nil, errors.New("invalid file operation format")
				}

				content = strings.Join(lines[contentStartIndex:contentEndIndex], "\n")
			}

			fileOperations = append(fileOperations, FileOperation{
				Type:     opType,
				FileName: fileName,
				Content:  content,
			})

		}
	}

	return fileOperations, nil
}
