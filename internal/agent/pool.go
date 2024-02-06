package agent

import (
	"time"

	"github.com/amirhnajafiz/ghoster/internal/agent/worker"
)

type Pool struct {
	limit int
	inuse int
	pipe  chan int
}

func NewPool(limit int) *Pool {
	// create pool with internal channels
	pool := &Pool{
		limit: limit,
		inuse: 0,
		pipe:  make(chan int),
	}

	return pool
}

func (p *Pool) listen() {
	for {
		<-p.pipe
		p.inuse--
	}
}

func (p *Pool) borrow() *worker.Worker {
	for {
		if p.limit == p.inuse {
			time.Sleep(1 * time.Second)
			continue
		}

		p.inuse++

		w := worker.New(p.pipe)
		go w.Work()

		return w
	}
}
