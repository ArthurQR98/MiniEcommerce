package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/services"
)

// TODO: reparar los mensajes a ingles
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "Error"+err.Error(), 400)
		return
	}
	register := models.Category{
		Name:        category.Name,
		Description: category.Description,
	}
	_, status, err := services.CreateCategory(register)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro, intente nuevamente.", 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado insertar la categoria", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
