package main

import (
	"github.com/edusantanaw/games_match.git/db"
	"github.com/edusantanaw/games_match.git/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.OpenConnection()
	routes.Routes(r)
	r.Run()
}
