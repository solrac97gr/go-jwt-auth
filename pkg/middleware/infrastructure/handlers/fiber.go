package handlers

import (
	"github.com/gofiber/fiber/v2"
	logger "github.com/solrac97gr/go-jwt-auth/pkg/logger/domain/ports"
	"github.com/solrac97gr/go-jwt-auth/pkg/middleware/domain/ports"
)

type FiberMdlHandlers struct {
	service ports.MiddlewareApplication
	logger  logger.LoggerApplication
}

var _ ports.MiddlewareHandlers = (*FiberMdlHandlers)(nil)

func NewFiberMdlHandlers(service ports.MiddlewareApplication, logger logger.LoggerApplication) *FiberMdlHandlers {
	return &FiberMdlHandlers{
		service: service,
		logger:  logger,
	}
}

func (f *FiberMdlHandlers) Authenticate(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
