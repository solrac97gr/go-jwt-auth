package application

import (
	"encoding/json"
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

func NewConfigService(repository ports.ConfigRepository, val validator.ValidatorApplication, logger logger.LoggerApplication) *ConfigService {
	return &ConfigService{
		repository: repository,
		validator:  val,
		logger:     logger,
	}
}

func (c *ConfigService) Config() error {
	configuration := new(models.Config)
	file, err := c.repository.GetConfigFile()
	if err != nil {
		return err
	}
	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&configuration); err != nil {
		return err
	}
	c.configuration = configuration
	return nil
}

func (c *ConfigService) GetConfig() (*models.Config, error) {
	if c.configuration == nil {
		if err := c.Config(); err != nil {
			return nil, err
		}
	}
	if err := c.validator.Struct(c.configuration); err != nil {
		return c.configuration, nil
	}
	return c.configuration, nil
}
