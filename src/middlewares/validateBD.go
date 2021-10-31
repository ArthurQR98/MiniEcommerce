package middlewares

import (
	"net/http"

	"github.com/ArthurQR98/e-commerce/config"
)

func ValidateDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if config.CheckConnection() == 0 {
			http.Error(w, "Lost connection with database", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
