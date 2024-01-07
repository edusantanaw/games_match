package main

import (
	"time"

	"github.com/edusantanaw/games_match.git/config"
	"github.com/edusantanaw/games_match.git/db"
	"github.com/edusantanaw/games_match.git/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.SetAllEnvs()
	r := gin.Default()
	db.OpenConnection()
	routes.Routes(r)
	r.Run()
}

func test(msg string, minute int) {
	time.After(time.Duration(minute))
	println(msg)
}
