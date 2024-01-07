package routes

import (
	"fmt"

	"github.com/edusantanaw/games_match.git/adapters"
	"github.com/edusantanaw/games_match.git/internal/controllers"
	authentication "github.com/edusantanaw/games_match.git/pkg/authentication"
	structs "github.com/edusantanaw/games_match.git/pkg/utils/structs"
	"github.com/gin-gonic/gin"
)

func gameRoutes(route *gin.Engine) {
	gamesGroup := route.Group("/game")
	gamesGroup.Use(func(ctx *gin.Context) {
		userId, err := authentication.ExtractTokenMetadata(ctx)
		if err != nil {
			ctx.JSON(401, err.Error())
			return
		}
		fmt.Println(userId)
		ctx.Next()
	})
	gamesGroup.POST("/", adapters.GinAdapter[*controllers.CreateGameData](controllers.RegisterGame))
	route.GET("/game", adapters.GinAdapter[structs.IPagination](controllers.LoadGames))
}
