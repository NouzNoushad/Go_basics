package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"http-user/database"
	"http-user/models"
	"net/http"

	"github.com/gorilla/mux"
)

func JsonResponse(w http.ResponseWriter, statusCode int, responseMessage map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(responseMessage)
}

// Create a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	query := "INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id"
	err := database.DB.QueryRow(query, user.Username, user.Email).Scan(&user.Id)
	if err != nil {
		JsonResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error": "Error creating user",
		})
		return
	}

	// success
	JsonResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "User created successfully",
		"user":    user,
	})
}

// Get users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	rows, err := database.DB.Query("SELECT id, username, email FROM users")
	if err != nil {
		JsonResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error": "Error fetching user",
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Email); err != nil {
			JsonResponse(w, http.StatusInternalServerError, map[string]interface{}{
				"error": "Error scanning user",
			})
			return
		}
		users = append(users, user)
	}

	// success
	JsonResponse(w, http.StatusOK, map[string]interface{}{
		"data": users,
		"items": fmt.Sprintf("%d items", len(users)),
	})
}

// Get user by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var user models.User
	err := database.DB.QueryRow("SELECT id, username, email FROM users WHERE id=$1", id).Scan(&user.Id, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			JsonResponse(w, http.StatusNotFound, map[string]interface{}{
				"error": "User not found",
			})
		} else {
			JsonResponse(w, http.StatusInternalServerError, map[string]interface{}{
				"error": "Error fetching user",
			})
		}
		return
	}

	// success
	JsonResponse(w, http.StatusOK, map[string]interface{}{
		"data": user,
	})
}