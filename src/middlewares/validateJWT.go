package middlewares

import (
	"net/http"

	"github.com/ArthurQR98/e-commerce/src/auth"
)

func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header.Get("Authorization")) == 0 {
			http.Error(w, "Token not sent ", http.StatusBadRequest)
			return
		}
		_, _, _, err := auth.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Token not valid - "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
