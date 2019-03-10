package controllers

import (
	"net/http"
	"go-api-login/api/utils"
	"go-api-login/api/models"
)

func PublicRoute(w http.ResponseWriter, r *http.Request) {
	utils.ToJson(w, "Rota pública", http.StatusOK)
}

func ProtectedRoute(w http.ResponseWriter, r *http.Request) {
	jwtParams, err := utils.JwtExtract(r)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnauthorized)
		return
	}
	email, ok := jwtParams["user_email"].(string)
	if !ok {
		utils.ToJson(w, "Payload inválido", http.StatusUnauthorized)
		return
	}
	user := models.GetUserByEmail(email)
	if user.Id == 0 {
		utils.ToJson(w, "Usuário não encontrado", http.StatusUnauthorized)
		return
	}
	utils.ToJson(w, user, http.StatusOK)
}