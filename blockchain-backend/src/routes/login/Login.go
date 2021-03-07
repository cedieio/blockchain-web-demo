package login

import (
	"backend/src/lib/token"
	"backend/src/models/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

var users = make(map[string]uint64)

func init() {
	users["cedieggg@gmail.com"] = 1
	users["cedieggg2@gmail.com"] = 2
}

// Post is used for /login post request, handles and validates login
func Post(c *gin.Context) {
	var json user.Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if json.Username == "" || json.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid credentials",
		})
		return
	}
	userID := users[json.Username]
	if userID != 0 {
		token, err := token.CreateToken(userID, "jdnfksdmfksd")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "You're loggin in !",
			"token":  token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "user not found",
		})
	}
}
