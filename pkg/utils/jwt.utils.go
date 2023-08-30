package utils

import (
	"os"
	"time"

	"github.com/form3tech-oss/jwt-go"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(adminUsername string) (string, error) {
	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set token claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = adminUsername
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Generate the token
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
