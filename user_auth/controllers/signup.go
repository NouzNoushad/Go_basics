package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"user_auth/config"
	"user_auth/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// validation
func isEmptyField(field string) bool {
	return field == ""
}

// check email already taken
func isEmailTaken(email string) bool {
	var user models.User
	result := config.DB.Where("email = ?", email).First(&user)
	return result.Error == nil
}

// validate password
func validPassword(password string) bool {
	// password length
	if len(password) < 8 {
		return false
	}

	// check if password contains at least one uppercase letter
	hasUpperCase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	if !hasUpperCase {
		return false
	}

	// check if password contains atleast one number
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	if !hasNumber {
		return false
	}

	// check if password contains atleast one special character
	hasSpecialCharacter := regexp.MustCompile(`[!@#$^&*(),.?":{}|<>]`).MatchString(password)
	return hasSpecialCharacter
}

// hash password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	user.Id = uuid.New().String()

	// username
	if isEmptyField(user.Username) {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	// email
	if isEmptyField(user.Email) {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	// check email taken
	if isEmailTaken(user.Email) {
		http.Error(w, "Email already taken", http.StatusBadRequest)
		return
	}

	// check password
	if !validPassword(user.Password) {
		http.Error(w, "Password must be at least 8 characters long, contain at least one uppercase letter, one number and one special character", http.StatusBadRequest)
		return
	}

	// hash password
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	// save user to database
	if err := config.DB.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// success
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "User created", "user": user})
}
