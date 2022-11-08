package ports

import (
	"github.com/solrac97gr/go-jwt-auth/pkg/config/domain/models"
	"os"
)

type ConfigApplication interface {
	Config() error
	GetConfig() (*models.Config, error)
}

type ConfigRepository interface {
	GetConfigFile() (*os.File, error)
}
