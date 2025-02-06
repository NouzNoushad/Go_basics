package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"user_auth/config"
	"user_auth/controllers"

	"github.com/gorilla/mux"
)

func main() {
	// initialize db
	config.InitDB()

	// setup router
	router := mux.NewRouter()
	router.HandleFunc("/signup", controllers.SignupHandler).Methods("POST")

	// port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8060"
	}

	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}