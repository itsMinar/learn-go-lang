package user

import (
	"encoding/json"
	"net/http"

	"my_server/config"
	"my_server/database"
	"my_server/utils"
)

type RequestLogin struct {
	Email    string `string:"email"`
	Password string `string:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
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

	cnf := config.GetConfig()

	accessToken, err := utils.CreateJwt(cnf.JwtSecretKey, utils.Payload{
		Sub:         usr.ID,
		FirstName:   usr.FirstName,
		LastName:    usr.LastName,
		Email:       usr.Email,
		IsShopOwner: usr.IsShopOwner,
	})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, accessToken, http.StatusOK)
}
