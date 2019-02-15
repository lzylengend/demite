package router

import (
	"demite/controller/class_api"
	"demite/controller/middleware"
	"demite/controller/user_api"
	"demite/controller/wx_user_api"
	"github.com/gin-gonic/gin"
)

func Init(g *gin.Engine) {
	manage := g.Group("/manage", middleware.LogReq)
	{
		user := manage.Group("/user")
		{
			user.POST("/login", user_api.Login)
			user.POST("/logout", middleware.CheckSession, user_api.Logout)
			user.POST("/list", middleware.CheckSession, user_api.ListUser)
			user.POST("/add", middleware.CheckSession, user_api.AddUser)
			user.POST("/update", middleware.CheckSession, user_api.UpdateUser)
			user.POST("/delete", middleware.CheckSession, user_api.DeleteUser)
			user.POST("/updatepassword", middleware.CheckSession, user_api.UpdatePwd)
		}

		class := manage.Group("/class", middleware.CheckSession)
		{
			class.POST("/add", class_api.AddClass)
			class.POST("/list", class_api.ListClass)
			class.POST("/update", class_api.UpdateClass)
		}
	}

	mini := g.Group("/api", middleware.LogReq)
	{
		wxUser := mini.Group("/wxuser", middleware.LogReq)
		{
			wxUser.POST("/login", wx_user_api.Login)
		}
	}

}
