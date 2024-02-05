package agent

import (
	"github.com/amirhnajafiz/ghoster/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
)

type Agent struct {
	DB         *mongo.Database
	Logger     logger.Logger
	Collection string
	Channel    chan string
	Pool       chan string
}

func (a Agent) Listen() {
	a.Pool = NewPool(10)

	for {
		path := <-a.Channel

		a.Pool <- path
	}
}
