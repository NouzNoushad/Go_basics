package main

import (
	"chat/database"
	"chat/routes"
)

func main() {
	database.InitDB()

	r := routes.Routes()

	r.Run(":8015")
}