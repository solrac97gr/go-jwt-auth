package repositories

import (
	"context"
	"github.com/solrac97gr/go-jwt-auth/internal/user/domain/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *UserMongoDB) FindByCredentials(credentials *models.AuthRequest) (*models.User, error) {
	user := new(models.User)
	user.Email = credentials.Email
	user.Password = user.Hash256Password(credentials.Password)

	result := repo.database.Collection("users").FindOne(context.Background(), bson.D{
		{Key: "email", Value: user.Email},
		{Key: "password", Value: user.Password},
	})

	if result.Err() != nil {
		repo.logger.Error("Error finding user", result.Err())
		return nil, result.Err()
	}

	err := result.Decode(user)
	if err != nil {
		repo.logger.Error("Error decoding user", err)
		return nil, err
	}
	user.Password = ""
	return user, nil
}
