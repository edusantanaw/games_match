package routes

import "github.com/gin-gonic/gin"

func Routes(route *gin.Engine) {
	gameRoutes(route)
	usersRoute(route)
}
