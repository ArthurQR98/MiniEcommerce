package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ArthurQR98/e-commerce/src/auth"
	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/services"
)

// TODO: reparar los mensajes a ingles
func CreateReview(w http.ResponseWriter, r *http.Request) {
	var review models.Review
	err := json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		http.Error(w, "Ocurrio un error", 400)
		return
	}
	register := models.Review{
		PostDate: time.Now(),
		Title:    review.Title,
		Body:     review.Body,
		Customer: auth.IDCustomer,
		Rating:   review.Rating,
	}
	_, status, err := services.CreateReview(register)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro, intente nuevamente. "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado insertar el review", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
