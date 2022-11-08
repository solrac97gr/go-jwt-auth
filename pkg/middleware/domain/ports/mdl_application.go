package ports

import "github.com/solrac97gr/go-jwt-auth/pkg/middleware/domain/models"

type MiddlewareApplication interface {
	Authenticate(token string) (code int, err error)
	Login(request *models.AuthRequest) (code int, authResponse *models.AuthResponse, err error)
}
