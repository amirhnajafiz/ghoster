package agent

import "github.com/amirhnajafiz/ghoster/internal/agent/worker"

type Pool struct {
}

func NewPool(number int) (chan string, chan bool) {
	// create pool channels
	channel := make(chan string)
	termination := make(chan bool)

	// create workers
	for i := 0; i < number; i++ {
		w := worker.Worker{
			Channel:   channel,
			Terminate: termination,
		}

		go w.Work()
	}

	return channel, termination
}
