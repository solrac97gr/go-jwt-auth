package ports

import (
	"github.com/gofiber/fiber/v2"
)

type MiddlewareHandlers interface {
	Authenticate(ctx *fiber.Ctx) error
}
