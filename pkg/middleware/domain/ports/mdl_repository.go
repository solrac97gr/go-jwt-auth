package ports

import "github.com/solrac97gr/go-jwt-auth/pkg/middleware/domain/models"

type MiddlewareRepository interface {
	Login(request *models.AuthRequest) error
	Authenticate(token string) error
}
