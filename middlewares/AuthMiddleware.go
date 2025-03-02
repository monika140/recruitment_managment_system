package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Defines the secret key for JWT.
var jwtKey = []byte("my_secrete_key")

// Function to create an authentication middleware.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" { //Checks if the token is empty.
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		//Parses the token.
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}

			return jwtKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		//Validates the token and extracts claims.
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", claims["user_id"])
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		}
	}
}

func ApplicantOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ensure AuthMiddleware has been called
		userID, exists := c.Get("user_id")
		if !exists {
			fmt.Println("No user found in context")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No user found"})
			c.Abort()
			return
		}

		userType, exists := c.Get("user_type")
		if !exists || userType != "applicant" {
			fmt.Printf("User %v does not have permission. UserType: %v", userID, userType)
			c.JSON(http.StatusForbidden, gin.H{"error": "Only applicants can access this endpoint"})
			c.Abort()
			return
		}

		fmt.Printf("User %v is authorized as applicant", userID)
		c.Next()
	}
}
