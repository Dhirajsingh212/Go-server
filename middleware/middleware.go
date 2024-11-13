package middleware

import (
	"net/http"
	"strings"

	"github.com/Dhirajsingh212/backend/utils"
	"github.com/gin-gonic/gin"
)

func ProtectedCheck(c *gin.Context) {
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	tokenString := strings.Split(cookie.String(), "=")[1]
	isValid := utils.VerifyToken(tokenString)
	if !isValid {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}
	c.Next()
}