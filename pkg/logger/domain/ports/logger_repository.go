package ports

import "github.com/solrac97gr/go-jwt-auth/pkg/logger/domain/models"

type LoggerRepository interface {
	Save(log *models.Log) error
}
