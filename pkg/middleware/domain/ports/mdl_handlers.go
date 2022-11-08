package ports

import (
	"github.com/gofiber/fiber/v2"
)

type MiddlewareHandlers interface {
	Login(ctx *fiber.Ctx) error
	Authenticate(ctx *fiber.Ctx) error
}
