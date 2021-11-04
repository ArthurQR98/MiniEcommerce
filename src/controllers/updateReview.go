package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/services"
)

func UpdateReview(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Id is must be send", http.StatusBadRequest)
		return
	}
	var review models.Review
	err := json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		http.Error(w, "Incorrect Data "+err.Error(), 400)
		return
	}
	var status bool
	status, err = services.UpdateReview(review, string(ID))
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
