package routes

import (
	"docker-product/handlers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/create-product", handlers.CreateProduct).Methods("POST")
	router.HandleFunc("/get-products", handlers.GetProducts).Methods("GET")
	router.HandleFunc("/get-product/{id}", handlers.GetProductDetails).Methods("GET")
	router.HandleFunc("/delete-product/{id}", handlers.DeleteProductDetails).Methods("DELETE")
	router.HandleFunc("/update-product/{id}", handlers.UpdateProductDetails).Methods("PUT")

	return router
}
