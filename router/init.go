package router

import (
	"demite/controller/class_api"
	"demite/controller/file_api"
	"demite/controller/middleware"
	"demite/controller/order_api"
	"demite/controller/place_api"
	"demite/controller/product_api"
	"demite/controller/user_api"
	"demite/controller/wx_user_api"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"reflect"
)

type MyRouter interface {
	GetRequest() interface{}
	GetResponse() interface{}
	GetApi() string
	GetDesc() string
}

var routerMap map[string]MyRouter

func Init(g *gin.Engine) {
	routerMap = make(map[string]MyRouter)

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
			MyRouterPost(class, "/add", class_api.AddClassApi{}, class_api.AddClass)
			MyRouterPost(class, "/list", class_api.ListClassApi{}, class_api.ListClass)
			MyRouterPost(class, "/update", class_api.UpdateClassApi{}, class_api.UpdateClass)
		}

		place := manage.Group("/place", middleware.CheckSession)
		{
			MyRouterPost(place, "/list", place_api.ListPlaceApi{}, place_api.ListPlace)
		}

		file := manage.Group("/file", middleware.CheckSession)
		{
			MyRouterPost(file, "/upload", file_api.UploadFileApi{}, file_api.UploadFile)
			MyRouterPost(file, "/download", file_api.DownloadFileApi{}, file_api.DownloadFile)
		}

		produce := manage.Group("/product", middleware.CheckSession)
		{
			MyRouterPost(produce, "/add", product_api.AddProductApi{}, product_api.AddProduct)
			MyRouterPost(produce, "/list", product_api.ListProductApi{}, product_api.ListProduct)
			MyRouterPost(produce, "/update", product_api.UpdateProductApi{}, product_api.UpdateProduct)
		}
	}

	mini := g.Group("/api", middleware.LogReq)
	{
		mini.POST("/login", wx_user_api.Login)
		wxUser := mini.Group("/wxuser", middleware.CheckWxSession)
		{
			MyRouterPost(wxUser, "/addorder", order_api.AddOrderApi{}, order_api.AddOrder)
		}
	}
}

func MyRouterPost(group *gin.RouterGroup, path string, r MyRouter, handleFun ...gin.HandlerFunc) {
	group.POST(path, handleFun...)
	routerMap[reflect.TypeOf(r).PkgPath()+"."+r.GetApi()] = r
}

func DoDoc(g *gin.Engine) error {
	doc, err := os.OpenFile("/tmp/doc", os.O_CREATE|os.O_WRONLY, 06667)
	if err != nil {
		return err
	}

	defer doc.Close()

	for _, v := range g.Routes() {
		tmp := ""
		if res, ok := routerMap[string(v.Handler)]; ok {
			req, err := json.Marshal(res.GetRequest())
			if err != nil {
				return err
			}
			rsp, err := json.Marshal(res.GetResponse())
			if err != nil {
				return err
			}

			tmp = fmt.Sprintf("desc : %s\npath : %s\nreq : %s\nresp : %s\n\n", res.GetDesc(), v.Path, string(req), rsp)
			_, err = doc.WriteString(tmp)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
