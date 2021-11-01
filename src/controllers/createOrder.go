package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ArthurQR98/e-commerce/src/auth"
	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/services"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Error creating order", 400)
		return
	}
	register := models.Order{
		User:  auth.IDCustomer,
		State: order.State,
		Items: order.Items,
		Date:  time.Now(),
	}
	register.Items.Quantity = len(register.Items.Product)
	_, status, err := services.CreateOrder(register)
	if err != nil {
		http.Error(w, "An error occurred while trying to register, please try again. "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "The order could not be saved", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
