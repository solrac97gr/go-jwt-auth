package application

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/solrac97gr/go-jwt-auth/internal/user/domain/models"
	"time"
)

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
