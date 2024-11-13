package controllers

import (
	"net/http"

	"github.com/Dhirajsingh212/backend/database"
	"github.com/Dhirajsingh212/backend/models"
	"github.com/Dhirajsingh212/backend/utils"
	"github.com/gin-gonic/gin"
)

type UserInputs struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginInputs struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignupUser(c *gin.Context) {
	var userDetails UserInputs
	if err := c.BindJSON(&userDetails); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	hp := utils.HashPassowrd(userDetails.Password)
	user := models.User{Username: userDetails.Username, Email: userDetails.Email, Password: hp}
	database.DB.Create(&user)

	token := utils.GenerateToken(user.Username)

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("token", token, 3600*24, "", "", true, false)
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func SignInUser(c *gin.Context) {
	var body UserLoginInputs
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var UserDetails UserInputs
	database.DB.Where("username = ?", body.Username).First(&UserDetails)

	if UserDetails.Username == "" {
		c.JSON(http.StatusBadGateway, gin.H{"success": false, "message": "User does not exits"})
		return
	}

	isValidPassword := utils.VerifyPassword(UserDetails.Password, body.Password)

	if !isValidPassword {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Wrong password"})
		return
	}

	token := utils.GenerateToken(UserDetails.Username)
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("token", token, 3600*24, "", "", true, false)
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func GetAllUser(c *gin.Context) {
	var users []models.User

	database.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func DeleteUserById(c *gin.Context) {
	var user models.User
	if err := database.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}
	database.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": "success"})
}
