package main

import (
	"github.com/amirhnajafiz/ghoster/internal/agent"
	"github.com/amirhnajafiz/ghoster/internal/config"
	"github.com/amirhnajafiz/ghoster/internal/handler/http"
	"github.com/amirhnajafiz/ghoster/internal/storage/mongodb"
	"github.com/amirhnajafiz/ghoster/pkg/logger"
)

func main() {
	// load configs
	cfg := config.Load()

	// create a new logger
	log := logger.New(logger.Level(cfg.Logger.Level))

	// open mongodb connection
	db, err := mongodb.NewConnection(cfg.MongoDB)
	if err != nil {
		panic(err)
	}

	// create a new agent
	a := agent.New(cfg.Agent.PoolSize)

	// create a new handler
	h := http.HTTP{
		Agent:      a,
		DB:         db,
		Logger:     log.Trace("http"),
		Collection: cfg.MongoDB.Collection,
	}

	// register http handler
	h.Register(cfg.HTTP.Port)
}
