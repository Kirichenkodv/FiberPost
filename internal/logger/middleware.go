package logger

import (
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Middleware(log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {

		start := time.Now()

		err := c.Next()

		latency := time.Since(start)

		log.Info("http.request",
			"status", c.Response().StatusCode(),
			"method", c.Method(),
			"path", c.Path(),
			"latency", latency,
			"ip", c.IP(),
			"user_agent", c.Get(fiber.HeaderUserAgent),
		)

		return err
	}
}
