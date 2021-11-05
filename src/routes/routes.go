package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ArthurQR98/e-commerce/src/controllers"
	"github.com/ArthurQR98/e-commerce/src/middlewares"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Routes() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("E-commerce APIRestful - Golang")
	}).Methods("GET")

	router.HandleFunc("/auth/sign-up", middlewares.ValidateDB(controllers.Register)).Methods("POST")
	router.HandleFunc("/auth/sign-in", middlewares.ValidateDB(controllers.Login)).Methods("POST")

	router.HandleFunc("/users", middlewares.ValidateDB(middlewares.ValidateJWT(controllers.ReadCustomers))).Methods("GET")
	router.HandleFunc("/users", middlewares.ValidateDB(middlewares.ValidateJWT(controllers.UpdateCustomer))).Methods("PUT")

	router.HandleFunc("/categories", middlewares.ValidateDB(middlewares.ValidateJWT(controllers.CreateCategory))).Methods("POST")
	router.HandleFunc("/categories", middlewares.ValidateDB(middlewares.ValidateJWT(controllers.ReadCategorys))).Methods("GET")
	router.HandleFunc("/categories", middlewares.ValidateDB(middlewares.ValidateJWT(controllers.UpdateCategory))).Methods("PUT")

	router.HandleFunc("/reviews", middlewares.ValidateDB(middlewares.ValidateJWT(controllers.CreateReview))).Methods("POST")
	router.HandleFunc("/reviews", middlewares.ValidateDB(middlewares.ValidateJWT(controllers.ReadReviews))).Methods("GET")
	router.HandleFunc("/reviews", middlewares.ValidateDB(middlewares.ValidateJWT(controllers.UpdateReview))).Methods("PUT")

	router.HandleFunc("/products", middlewares.ValidateDB(middlewares.ValidateJWT(controllers.CreateProduct))).Methods("POST")
	router.HandleFunc("/products", middlewares.ValidateDB(middlewares.ValidateJWT(controllers.ReadProducts))).Methods("GET")
	router.HandleFunc("/products", middlewares.ValidateDB(middlewares.ValidateJWT(controllers.UpdateProduct))).Methods("PUT")
	router.HandleFunc("/products", middlewares.ValidateDB(middlewares.ValidateJWT(controllers.DeleteProduct))).Methods("DELETE")

	router.HandleFunc("/orders", middlewares.ValidateDB(middlewares.ValidateJWT(controllers.CreateOrder))).Methods("POST")
	router.HandleFunc("/orders", middlewares.ValidateDB(middlewares.ValidateJWT(controllers.ReadOrders))).Methods("GET")
	router.HandleFunc("/orders", middlewares.ValidateDB(middlewares.ValidateJWT(controllers.UpdateOrder))).Methods("PUT")

	PORT := os.Getenv("PORT")
	controller := cors.AllowAll().Handler(router)
	fmt.Println("-- Backend running on PORT " + PORT + " --")
	log.Fatal(http.ListenAndServe(":"+PORT, controller))
}
