package product

import (
	"net/http"

	"my_server/database"
	"my_server/utils"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, database.List(), http.StatusOK)
}
