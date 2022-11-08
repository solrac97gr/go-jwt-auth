package ports

import "github.com/solrac97gr/go-jwt-auth/internal/user/domain/models"

type UserApplication interface {
	Create(user models.User) error
}
