package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProtectedCheck(c *gin.Context) {
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(cookie)
	c.Next()
}
