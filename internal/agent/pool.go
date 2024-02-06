package agent

import (
	"github.com/amirhnajafiz/ghoster/internal/agent/worker"
)

type Pool struct {
	pipe  chan string
	stdin chan string
}

func NewPool(number int) *Pool {
	// create pool with internal channels
	pool := &Pool{
		pipe:  make(chan string),
		stdin: make(chan string),
	}

	// create workers
	for i := 0; i < number; i++ {
		w := worker.Worker{
			Pipe:   pool.pipe,
			Stdout: pool.stdin,
		}

		go w.Work()
	}

	return pool
}
