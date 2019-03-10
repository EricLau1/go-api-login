package controllers

import (
	"net/http"
	"encoding/json"
	"go-api-login/api/auth"
	"go-api-login/api/utils"
	"go-api-login/api/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body := utils.BodyParser(r)
	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnauthorized)
		return
	}
	token, err := auth.SignIn(user.Email, user.Password)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnauthorized)
		return
	}
	utils.ToJson(w, token, http.StatusOK)
}