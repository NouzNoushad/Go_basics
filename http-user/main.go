package main

import (
	"fmt"
	"http-user/database"
	"http-user/routes"
	"net/http"
)

func main() {
	database.InitDB()

	r := routes.SetupRoutes()

	port := "8020"

	fmt.Printf("Server is running on port %s...", port)

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
