package filemanager

import (
	"fmt"
	"os"
)

func createProjectDir(id string) error {
	if err := os.Mkdir(fmt.Sprintf("%s/%s", storageDir, id), os.ModeDir); err != nil {
		return fmt.Errorf("failed to create project dir: %w", err)
	}

	return nil
}

func createProjectMetaFile(path, name string) error {

}
