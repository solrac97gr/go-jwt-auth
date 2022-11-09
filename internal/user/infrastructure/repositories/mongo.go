package repositories

import (
	"context"
	"github.com/solrac97gr/go-jwt-auth/internal/user/domain/models"
	"github.com/solrac97gr/go-jwt-auth/internal/user/domain/ports"
	config "github.com/solrac97gr/go-jwt-auth/pkg/config/domain/ports"
	logger "github.com/solrac97gr/go-jwt-auth/pkg/logger/domain/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoDB struct {
	client       *mongo.Client
	database     *mongo.Database
	collection   *mongo.Collection
	configurator config.ConfigApplication
	logger       logger.LoggerApplication
}

var _ ports.UserRepository = (*UserMongoDB)(nil)

func NewUserMongoDB(config config.ConfigApplication, logger logger.LoggerApplication, client *mongo.Client) *UserMongoDB {
	cfg, err := config.GetConfig()
	if err != nil {
		logger.Error("Error getting configurator", err)
		return nil
	}
	return &UserMongoDB{
		configurator: config,
		logger:       logger,
		client:       client,
		database:     client.Database(cfg.Database.Name),
		collection:   client.Database(cfg.Database.Name).Collection("users"),
	}
}

func (repo *UserMongoDB) Save(user *models.User) error {
	_, err := repo.collection.InsertOne(context.Background(), user)
	if err != nil {
		repo.logger.Error("Error saving user", err)
		return err
	}
	return nil
}

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
