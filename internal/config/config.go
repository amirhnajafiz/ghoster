package config

import (
	"github.com/amirhnajafiz/ghoster/internal/config/agent"
	"github.com/amirhnajafiz/ghoster/internal/config/http"
	"github.com/amirhnajafiz/ghoster/internal/storage/mongodb"
	"github.com/amirhnajafiz/ghoster/internal/storage/s3"
)

type Config struct {
	HTTP    http.Config    `koanf:"http"`
	Agent   agent.Config   `koanf:"agent"`
	MongoDB mongodb.Config `koanf:"db"`
	S3      s3.Config      `koanf:"s3"`
}
