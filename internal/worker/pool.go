package worker

import (
	"context"

	"golang.org/x/sync/semaphore"
)

type Pool struct {
	Semaphore *semaphore.Weighted
}

func NewPool(limit int) Pool {
	instance := Pool{}

	instance.Semaphore = semaphore.NewWeighted(int64(limit))

	return instance
}

func (p Pool) Pull() {
	ctx := context.Background()

	p.Semaphore.Acquire(ctx, 1)
}

func (p Pool) Free() {
	p.Semaphore.Release(1)
}
