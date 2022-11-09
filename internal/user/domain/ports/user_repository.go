package ports

import "github.com/solrac97gr/go-jwt-auth/internal/user/domain/models"

type UserRepository interface {
	Save(user *models.User) error
	FindByCredentials(credentials *models.AuthRequest) (*models.User, error)
}
