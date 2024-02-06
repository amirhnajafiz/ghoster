package agent

import "github.com/amirhnajafiz/ghoster/pkg/logger"

type Agent struct {
	workerPool *Pool
	logger     logger.Logger
}

func New(l logger.Logger, poolSize int) *Agent {
	pool := NewPool(poolSize)
	go pool.listen()

	return &Agent{
		workerPool: pool,
		logger:     l,
	}
}

func (a Agent) NewWorker() Worker {
	return a.workerPool.borrow()
}
