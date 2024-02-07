package agent

import "fmt"

// Agent manages the workers by
// using a worker pool.
type Agent struct {
	workerPool *Pool
	timeout    int
}

func New(poolSize int, timeout int) *Agent {
	pool := NewPool(poolSize)

	// listening on workers to manage their status
	go pool.listen()

	return &Agent{
		workerPool: pool,
		timeout:    timeout,
	}
}

// NewWorker generates a new worker for
// clients.
func (a Agent) NewWorker() (Worker, error) {
	w, err := a.workerPool.borrow(a.timeout)
	if err != nil {
		return nil, fmt.Errorf("failed to create worker: %w", err)
	}

	return w, nil
}
