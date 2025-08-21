package handlers

import (
	"net/http"
	"strconv"

	"my_server/database"
	"my_server/utils"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("productId")

	id, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == id {
			utils.SendData(w, product, http.StatusOK)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}
