package ports

import "github.com/solrac97gr/go-jwt-auth/internal/user/domain/models"

type UserApplication interface {
	Create(registerRequest *models.RegisterRequest) error
	Login(credentials *models.AuthRequest) (authToken *models.AuthResponse, err error)
}
