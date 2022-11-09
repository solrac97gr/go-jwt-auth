package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	user "github.com/solrac97gr/go-jwt-auth/internal/user/domain/ports"
	mdl "github.com/solrac97gr/go-jwt-auth/pkg/middleware/domain/ports"
)

type Server struct {
	mdl  mdl.MiddlewareHandlers
	user user.UserHandlers
}

func NewServer(mdl mdl.MiddlewareHandlers, user user.UserHandlers) *Server {
	return &Server{mdl: mdl, user: user}
}

func (s *Server) Run(port string) error {
	app := fiber.New()
	app.Get("/swagger/*", swagger.New())

	v1Public := app.Group("/api/v1")
	v1Private := app.Group("/api/v1").Use(s.mdl.Authenticate)

	// User Endpoints
	v1Public.Post("/register", s.user.Register)
	v1Public.Post("/login", s.user.Login)

	// Test Endpoint
	v1Private.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")

	})

	err := app.Listen(":" + port)
	if err != nil {
		return err
	}
	return nil
}
