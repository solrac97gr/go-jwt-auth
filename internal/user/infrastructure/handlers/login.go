package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/solrac97gr/go-jwt-auth/internal/user/domain/models"
)

// Login godoc
// @Summary Login a user
// @Description Login a user
// @Accept json
// @Produce json
// @Tags User
// @Param Body body models.AuthRequest true "Register"
// @Success 200 {object} models.AuthResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /login [post]
func (hdl *UserHdl) Login(ctx *fiber.Ctx) error {
	reqBody := new(models.AuthRequest)
	if err := ctx.BodyParser(reqBody); err != nil {
		hdl.logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	authResponse, err := hdl.app.Login(reqBody)
	if err != nil {
		hdl.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(authResponse)
}
