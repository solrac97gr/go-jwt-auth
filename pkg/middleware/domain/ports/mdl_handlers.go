package ports

import (
	"github.com/gofiber/fiber/v2"
)

type MiddlewareHandlers interface {
	Authenticate() fiber.Handler
}
