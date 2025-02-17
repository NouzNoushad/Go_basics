package routes

import (
	"http-user/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/create-user", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/get-users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/get-user/{id}", handlers.GetUser).Methods("GET")

	return r
}
