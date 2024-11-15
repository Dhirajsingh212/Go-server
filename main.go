package main

import (
	"log"
	"net/http"

	"github.com/Dhirajsingh212/backend/controllers"
	"github.com/Dhirajsingh212/backend/database"
	"github.com/Dhirajsingh212/backend/middleware"
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

	r.GET("/health-check", greetingFunc)

	// USER ROUTES
	r.POST("/signup", controllers.SignupUser)
	r.POST("/signin", controllers.SignInUser)
	r.GET("/getAllUser", middleware.ProtectedCheck, controllers.GetAllUser)
	r.DELETE("/delete/:id", middleware.ProtectedCheck, controllers.DeleteUserById)
	r.GET("/user/:id", middleware.ProtectedCheck, controllers.GetSingleUser)
	r.Run("localhost:8080")
}

func greetingFunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Server is healthy"})
}
