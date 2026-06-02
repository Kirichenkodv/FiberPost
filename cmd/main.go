package main

import (
	"github.com/Kirichenkodv/FiberPost/config"
	"github.com/Kirichenkodv/FiberPost/internal/logger"
	"github.com/gofiber/fiber/v2"
)

func main() {

	cfg := config.Load()
	log := logger.New(cfg.LogFormat, cfg.LogLevel)

	app := fiber.New()
	app.Use(logger.Middleware(log))
	app.Listen(":8080")
}
