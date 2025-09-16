package product

import (
	"my_server/repo"
	"my_server/rest/middlewares"
)

type Handler struct {
	middlewares *middlewares.Middlewares
	productRepo repo.ProductRepo
}

func NewHandler(
	middlewares *middlewares.Middlewares,
	productRepo repo.ProductRepo,
) *Handler {
	return &Handler{
		middlewares: middlewares,
		productRepo: productRepo,
	}
}
