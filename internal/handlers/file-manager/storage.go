package filemanager

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/amirhnajafiz/ghoster/pkg/crypto"
	"github.com/amirhnajafiz/ghoster/pkg/model"
)

func createProjectDir(id string) error {
	if err := os.Mkdir(fmt.Sprintf("%s/%s", storageDir, id), os.ModeDir); err != nil {
		return fmt.Errorf("failed to create project dir: %w", err)
	}

	return nil
}

func createProjectMetaFile(path, name string) error {
	stamp := time.Now()

	meta := model.FileMeta{
		Name:      name,
		CreatedAt: stamp,
		Hash:      crypto.MD5Hash(fmt.Sprint("%s-%t", name, stamp)),
	}

	bytes, err := json.Marshal(meta)
	if err != nil {
		return fmt.Errorf("failed to create file meta struct: %w", err)
	}

	file, err := os.Create(fmt.Sprintf("%s/data.meta.json", path))
	if err != nil {
		return fmt.Errorf("failed to create meta.json file: %w", err)
	}

	file.Write(bytes)
	file.Close()

	return nil
}
