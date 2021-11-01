package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/services"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, "Error "+err.Error(), 400)
		return
	}
	_, status, err := services.CreateProduct(product)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro, intente nuevamente.", 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado insertar el producto", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
