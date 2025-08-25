package cmd

import (
	"fmt"
	"net/http"

	"my_server/middlewares"
)

func Serve() {
	manager := middlewares.NewManager()
	manager.Use(
		middlewares.Preflight,
		middlewares.Cors,
		middlewares.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	fmt.Println("Server is running on port 8000")
	err := http.ListenAndServe(":8000", wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
