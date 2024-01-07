package routes

import (
	"github.com/edusantanaw/games_match.git/adapters"
	"github.com/edusantanaw/games_match.git/controllers"
	"github.com/gin-gonic/gin"
)

func GameRoutes(route *gin.Engine) {
	route.POST("/game", adapters.GinAdapter[*controllers.CreateGameData](controllers.RegisterGame))
}
