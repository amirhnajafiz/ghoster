package agent

import "github.com/amirhnajafiz/ghoster/pkg/logger"

type Agent struct {
	workerPool *Pool
	logger     logger.Logger

	stdin  chan interface{}
	stdout chan interface{}
}

func New(logger logger.Logger) *Agent {
	return &Agent{
		workerPool: NewPool(),
		stdin:      make(chan interface{}),
		stdout:     make(chan interface{}),
		logger:     logger,
	}
}

func (a Agent) GetStdin() chan interface{} {
	return a.stdin
}

func (a Agent) GetStdout() chan interface{} {
	return a.stdout
}

func (a Agent) Listen() {
	for {
		path := <-a.stdin

		if err := a.workerPool.Throw(path.(string)); err != nil {
			a.stdout <- err.Error()
		}
	}
}
