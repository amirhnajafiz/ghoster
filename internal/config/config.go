package config

import "github.com/joho/godotenv"

type Config struct {
	HTTPPort         int
	MetricsPort      int
	MetricsNamespace string
	MetricsSubSystem string
	PoolSize         int
}

// Load env variables to a Config instance
func Load() Config {
	godotenv.Load()

	return Config{
		HTTPPort:         readIntFromEnv("HTTP_PORT"),
		MetricsPort:      readIntFromEnv("METRICS_PORT"),
		MetricsNamespace: readFromEnv("METRICS_NS"),
		MetricsSubSystem: readFromEnv("METRICS_SS"),
		PoolSize:         readIntFromEnv("POOL_SIZE"),
	}
}
