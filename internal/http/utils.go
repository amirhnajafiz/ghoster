package http

import (
	"fmt"
	"os"
)

func listDirectoryItems(path string) ([]string, error) {
	items, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	folders := make([]string, len(items))
	for index, item := range items {
		folders[index] = item.Name()
	}

	return folders, err
}
