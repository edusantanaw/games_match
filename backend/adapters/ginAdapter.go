package adapters

import (
	httpResponse "github.com/edusantanaw/games_match.git/pkg/utils/httpResponse"
	"github.com/gin-gonic/gin"
)

type IController[T comparable] func(data T) httpResponse.HttpResponse

func GinAdapter[T comparable](controller IController[T]) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var requestBody T
		if err := ctx.BindJSON(&requestBody); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := ctx.BindQuery(&requestBody); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		response := controller(requestBody)
		ctx.JSON(response.StatusCode, response.Body)
	}
}
