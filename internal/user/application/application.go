package application

import (
	"github.com/solrac97gr/go-jwt-auth/internal/user/domain/models"
	"github.com/solrac97gr/go-jwt-auth/internal/user/domain/ports"
	logger "github.com/solrac97gr/go-jwt-auth/pkg/logger/domain/ports"
	val "github.com/solrac97gr/go-jwt-auth/pkg/validator/domain/ports"
)

type UserApp struct {
	repo      ports.UserRepository
	validator val.ValidatorApplication
	logger    logger.LoggerApplication
}

var _ ports.UserApplication = (*UserApp)(nil)

func NewUserApp(repo ports.UserRepository, validator val.ValidatorApplication, logger logger.LoggerApplication) *UserApp {
	return &UserApp{repo: repo, validator: validator, logger: logger}
}

func (u *UserApp) Create(user *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserApp) Login(user *models.User) (authToken *models.AuthResponse, err error) {
	//TODO implement me
	panic("implement me")
}
