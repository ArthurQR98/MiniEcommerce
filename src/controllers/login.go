package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ArthurQR98/e-commerce/src/auth"
	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/services"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var customer models.Customer

	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, "User or password is invalid "+err.Error(), 400)
		return
	}
	if len(customer.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}
	user, find := services.LoginCustomer(customer.Email, customer.Password)
	if !find {
		http.Error(w, "User or password is invalid", 500)
		return
	}
	jwtKey, err := auth.GenerateJWT(user)
	if err != nil {
		http.Error(w, "Error generate token "+err.Error(), 400)
		return
	}

	resp := models.ResponseLogin{
		Token: jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
