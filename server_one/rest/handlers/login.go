package handlers

import (
	"encoding/json"
	"net/http"

	"my_server/database"
	"my_server/utils"
)

type RequestLogin struct {
	Email    string `string:"email"`
	Password string `string:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin database.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		http.Error(w, "Invalid Request data", http.StatusBadRequest)
		return
	}

	usr := database.Find(reqLogin.Email, reqLogin.Password)
	if usr == nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	utils.SendData(w, usr, http.StatusOK)
}
