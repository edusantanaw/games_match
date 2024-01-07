package routes

import (
	"github.com/edusantanaw/games_match.git/adapters"
	"github.com/edusantanaw/games_match.git/internal/controllers"
	"github.com/gin-gonic/gin"
)

func usersRoute(c *gin.Engine) {
	usersGroup := c.Group("/users")
	usersGroup.POST("/signup", adapters.GinAdapter[*controllers.RegisterUserData](controllers.ResgisterUser))
}
