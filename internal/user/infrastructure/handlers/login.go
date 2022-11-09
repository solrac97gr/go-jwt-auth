package handlers

import "github.com/gofiber/fiber/v2"

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
func (u *UserHdl) Login(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
