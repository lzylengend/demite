package router

import (
	"demite/controller/class_api"
	"demite/controller/file_api"
	"demite/controller/middleware"
	"demite/controller/place_api"
	"demite/controller/product_api"
	"demite/controller/user_api"
	"demite/controller/wx_user_api"
	"github.com/gin-gonic/gin"
)

func Init(g *gin.Engine) {
	manage := g.Group("/manage", middleware.LogReq)
	{
		manage.POST("/login", user_api.Login)
		user := manage.Group("/user", middleware.CheckSession)
		{
			user.POST("/logout", user_api.Logout)
			user.POST("/list", user_api.ListUser)
			user.POST("/add", user_api.AddUser)
			user.POST("/update", user_api.UpdateUser)
			user.POST("/delete", user_api.DeleteUser)
			user.POST("/updatepassword", user_api.UpdatePwd)
		}

		class := manage.Group("/class", middleware.CheckSession)
		{
			class.POST("/add", class_api.AddClass)
			class.POST("/list", class_api.ListClass)
			class.POST("/update", class_api.UpdateClass)
		}

		place := manage.Group("/place", middleware.CheckSession)
		{
			place.POST("/list", place_api.ListPlace)
		}

		file := manage.Group("/file", middleware.CheckSession)
		{
			file.POST("/upload", file_api.UploadFile)
			file.POST("/download", file_api.DownloadFile)
		}

		produce := manage.Group("/product", middleware.CheckSession)
		{
			produce.POST("/add", product_api.AddProduct)
			produce.POST("/list", product_api.ListProduct)
		}
	}

	mini := g.Group("/api", middleware.LogReq)
	{
		mini.POST("/login", wx_user_api.Login)
		//wxUser := mini.Group("/wxuser", middleware.LogReq)
		//{
		//
		//}
	}

}
