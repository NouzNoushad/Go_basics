package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := "host=localhost port=5432 user=postgres password=noushad dbname=http-user sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database not responding:", err)
	}
	fmt.Println("Connected to database")

	// create table
	CreateTable()
}

func CreateTable() {
	query := `
	CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		username TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL
	)`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Error creating table:", err)
		return
	}
	fmt.Println("Table created successfully")
}
