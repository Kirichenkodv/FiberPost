package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Debug("No .env file")
	} else {
		log.Debug(".env file loaded")
	}
}
