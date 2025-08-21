package cmd

import (
	"fmt"
	"net/http"

	"my_server/global_router"
	"my_server/handlers"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /products/{productId}", http.HandlerFunc(handlers.GetProductById))

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server is running on port 8000")
	err := http.ListenAndServe(":8000", globalRouter)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
