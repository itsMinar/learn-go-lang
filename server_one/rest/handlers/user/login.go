package user

import (
	"encoding/json"
	"net/http"

	"my_server/utils"
)

type RequestLogin struct {
	Email    string `string:"email"`
	Password string `string:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req RequestLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid Request data")
		return
	}

	usr, err := h.userRepo.Find(req.Email, req.Password)
	if err != nil {
		utils.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	if usr == nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid credentials")
		return
	}

	accessToken, err := utils.CreateJwt(h.cnf.JwtSecretKey, utils.Payload{
		Sub:         usr.ID,
		FirstName:   usr.FirstName,
		LastName:    usr.LastName,
		Email:       usr.Email,
		IsShopOwner: usr.IsShopOwner,
	})

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	utils.SendData(w, http.StatusOK, accessToken)
}
