package application

import (
	"encoding/json"
	"github.com/solrac97gr/go-jwt-auth/pkg/config/domain/models"
	"github.com/solrac97gr/go-jwt-auth/pkg/config/domain/ports"
	validator "github.com/solrac97gr/go-jwt-auth/pkg/validator/domain/ports"
	"log"
)

type ConfigService struct {
	repository    ports.ConfigRepository
	configuration *models.Config
	validator     validator.ValidatorApplication
}

func NewConfigService(repository ports.ConfigRepository) *ConfigService {
	return &ConfigService{
		repository: repository,
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
		log.Fatal("parsing config file", err.Error())
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
