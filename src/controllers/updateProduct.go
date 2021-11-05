package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/services"
	"github.com/ArthurQR98/e-commerce/src/utils"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Id is must be send", http.StatusBadRequest)
		return
	}
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Incorrect Data "+err.Error(), 400)
		return
	}

	find, _ := utils.IsExistsCategory(product.Categories)
	if !find {
		http.Error(w, "Category don't exist ", 400)
		return
	}

	if len(product.Reviews) > 0 {
		for _, v := range product.Reviews {
			find, _ := utils.IsExistsReview(v)
			if !find {
				http.Error(w, "Some of the reviews do not exist ", 400)
				return
			}
		}
	}

	var status bool
	status, err = services.UpdateProduct(product, string(ID))
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
