package main

import (
	"github.com/solrac97gr/go-jwt-auth/docs"
	"github.com/solrac97gr/go-jwt-auth/pkg/factories"
	"github.com/solrac97gr/go-jwt-auth/pkg/server"
)

// @contact.name   Carlos García
// @contact.email  cgarciarosales97@gmail.com
// @license.name   Apache 2.0
// @license.url    http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	factory := factories.NewFactory(
		"logs/log.csv",
		"./config/env.json",
	)

	// Create and Share the dependencies throw the factory [Only one instance]
	factory.InitializeValidator()
	configurator := factory.InitializeConfigurator()
	factory.InitializeLogger()
	factory.InitializeMongoDB()

	middlewares := factory.BuildMiddlewaresHandlers()
	usersHandlers := factory.BuildUserHandlers()

	server := server.NewServer(middlewares, usersHandlers, configurator)

	config, err := configurator.GetConfig()
	if err != nil {
		panic(err)
	}

	docs.SwaggerInfo.Title = config.App.Name
	docs.SwaggerInfo.Description = config.App.Description
	docs.SwaggerInfo.Version = config.App.Version
	docs.SwaggerInfo.Host = config.App.Host
	docs.SwaggerInfo.BasePath = config.App.BaseURL

	err = server.Run(config.Server.Port)
	if err != nil {
		panic(err)
	}
}
