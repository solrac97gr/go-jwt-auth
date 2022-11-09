package handlers

import (
	"github.com/solrac97gr/go-jwt-auth/internal/user/domain/ports"
	logger "github.com/solrac97gr/go-jwt-auth/pkg/logger/domain/ports"
)

type UserHdl struct {
	app    ports.UserApplication
	logger logger.LoggerApplication
}

var _ ports.UserHandlers = (*UserHdl)(nil)

func NewUserHdl(app ports.UserApplication, logger logger.LoggerApplication) *UserHdl {
	return &UserHdl{app: app, logger: logger}
}
