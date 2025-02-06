package routes

import (
	"docker-product/handlers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/create-product", handlers.CreateProduct).Methods("POST")

	return router
}
