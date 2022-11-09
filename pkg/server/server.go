package server

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/golang-jwt/jwt/v4"
	userHdl "github.com/solrac97gr/go-jwt-auth/internal/user/domain/ports"
	configurator "github.com/solrac97gr/go-jwt-auth/pkg/config/domain/ports"
	mdl "github.com/solrac97gr/go-jwt-auth/pkg/middleware/domain/ports"
)

type Server struct {
	mdl          mdl.MiddlewareHandlers
	user         userHdl.UserHandlers
	configurator configurator.ConfigApplication
}

func NewServer(mdl mdl.MiddlewareHandlers, user userHdl.UserHandlers, config configurator.ConfigApplication) *Server {
	return &Server{mdl: mdl, user: user, configurator: config}
}

func (s *Server) Run(port string) error {
	cfg, err := s.configurator.GetConfig()
	if err != nil {
		return err
	}
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/metrics", monitor.New(monitor.Config{Title: cfg.App.Name + " Metrics"}))
	app.Use(logger.New())
	app.Get("/docs/*", swagger.New())

	v1Public := app.Group("/api/v1")

	// User Endpoints
	v1Public.Post("/register", s.user.Register)
	v1Public.Post("/login", s.user.Login)

	v1Private := app.Group("/api/v1").Use(s.mdl.Authenticate())
	// Test Endpoint
	v1Private.Get("/secret", func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		email := claims["email"].(string)
		return c.SendString("Welcome 👋" + email)

	})

	err = app.Listen(":" + port)
	if err != nil {
		return err
	}
	return nil
}
