package main

import (
	"docker-product/database"
	"docker-product/routes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	database.InitDB()

	router := routes.Router()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
