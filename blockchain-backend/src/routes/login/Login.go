package login

import (
	"backend/src/models/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Post is used for /login post request, handles and validates login
func Post(c *gin.Context) {
	var json user.Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	if json.Username == "" || json.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid credentials",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "You're loggin in !",
	})
}
