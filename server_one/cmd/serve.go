package cmd

import (
	"fmt"
	"net/http"

	"my_server/global_router"
	"my_server/middlewares"
)

func Serve() {
	manager := middlewares.NewManager()
	manager.Use(middlewares.Logger, middlewares.Hudai)

	mux := http.NewServeMux()

	initRoutes(mux, manager)

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server is running on port 8000")
	err := http.ListenAndServe(":8000", globalRouter)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
