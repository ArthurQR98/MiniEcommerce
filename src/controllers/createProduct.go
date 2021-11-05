package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/services"
	"github.com/ArthurQR98/e-commerce/src/utils"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, "Error creating product "+err.Error(), 400)
		return
	}

	find, _ := utils.FindIfExistCategory(product.Categories)
	if !find {
		http.Error(w, "Category don't exist ", 400)
		return
	}

	if len(product.Reviews) > 0 {
		for _, v := range product.Reviews {
			find, _ := utils.FindIfExistReview(v)
			if !find {
				http.Error(w, "Some of the reviews do not exist ", 400)
				return
			}
		}
	}

	_, status, err := services.CreateProduct(product)
	if err != nil {
		http.Error(w, "An error occurred while trying to register, please try again. ", 400)
		return
	}
	if !status {
		http.Error(w, "The product could not be saved.", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
