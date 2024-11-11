package main

import (
	"log"
	"net/http"

	"github.com/Dhirajsingh212/backend/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Env not loaded")
	}

	database.ConnectDB()

	r := gin.Default()

	r.GET("/", greetingFunc)
	r.Run("localhost:8080")
}

func greetingFunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
