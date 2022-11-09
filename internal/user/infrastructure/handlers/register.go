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
// @Param Body body models.RegisterRequest true "Register"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /register [post]
func (hdl *UserHdl) Register(ctx *fiber.Ctx) error {
	authReq := new(models.RegisterRequest)
	if err := ctx.BodyParser(authReq); err != nil {
		hdl.logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := hdl.app.Create(authReq)
	if err != nil {
		hdl.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
	})
}
