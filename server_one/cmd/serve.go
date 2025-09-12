package cmd

import (
	"my_server/config"
	"my_server/rest"
	"my_server/rest/handlers/product"
	"my_server/rest/handlers/user"
	"my_server/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	middlewares := middlewares.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler()

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
