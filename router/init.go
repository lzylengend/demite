package router

import (
	"demite/controller/user_api"
	"github.com/gin-gonic/gin"
)

func Init(g *gin.Engine) {
	user := g.Group("/user")
	{
		user.POST("/login", user_api.Login)
	}
}
