package config

import (
	"github.com/amirhnajafiz/ghoster/internal/config/agent"
	"github.com/amirhnajafiz/ghoster/internal/config/http"
	"github.com/amirhnajafiz/ghoster/internal/config/logger"
	"github.com/amirhnajafiz/ghoster/internal/storage/mongodb"
)

func Default() Config {
	return Config{
		HTTP: http.Config{
			Port: 5000,
		},
		Agent: agent.Config{
			PoolSize: 10,
		},
		MongoDB: mongodb.Config{
			URI:        "",
			Database:   "",
			Collection: "",
		},
		Logger: logger.Config{
			Level: 1,
		},
	}
}
