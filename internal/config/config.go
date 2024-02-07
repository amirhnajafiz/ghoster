package config

import (
	"os"
	"strconv"

	"github.com/amirhnajafiz/ghoster/internal/config/agent"
	"github.com/amirhnajafiz/ghoster/internal/config/http"
	"github.com/amirhnajafiz/ghoster/internal/config/logger"
	"github.com/amirhnajafiz/ghoster/internal/storage/mongodb"
)

type Config struct {
	HTTP    http.Config
	Agent   agent.Config
	MongoDB mongodb.Config
	Logger  logger.Config
}

func Load() Config {
	instance := Default()

	instance.HTTP.Port, _ = strconv.Atoi(os.Getenv("http__port"))

	instance.Agent.PoolSize, _ = strconv.Atoi(os.Getenv("agent__pool_size"))
	instance.Agent.Timeout, _ = strconv.Atoi(os.Getenv("agent__timeout"))

	instance.Logger.Level, _ = strconv.Atoi(os.Getenv("logger__level"))

	instance.MongoDB.URI = os.Getenv("mongodb__uri")
	instance.MongoDB.Database = os.Getenv("mongodb__database")
	instance.MongoDB.Collection = os.Getenv("mongodb__collection")

	return instance
}
