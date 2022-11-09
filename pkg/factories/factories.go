package factories

import (
	"context"
	userApp "github.com/solrac97gr/go-jwt-auth/internal/user/application"
	userHdl "github.com/solrac97gr/go-jwt-auth/internal/user/infrastructure/handlers"
	userRepo "github.com/solrac97gr/go-jwt-auth/internal/user/infrastructure/repositories"
	configApp "github.com/solrac97gr/go-jwt-auth/pkg/config/application"
	configRepo "github.com/solrac97gr/go-jwt-auth/pkg/config/infrastructure/repositories"
	loggerApp "github.com/solrac97gr/go-jwt-auth/pkg/logger/application"
	loggerRepo "github.com/solrac97gr/go-jwt-auth/pkg/logger/infrastructure/repositories"
	mdlApp "github.com/solrac97gr/go-jwt-auth/pkg/middleware/application"
	mdlHdl "github.com/solrac97gr/go-jwt-auth/pkg/middleware/infrastructure/handlers"
	mdlRepo "github.com/solrac97gr/go-jwt-auth/pkg/middleware/infrastructure/repositories"
	valApp "github.com/solrac97gr/go-jwt-auth/pkg/validator/application"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

const MongoClientTimeout = 10

type Factory struct {
	//Variables
	logFilePath    string
	configFilePath string
	//Packages
	validator    *valApp.Validator
	configurator *configApp.ConfigService
	logger       *loggerApp.Logger
	dbClient     *mongo.Client
}

func NewFactory(logFilePath string, configFilePath string) *Factory {
	return &Factory{
		logFilePath:    logFilePath,
		configFilePath: configFilePath,
	}
}

func (f *Factory) InitializeMongoDB() *mongo.Client {
	if f.dbClient != nil {
		return f.dbClient
	}
	configurator := f.InitializeConfigurator()
	cfg, err := configurator.GetConfig()
	if err != nil {
		panic(err)
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), MongoClientTimeout*time.Second)
	defer cancelFunc()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		cfg.Database.URL,
	))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	f.dbClient = client
	return client
}

func (f *Factory) InitializeValidator() *valApp.Validator {
	if f.validator == nil {
		app := valApp.NewValidator()
		f.validator = app
		return app
	}
	return f.validator
}

func (f *Factory) InitializeLogger() *loggerApp.Logger {
	if f.configurator == nil {
		validator := f.InitializeValidator()
		path := f.logFilePath

		repo := loggerRepo.NewCSVFile(path)
		app := loggerApp.NewLogger(repo, validator)
		f.logger = app
		return app
	}
	return f.logger
}

func (f *Factory) InitializeConfigurator() *configApp.ConfigService {
	if f.configurator == nil {
		validator := f.InitializeValidator()
		logger := f.InitializeLogger()
		path := f.configFilePath

		repo := configRepo.NewJSONRepository(path)
		app := configApp.NewConfigService(repo, validator, logger)
		err := app.Config()
		if err != nil {
			panic(err)
		}
		f.configurator = app
		return app
	}
	return f.configurator
}

func (f *Factory) BuildMiddlewaresHandlers() *mdlHdl.FiberMdlHandlers {
	configurator := f.InitializeConfigurator()
	logger := f.InitializeLogger()
	validator := f.InitializeValidator()

	repo := mdlRepo.NewMongoMdlRepository(configurator, logger)
	app := mdlApp.NewFiberMiddlewares(repo, validator, logger)
	return mdlHdl.NewFiberMdlHandlers(app, logger, configurator)
}

func (f *Factory) BuildUserHandlers() *userHdl.UserHdl {
	configurator := f.InitializeConfigurator()
	logger := f.InitializeLogger()
	validator := f.InitializeValidator()
	dbClient := f.InitializeMongoDB()

	repo := userRepo.NewUserMongoDB(configurator, logger, dbClient)
	app := userApp.NewUserApp(repo, validator, logger, configurator)
	return userHdl.NewUserHdl(app, logger)
}
