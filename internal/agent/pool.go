package agent

import (
	"github.com/amirhnajafiz/ghoster/internal/agent/worker"
)

type Pool struct {
	pipe chan string
}

func NewPool(number int) *Pool {
	// create pool with internal channels
	pool := &Pool{
		pipe: make(chan string),
	}

	// create workers
	for i := 0; i < number; i++ {
		w := worker.Worker{
			Pipe: pool.pipe,
		}

		go w.Work()
	}

	return pool
}

func (p Pool) Add(msg string) {
	p.pipe <- msg
}
