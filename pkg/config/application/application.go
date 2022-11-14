package application

import (
	"github.com/solrac97gr/go-jwt-auth/pkg/config/domain/models"
	"github.com/solrac97gr/go-jwt-auth/pkg/config/domain/ports"
	logger "github.com/solrac97gr/go-jwt-auth/pkg/logger/domain/ports"
	validator "github.com/solrac97gr/go-jwt-auth/pkg/validator/domain/ports"
)

type ConfigService struct {
	repository    ports.ConfigRepository
	configuration *models.Config
	validator     validator.ValidatorApplication
	logger        logger.LoggerApplication
}

var _ ports.ConfigApplication = (*ConfigService)(nil)

func NewConfigService(
	repository ports.ConfigRepository,
	val validator.ValidatorApplication,
	logger logger.LoggerApplication,
) *ConfigService {
	return &ConfigService{
		repository: repository,
		validator:  val,
		logger:     logger,
	}
}
