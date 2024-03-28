package config

type Config struct {
	HTTPPort         int
	MetricsPort      int
	MetricsNamespace string
	MetricsSubSystem string
}

func Load() Config {
	return Config{
		HTTPPort:         readIntFromEnv("HTTP_PORT"),
		MetricsPort:      readIntFromEnv("METRICS_PORT"),
		MetricsNamespace: readFromEnv("METRICS_NS"),
		MetricsSubSystem: readFromEnv("METRICS_SS"),
	}
}
