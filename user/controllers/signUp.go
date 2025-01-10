package controllers

import (
	"user/config"
	"user/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// signup user
func SignUp(c *gin.Context) {
	var user models.User

	// if err := c.ShouldBind(&user); err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }

	user.Name = c.PostForm("name")
	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")

	if user.Name == "" || user.Email == "" || user.Password == "" {
		c.JSON(400, gin.H{"error": "All fields are required"})
		return
	}

	user.Id = uuid.New().String()

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error hashing password"})
		return
	}

	// save the user
	user.Password = string(hash)
	if err := config.Db.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(200, gin.H{"message": "User created"})
}
