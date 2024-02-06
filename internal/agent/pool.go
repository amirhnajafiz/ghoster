package agent

import "github.com/amirhnajafiz/ghoster/internal/agent/worker"

type Pool struct {
	Channel chan string
}

func NewPool(number int) Pool {
	// create pool channels
	channel := make(chan string)

	// create workers
	for i := 0; i < number; i++ {
		w := worker.Worker{
			Channel: channel,
		}

		go w.Work()
	}

	return Pool{
		Channel: channel,
	}
}
