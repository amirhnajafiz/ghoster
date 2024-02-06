package agent

import "github.com/amirhnajafiz/ghoster/internal/agent/worker"

type Pool struct {
	Channel     chan string
	Termination chan bool
	Pipe        chan int
}

func NewPool(number int) Pool {
	// create pool channels
	channel := make(chan string)
	termination := make(chan bool)
	pipe := make(chan int)

	// create workers
	for i := 0; i < number; i++ {
		w := worker.Worker{
			Channel:   channel,
			Terminate: termination,
			Pipe:      pipe,
		}

		go w.Work()
	}

	return Pool{
		Channel:     channel,
		Termination: termination,
		Pipe:        pipe,
	}
}
