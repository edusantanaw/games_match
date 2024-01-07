package routes

import (
	"github.com/edusantanaw/games_match.git/adapters"
	"github.com/edusantanaw/games_match.git/internal/controllers"
	structs "github.com/edusantanaw/games_match.git/pkg/utils/structs"
	"github.com/gin-gonic/gin"
)

func GameRoutes(route *gin.Engine) {
	route.POST("/game", adapters.GinAdapter[*controllers.CreateGameData](controllers.RegisterGame))
	route.GET("/game", adapters.GinAdapter[structs.IPagination](controllers.LoadGames))
}
