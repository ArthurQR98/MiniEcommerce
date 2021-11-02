package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ArthurQR98/e-commerce/src/services"
)

func ReadCategorys(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "You must send the page parameter", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "You must send the page parameter with a value greater than 0", http.StatusBadRequest)
		return
	}
	pag := int64(page)
	response, isCorrect := services.GetCategorys(pag)
	if !isCorrect {
		http.Error(w, "Error reading categories", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
