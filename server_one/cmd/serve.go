package cmd

import (
	"my_server/config"
	"my_server/repo"
	"my_server/rest"
	"my_server/rest/handlers/product"
	"my_server/rest/handlers/user"
	"my_server/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()

	middlewares := middlewares.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)

	server.Start()
}
