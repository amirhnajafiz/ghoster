package config

type Config struct {
	HTTPPort int
}

func Load() Config {
	return Config{
		HTTPPort: readIntFromEnv("HTTP_PORT"),
	}
}
