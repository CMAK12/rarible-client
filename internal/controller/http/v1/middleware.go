package http_v1

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nordew/go-errx"
	"go.uber.org/zap"
)

func loggerMiddleware(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		stop := time.Since(start)

		logger.Info(
			"HTTP Request",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("latency", stop),
		)

		return err
	}
}

func responseWrapper(handler func(c *fiber.Ctx) (map[string]any, error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		response, err := handler(c)
		if err != nil {
			return handleError(c, err)
		}

		return c.JSON(response)
	}
}

func handleError(c *fiber.Ctx, err error) error {
	switch {
	case errx.IsCode(err, errx.NotFound):
		return writeError(c, fiber.StatusNotFound, err)
	case errx.IsCode(err, errx.BadRequest):
		return writeError(c, fiber.StatusBadRequest, err)
	case errx.IsCode(err, errx.Internal):
		return writeError(c, fiber.StatusInternalServerError, err)
	case errx.IsCode(err, errx.Unauthorized):
		return writeError(c, fiber.StatusUnauthorized, err)
	case errx.IsCode(err, errx.Forbidden):
		return writeError(c, fiber.StatusForbidden, err)
	default:
		return writeError(c, fiber.StatusInternalServerError, fmt.Errorf("unexpected error: %v", err))
	}
}

func writeError(c *fiber.Ctx, statusCode int, err error) error {
	response := fiber.Map{
		"error": err.Error(),
	}

	return c.Status(statusCode).JSON(response)
}
