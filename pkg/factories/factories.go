package factories

import (
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
)

type Factory struct {
	//Variables
	logFilePath    string
	configFilePath string
	//Packages
	validator    *valApp.Validator
	configurator *configApp.ConfigService
	logger       *loggerApp.Logger
}

func NewFactory(logFilePath string, configFilePath string) *Factory {
	return &Factory{
		logFilePath:    logFilePath,
		configFilePath: configFilePath,
	}
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
		validator := f.validator
		logger := f.logger
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
	configurator := f.configurator
	logger := f.logger
	validator := f.validator

	repo := mdlRepo.NewMongoMdlRepository(configurator, logger)
	app := mdlApp.NewFiberMiddlewares(repo, validator, logger)
	return mdlHdl.NewFiberMdlHandlers(app, logger)
}

func (f *Factory) BuildUserHandlers() *userHdl.UserHdl {
	configurator := f.configurator
	logger := f.logger
	validator := f.validator

	repo := userRepo.NewUserMongoDB(configurator, logger)
	app := userApp.NewUserApp(repo, validator, logger)
	return userHdl.NewUserHdl(app, logger)
}
