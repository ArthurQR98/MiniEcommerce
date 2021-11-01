package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/services"
	"github.com/asaskevich/govalidator"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, "Data error received "+err.Error(), 400)
		return
	}
	if len(customer.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}
	if len(customer.Password) <= 6 {
		http.Error(w, "Something 6 character password", 400)
		return
	}
	_, find, _ := services.CheckIfExistUser(customer.Email)
	if find {
		http.Error(w, "Customer is already registered", 400)
		return
	}

	if !govalidator.IsExistingEmail(customer.Email) {
		http.Error(w, "Email is not valid", 500)
		return
	}
	_, status, err := services.RegisterCustomer(customer)
	if err != nil {
		http.Error(w, "Error ocurred while registering "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "Could not insert customer record "+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
