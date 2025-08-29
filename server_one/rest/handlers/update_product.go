package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"my_server/database"
	"my_server/utils"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid product data", http.StatusBadRequest)
		return
	}

	newProduct.ID = id

	database.Update(newProduct)

	utils.SendData(w, "Successfully Updated Product", http.StatusCreated)
}
