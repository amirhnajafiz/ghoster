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
	PoolSize   int
}

func (a Agent) Done() {

}

func (a Agent) Listen() {
	counter := 10
	a.WorkerPool = NewPool(counter)

	for {
		path := <-a.Channel

		a.WorkerPool.Channel <- path
	}
}
