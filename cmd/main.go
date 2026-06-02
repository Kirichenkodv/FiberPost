package main

import (
	"github.com/Kirichenkodv/FiberPost/config"
	"github.com/Kirichenkodv/FiberPost/internal/logger"
	"github.com/gofiber/fiber/v2"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {

	cfg := config.Load()
	log := logger.New(cfg.LogFormat, cfg.LogLevel)

	app := fiber.New()
	app.Use(slogfiber.New(log))

	app.Listen(":8080")
}
