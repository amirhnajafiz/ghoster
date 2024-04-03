package cexe

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"golang.org/x/sync/semaphore"
)

// CExe is a component for executing ghoster projects (aka functions)
type CExe struct {
	semaphore *semaphore.Weighted
}

func New(pool int) *CExe {
	return &CExe{
		semaphore: semaphore.NewWeighted(int64(pool)),
	}
}

func (c *CExe) Execute(path string, args []string) ([]byte, time.Duration, error) {
	// get a resource to continue
	ctx := context.Background()
	c.semaphore.Acquire(ctx, 1)
	defer func() {
		c.semaphore.Release(1)
	}()

	// generate function command
	cmd := exec.Command("go", args...)
	cmd.Dir = path

	now := time.Now()

	// execute function
	bytes, err := cmd.Output()
	if err != nil {
		return nil, 0, fmt.Errorf("cexe failed to execute project: %v", err)
	}

	return bytes, time.Since(now), nil
}
