package agent

import "github.com/amirhnajafiz/ghoster/pkg/logger"

// Agent manages the workers by
// using a worker pool.
type Agent struct {
	workerPool *Pool
	logger     logger.Logger
}

func New(l logger.Logger, poolSize int) *Agent {
	pool := NewPool(poolSize)

	// listening on workers to manage their status
	go pool.listen()

	return &Agent{
		workerPool: pool,
		logger:     l,
	}
}

// NewWorker generates a new worker for
// clients.
func (a Agent) NewWorker() Worker {
	return a.workerPool.borrow()
}
