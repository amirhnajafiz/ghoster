package agent

import "github.com/amirhnajafiz/ghoster/internal/agent/worker"

type Pool struct{}

func NewPool() *Pool {
	// create pool with internal channels
	pool := &Pool{}

	return pool
}

func (p *Pool) Throw(input string) error {
	return worker.Worker{}.Work(input)
}
