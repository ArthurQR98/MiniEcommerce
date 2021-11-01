package middlewares

import (
	"net/http"

	"github.com/ArthurQR98/e-commerce/src/auth"
)

// TODO: reparar los mensajes a ingles
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header.Get("Authorization")) == 0 {
			http.Error(w, "Error token no enviado", http.StatusBadRequest)
			return
		}
		_, _, _, err := auth.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error token - "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
