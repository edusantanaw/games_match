package main

import "github.com/gin-gonic/gin"

var PORT = "3000"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(PORT)
}
