package cmd

import (
	"fmt"
	"my_server/config"
	"my_server/infra/db"
	"my_server/repo"
	"my_server/rest"
	"my_server/rest/handlers/product"
	"my_server/rest/handlers/user"
	"my_server/rest/middlewares"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println("Failed to connect to database: ", err)
		os.Exit(1)
	}

	err = db.Migrate(dbCon, "./migrations")
	if err != nil {
		fmt.Println("Failed to migrate database: ", err)
		os.Exit(1)
	}

	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

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
