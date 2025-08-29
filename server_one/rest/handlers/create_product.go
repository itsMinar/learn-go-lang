package handlers

import (
	"encoding/json"
	"net/http"

	"my_server/database"
	"my_server/utils"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid product data", http.StatusBadRequest)
		return
	}

	createdProduct := database.Store(newProduct)

	utils.SendData(w, createdProduct, http.StatusCreated)
}
