package rest

import (
	"net/http"

	"my_server/rest/handlers"
	"my_server/rest/middlewares"
)

func initRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		))
	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProduct),
		))
	mux.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetProduct),
		))
	mux.Handle(
		"PUT /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.UpdateProduct),
		))
	mux.Handle(
		"DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.DeleteProduct),
		))

	mux.Handle(
		"POST /users",
		manager.With(
			http.HandlerFunc(handlers.CreateUser),
		))
	mux.Handle(
		"POST /users/login",
		manager.With(
			http.HandlerFunc(handlers.Login),
		))
}
