package worker

import (
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
