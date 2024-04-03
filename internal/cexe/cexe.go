package cexe

import (
	"fmt"
	"os/exec"
	"time"
)

func Execute(path string, args []string) ([]byte, time.Duration, error) {
	cmd := exec.Command("go", args...)
	cmd.Dir = path

	now := time.Now()

	bytes, err := cmd.Output()
	if err != nil {
		return nil, 0, fmt.Errorf("cexe failed to execute project: %v", err)
	}

	return bytes, time.Since(now), nil
}
