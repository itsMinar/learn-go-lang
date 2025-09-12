package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"my_server/config"
	"my_server/rest/handlers/product"
	"my_server/rest/handlers/user"
	"my_server/rest/middlewares"
)

type Server struct {
	// add dependencies
	cnf            *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

// inject dependencies
func NewServer(cnf *config.Config, productHandler *product.Handler, userHandler *user.Handler) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (server *Server) Start() {
	manager := middlewares.NewManager()
	manager.Use(
		middlewares.Preflight,
		middlewares.Cors,
		middlewares.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("Server is running on port", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
