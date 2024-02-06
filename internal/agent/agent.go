package agent

import (
	"github.com/amirhnajafiz/ghoster/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
)

type Agent struct {
	WorkerPool Pool
	DB         *mongo.Database
	Logger     logger.Logger
	Channel    chan string
	Collection string
}

func (a Agent) Listen() {
	counter := 10
	a.WorkerPool = NewPool(counter)

	for {
		select {
		case path := <-a.Channel:
			a.WorkerPool.Channel <- path
		case <-a.WorkerPool.Pipe:
			counter--
		}
	}
}
