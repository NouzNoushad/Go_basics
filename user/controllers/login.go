package controllers

import (
	"time"
	"user/config"
	"user/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// generate jwt token
func generateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("Thisisasecretkey"))
}

func Login(c *gin.Context) {
	var userInput models.User
	var user models.User

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// check user exists
	if err := config.Db.Where("email = ?", userInput.Email).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	// generate jwt token
	token, err := generateToken(user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, gin.H{"message": "User login success", "token": token})
}