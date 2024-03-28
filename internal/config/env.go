package config

import (
	"os"
	"strconv"
)

func readFromEnv(key string) string {
	return os.Getenv(key)
}

func readIntFromEnv(key string) int {
	tmp, _ := strconv.Atoi(readFromEnv(key))

	return tmp
}
