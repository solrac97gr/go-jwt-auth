package application

import (
	"github.com/solrac97gr/go-jwt-auth/internal/user/domain/ports"
	configurator "github.com/solrac97gr/go-jwt-auth/pkg/config/domain/ports"
	logger "github.com/solrac97gr/go-jwt-auth/pkg/logger/domain/ports"
	val "github.com/solrac97gr/go-jwt-auth/pkg/validator/domain/ports"
)

type UserApp struct {
	repo         ports.UserRepository
	validator    val.ValidatorApplication
	logger       logger.LoggerApplication
	configurator configurator.ConfigApplication
}

var _ ports.UserApplication = (*UserApp)(nil)

func NewUserApp(
	repo ports.UserRepository,
	validator val.ValidatorApplication,
	logger logger.LoggerApplication,
	configurator configurator.ConfigApplication,
) *UserApp {
	return &UserApp{
		repo:         repo,
		validator:    validator,
		logger:       logger,
		configurator: configurator,
	}
}
