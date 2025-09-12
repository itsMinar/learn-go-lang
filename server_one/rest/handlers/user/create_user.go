package user

import (
	"encoding/json"
	"net/http"

	"my_server/database"
	"my_server/utils"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid Request data", http.StatusBadRequest)
		return
	}

	createdUser := newUser.Store()

	utils.SendData(w, createdUser, http.StatusCreated)
}
