package config

import (
	"github.com/amirhnajafiz/ghoster/internal/config/agent"
	"github.com/amirhnajafiz/ghoster/internal/config/http"
	"github.com/amirhnajafiz/ghoster/internal/storage/mongodb"
)

type Config struct {
	HTTP    http.Config    `koanf:"http"`
	Agent   agent.Config   `koanf:"agent"`
	MongoDB mongodb.Config `koanf:"db"`
}
