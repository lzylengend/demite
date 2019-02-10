package router

import (
	"demite/controller/middleware"
	"demite/controller/user_api"
	"github.com/gin-gonic/gin"
)

func Init(g *gin.Engine) {
	user := g.Group("/user")
	{
		user.POST("/login", user_api.Login)
		user.POST("/logout", middleware.CheckSession, user_api.Logout)
		user.POST("/list", middleware.CheckSession, user_api.ListUser)
		user.POST("/add", middleware.CheckSession, user_api.AddUser)
		user.POST("/update", middleware.CheckSession, user_api.UpdateUser)
		user.POST("/delete", middleware.CheckSession, user_api.DeleteUser)
		user.POST("/updatepassword", middleware.CheckSession, user_api.UpdatePwd)
	}
}
