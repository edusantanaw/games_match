package main

import (
	"github.com/edusantanaw/games_match.git/adapters"
	"github.com/edusantanaw/games_match.git/controllers"
	"github.com/edusantanaw/games_match.git/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.OpenConnection()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/game", adapters.GinAdapter[*controllers.CreateGameData](controllers.RegisterGame))
	r.Run()
}
