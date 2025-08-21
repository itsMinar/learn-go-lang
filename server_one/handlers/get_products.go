package handlers

import (
	"net/http"

	"my_server/database"
	"my_server/utils"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, database.ProductList, http.StatusOK)
}
