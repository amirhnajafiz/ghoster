package gc

import (
	"fmt"
	"os"
	"strings"
)

func deleteFilesWithStartToken(path, token string) (int, error) {
	items, err := os.ReadDir(path)
	if err != nil {
		return 0, fmt.Errorf("failed to read directory: %w", err)
	}

	count := 0

	for _, item := range items {
		if strings.HasPrefix(item.Name(), token) {
			_ = os.RemoveAll(fmt.Sprintf("%s/%s", path, item.Name()))
			count++
		}
	}

	return count, nil
}
