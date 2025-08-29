package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"my_server/config"
	"my_server/rest/middlewares"
)

func Start(cnf config.Config) {
	manager := middlewares.NewManager()
	manager.Use(
		middlewares.Preflight,
		middlewares.Cors,
		middlewares.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("Server is running on port", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
