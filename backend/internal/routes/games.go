package routes

import (
	"github.com/edusantanaw/games_match.git/adapters"
	"github.com/edusantanaw/games_match.git/internal/controllers"
	authentication "github.com/edusantanaw/games_match.git/pkg/authentication"
	structs "github.com/edusantanaw/games_match.git/pkg/utils/structs"
	"github.com/gin-gonic/gin"
)

func gameRoutes(route *gin.Engine) {
	gamesGroup := route.Group("/game")
	gamesGroup.Use(AuthMiddleware())
	gamesGroup.POST("/", adapters.GinAdapter[*controllers.CreateGameData](controllers.RegisterGame))
	gamesGroup.GET("", adapters.GinAdapter[structs.IPagination](controllers.LoadGames))
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := authentication.ExtractTokenMetadata(ctx)
		if err != nil {
			println(err.Error())
			ctx.JSON(401, err.Error())
			ctx.Abort()
			return
		}
		println(userId)
		ctx.Next()
	}
}
