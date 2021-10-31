package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ArthurQR98/e-commerce/src/middlewares"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Routes() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Hello Golang")
	}).Methods("GET")

	//TODO: probar los nuevo endpoints
	router.HandleFunc("/auth/sign-up", middlewares.ValidateDB(Register)).Methods("POST")
	router.HandleFunc("/auth/sign-in", middlewares.ValidateDB(Login)).Methods("POST")

	PORT := os.Getenv("PORT")
	controller := cors.AllowAll().Handler(router)
	fmt.Println("-- Backend running on PORT " + PORT + " --")
	log.Fatal(http.ListenAndServe(":"+PORT, controller))
}
