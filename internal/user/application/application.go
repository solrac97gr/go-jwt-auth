package application

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/solrac97gr/go-jwt-auth/internal/user/domain/models"
	"github.com/solrac97gr/go-jwt-auth/internal/user/domain/ports"
	configurator "github.com/solrac97gr/go-jwt-auth/pkg/config/domain/ports"
	customErr "github.com/solrac97gr/go-jwt-auth/pkg/custom-errors"
	logger "github.com/solrac97gr/go-jwt-auth/pkg/logger/domain/ports"
	val "github.com/solrac97gr/go-jwt-auth/pkg/validator/domain/ports"
	"time"
)

type UserApp struct {
	repo         ports.UserRepository
	validator    val.ValidatorApplication
	logger       logger.LoggerApplication
	configurator configurator.ConfigApplication
}

var _ ports.UserApplication = (*UserApp)(nil)

func NewUserApp(repo ports.UserRepository, validator val.ValidatorApplication, logger logger.LoggerApplication, configurator configurator.ConfigApplication) *UserApp {
	return &UserApp{repo: repo, validator: validator, logger: logger, configurator: configurator}
}

func (app *UserApp) Create(registerReq *models.RegisterRequest) error {
	if registerReq.Password != registerReq.PasswordConfirmation {
		return customErr.ErrPasswordConfirmation
	}
	user := new(models.User)
	user.Email = registerReq.Email
	user.Password = registerReq.Password
	user.CreatedAt = time.Now()

	if err := app.validator.Struct(user); err != nil {
		app.logger.Error("Error validating user", err)
		return err
	}
	// Hash password after the validation
	user.Password = user.Hash256Password(registerReq.Password)

	err := app.repo.Save(user)
	if err != nil {
		app.logger.Error("Error creating user", err)
		return err
	}

	app.logger.Info("User created", user)
	return nil
}

func (app *UserApp) Login(credentials *models.AuthRequest) (authToken *models.AuthResponse, err error) {
	user, err := app.repo.FindByCredentials(credentials)
	if err != nil {
		app.logger.Error("Error finding user", err)
		return nil, err
	}

	config, err := app.configurator.GetConfig()
	if err != nil {
		app.logger.Error("Error getting config", err)
		return nil, err
	}

	day := time.Hour * 24

	claims := jwt.MapClaims{
		"ID":         user.ID,
		"email":      user.Email,
		"created_at": user.CreatedAt,
		"exp":        time.Now().Add(day * time.Duration(config.JWT.ExpiresIn)).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.JWT.Secret))
	if err != nil {
		return nil, err
	}
	return &models.AuthResponse{Token: t}, nil
}
