package utils

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassowrd(password string) string {
	byteData, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal("Failed to hash password")
	}
	return string(byteData)
}

func VerifyPassword(dbPassword string, reqPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(reqPassword))
	return err == nil
}

func GenerateToken(username string) string {
	secretKey := os.Getenv("SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return ""
	}
	return tokenString
}

func VerifyToken(tokenString string) bool {
	secretKey := os.Getenv("SECRET_KEY")
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return false
	}
	if !token.Valid {
		return false
	}

	return true
}
