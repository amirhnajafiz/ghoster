package agent

import (
	"context"
	"time"

	"github.com/amirhnajafiz/ghoster/internal/agent/worker"

	"golang.org/x/sync/semaphore"
)

// Pool manages the workers.
type Pool struct {
	semaphore *semaphore.Weighted
	pipe      chan int
}

func NewPool(limit int) *Pool {
	// create pool with internal channels and a semaphore
	pool := &Pool{
		semaphore: semaphore.NewWeighted(int64(limit)),
		pipe:      make(chan int),
	}

	return pool
}

// listen on workers status.
func (p *Pool) listen() {
	for {
		// wait for a worker to finish
		<-p.pipe

		// release semaphore
		p.semaphore.Release(1)
	}
}

// borrow creates a new worker, but before doing that
// it trys to acquire the pool semaphore.
func (p *Pool) borrow(timeout int) (*worker.Worker, error) {
	// create a new context with 10 seconds limit
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	for {
		// wait for a free worker
		if err := p.semaphore.Acquire(ctx, 1); err != nil {
			return nil, err
		}

		// create a new worker
		w := worker.New(p.pipe)
		go w.Work()

		return w, nil
	}
}
