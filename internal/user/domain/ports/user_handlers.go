package ports

import "github.com/gofiber/fiber/v2"

type UserHandlers interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}
