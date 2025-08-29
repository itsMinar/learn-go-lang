package handlers

import (
	"net/http"
	"strconv"

	"my_server/database"
	"my_server/utils"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	database.Delete(id)

	utils.SendData(w, "Successfully Deleted Product", http.StatusOK)
}
