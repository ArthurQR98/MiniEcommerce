package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ArthurQR98/e-commerce/src/auth"
	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/services"
)

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, "Incorrect Data"+err.Error(), 400)
		return
	}
	var status bool
	status, err = services.UpdateCustomer(customer, auth.IDCustomer.Hex())
	if err != nil {
		http.Error(w, "Error occurred while updating the data "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "Modification failed "+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}
