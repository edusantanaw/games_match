package main

import (
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
	r.POST("/game", func(ctx *gin.Context) {
		var requestBody controllers.CreateGameData
		if err := ctx.BindJSON(&requestBody); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		gameJson, err := controllers.RegisterGame(requestBody)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gameJson)
	})
	r.Run()
}
