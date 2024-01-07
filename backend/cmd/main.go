package main

import (
	"github.com/edusantanaw/games_match.git/cmd/config"
	"github.com/edusantanaw/games_match.git/internal/db"
	"github.com/edusantanaw/games_match.git/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.SetAllEnvs()
	r := gin.Default()
	db.OpenConnection()
	routes.Routes(r)
	r.Run()
}
