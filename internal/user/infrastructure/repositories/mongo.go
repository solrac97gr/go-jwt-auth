package repositories

import (
	"github.com/solrac97gr/go-jwt-auth/internal/user/domain/models"
	"github.com/solrac97gr/go-jwt-auth/internal/user/domain/ports"
	config "github.com/solrac97gr/go-jwt-auth/pkg/config/domain/ports"
	logger "github.com/solrac97gr/go-jwt-auth/pkg/logger/domain/ports"
)

type UserMongoDB struct {
	config config.ConfigApplication
	logger logger.LoggerApplication
}

var _ ports.UserRepository = (*UserMongoDB)(nil)

func NewUserMongoDB(config config.ConfigApplication, logger logger.LoggerApplication) *UserMongoDB {
	return &UserMongoDB{config: config, logger: logger}
}

func (u *UserMongoDB) Save(user *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserMongoDB) FindByCredentials(credentials *models.AuthRequest) error {
	//TODO implement me
	panic("implement me")
}
