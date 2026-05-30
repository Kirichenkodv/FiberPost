package main

import (
	"github.com/Kirichenkodv/FiberPost/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Init()
	app := fiber.New()
	app.Listen(":8080")
}
