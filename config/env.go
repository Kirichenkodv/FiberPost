package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	LogFormat string
	LogLevel  string
}

func Load() Config {
	godotenv.Load()
	logFormat := os.Getenv("LOG_FORMAT")
	logLevel := os.Getenv("LOG_LEVEL")

	return Config{
		LogFormat: logFormat,
		LogLevel:  logLevel,
	}
}
