package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/services"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "Error creating category"+err.Error(), 400)
		return
	}
	register := models.Category{
		Name:        category.Name,
		Description: category.Description,
	}
	_, status, err := services.CreateCategory(register)
	if err != nil {
		http.Error(w, "An error occurred while trying to insert the record, please try again. ", 400)
		return
	}
	if !status {
		http.Error(w, "The category could not be inserted ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
