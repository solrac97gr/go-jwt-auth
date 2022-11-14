package repositories

import (
	"context"
	"github.com/solrac97gr/go-jwt-auth/internal/user/domain/models"
)

func (repo *UserMongoDB) Save(user *models.User) error {
	_, err := repo.collection.InsertOne(context.Background(), user)
	if err != nil {
		repo.logger.Error("Error saving user", err)
		return err
	}
	return nil
}
