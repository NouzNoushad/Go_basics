package middleware

import (
	"fmt"
	"strings"
	"user/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

		if tokenString == "" {
			c.JSON(400, gin.H{"error": "Authorization token is missing"})
			c.Abort()
			return
		}

		// validate token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return config.SecretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Token expired"})
			c.Abort()
			return
		}

		c.Next()
	}
}