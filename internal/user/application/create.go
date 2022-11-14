package application

import (
	"github.com/solrac97gr/go-jwt-auth/internal/user/domain/models"
	customErr "github.com/solrac97gr/go-jwt-auth/pkg/custom-errors"
	"time"
)

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
