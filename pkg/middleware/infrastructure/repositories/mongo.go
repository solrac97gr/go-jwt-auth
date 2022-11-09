package repositories

import (
	"github.com/solrac97gr/go-jwt-auth/pkg/config/domain/ports"
	logger "github.com/solrac97gr/go-jwt-auth/pkg/logger/domain/ports"
	mdl "github.com/solrac97gr/go-jwt-auth/pkg/middleware/domain/ports"
)

type MongoMdlRepository struct {
	config ports.ConfigApplication
	logger logger.LoggerApplication
}

var _ mdl.MiddlewareRepository = (*MongoMdlRepository)(nil)

func NewMongoMdlRepository(config ports.ConfigApplication, logger logger.LoggerApplication) *MongoMdlRepository {
	return &MongoMdlRepository{
		config: config,
		logger: logger,
	}
}

func (m *MongoMdlRepository) Authenticate(token string) (err error) {
	//TODO implement me
	panic("implement me")
}
