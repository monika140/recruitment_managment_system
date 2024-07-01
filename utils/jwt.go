package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Defines the secret key for JWT.
var jwtKey = []byte("my_secrete_key")

// Function to generate a JWT token.
func GenerateToken(userID uint) (string, error) {
	//Creates a new JWT token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	//Signs the token with the secret key.
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
// func ValidateJWT(tokenString string) (*jwt.Token, error) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		return token, nil
// 	}

// 	return nil, errors.New("invalid token")
// }
