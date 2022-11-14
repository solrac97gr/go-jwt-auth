package repositories

import (
	"github.com/solrac97gr/go-jwt-auth/internal/user/domain/ports"
	config "github.com/solrac97gr/go-jwt-auth/pkg/config/domain/ports"
	logger "github.com/solrac97gr/go-jwt-auth/pkg/logger/domain/ports"
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
