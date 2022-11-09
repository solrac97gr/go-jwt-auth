package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/solrac97gr/go-jwt-auth/internal/user/domain/models"
)

// Register godoc
// @Summary Register a new user
// @Description Register a new user
// @Accept json
// @Produce json
// @Tags User
// @Param Body body models.AuthRequest true "Register"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /register [post]
func (u *UserHdl) Register(ctx *fiber.Ctx) error {
	authReq := models.AuthRequest{}
	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Hello, World 👋!",
		"body":    authReq,
	})
}
