package controllers

import (
	"net/http"

	"github.com/ArthurQR98/e-commerce/src/services"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Id is must be sent", http.StatusBadRequest)
		return
	}
	err := services.DeleteProduct(ID)
	if err != nil {
		http.Error(w, "Failed to delete", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
