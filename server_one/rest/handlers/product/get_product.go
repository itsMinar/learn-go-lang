package product

import (
	"net/http"
	"strconv"

	"my_server/database"
	"my_server/utils"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product := database.GetSingle(id)
	if product == nil {
		utils.SendError(w, http.StatusNotFound, "Product not found")
		return
	}

	utils.SendData(w, product, http.StatusCreated)
}
