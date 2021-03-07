package main

import (
	"backend/src/routes/login"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	routes := gin.Default()

	routes.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"response": "Hey world !",
		})
	})

	routes.POST("/login", login.Post)

	err := routes.Run(":9090")
	if err != nil {
		log.Fatalf("Error in running gin ! : %+v \n", err)
	}
}
