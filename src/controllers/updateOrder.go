package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/services"
	"github.com/ArthurQR98/e-commerce/src/utils"
)

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Id is must be send", http.StatusBadRequest)
		return
	}
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Incorrect Data "+err.Error(), 400)
		return
	}

	if len(order.Items.Product) > 0 {
		for _, v := range order.Items.Product {
			find, _ := utils.IsExistsProduct(v)
			if !find {
				http.Error(w, "Some of the products do not exist ", 400)
				return
			}
		}
	}

	var status bool
	status, err = services.UpdateOrder(order, string(ID))
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
