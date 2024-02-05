package agent

import "github.com/amirhnajafiz/ghoster/internal/agent/worker"

func NewPool(number int) chan string {
	// create a new channel
	channel := make(chan string)

	// create workers
	for i := 0; i < number; i++ {
		w := worker.Worker{
			Channel: channel,
		}

		go w.Work()
	}

	return channel
}
