package middlewares

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Logger(c *fiber.Ctx) error {
	startTime := time.Now()
	duration := time.Since(startTime)

	logger := log.Info()
	if c.Response().StatusCode() != http.StatusOK {
		logger = log.Error().Bytes("body", c.Body())
	}

	logger.
		Str("protocol", "http").
		Str("method", c.Method()).
		Str("path", c.OriginalURL()).
		Int("status_code", c.Response().StatusCode()).
		Str("status_text", http.StatusText(c.Response().StatusCode())).
		Dur("duration", duration).
		Msg("received an HTTP request")

	return c.Next()
}
