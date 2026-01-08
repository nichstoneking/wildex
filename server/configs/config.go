package configs

import (
	"os"
)

type Config struct {
	ServerPort  string
	Environment string
}

func LoadConfig() *Config {
	config := &Config{
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}
	return config
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
