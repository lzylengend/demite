package router

import (
	"demite/controller/class_api"
	"demite/controller/drug_api"
	"demite/controller/drug_class_api"
	"demite/controller/file_api"
	"demite/controller/goods_api"
	"demite/controller/middleware"
	"demite/controller/order_api"
	"demite/controller/place_api"
	"demite/controller/user_api"
	"demite/controller/wx_user_api"
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/gin-gonic/gin"
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

	g.GET("test", file_api.Test)

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

		//produce := manage.Group("/product", middleware.CheckSession)
		//{
		//	MyRouterPost(produce, "/add", product_api.AddProductApi{}, product_api.AddProduct)
		//	MyRouterPost(produce, "/list", product_api.ListProductApi{}, product_api.ListProduct)
		//	MyRouterPost(produce, "/update", product_api.UpdateProductApi{}, product_api.UpdateProduct)
		//}

		drugClass := manage.Group("/druclass", middleware.CheckSession)
		{
			MyRouterPost(drugClass, "/add", drug_class_api.AddDrugClassApi{}, drug_class_api.AddDrugClass)
			MyRouterPost(drugClass, "/list", drug_class_api.ListDrugClassApi{}, drug_class_api.ListDrugClass)
			MyRouterPost(drugClass, "/update", drug_class_api.UpdateDrugClassApi{}, drug_class_api.UpdateDrugClass)
			MyRouterPost(drugClass, "/del", drug_class_api.DelDrugClassApi{}, drug_class_api.DelDrugClass)
		}

		drug := manage.Group("/drug", middleware.CheckSession)
		{
			MyRouterPost(drug, "/add", drug_api.AddDrugApi{}, drug_api.AddDrug)
			MyRouterPost(drug, "/list", drug_api.ListDrugApi{}, drug_api.ListDrug)
			MyRouterPost(drug, "/update", drug_api.UpdateDrugApi{}, drug_api.UpdateDrug)
			MyRouterPost(drug, "/del", drug_api.DelDrugApi{}, drug_api.DelDrug)
		}

		goods := manage.Group("/goods", middleware.CheckSession)
		{
			MyRouterPost(goods, "/add", goods_api.GoodsAddApi{}, goods_api.GoodsAdd)
			MyRouterPost(goods, "/list", goods_api.GoodsListApi{}, goods_api.GoodsList)
			MyRouterPost(goods, "/update", goods_api.GoodsUpdateApi{}, goods_api.GoodsUpdate)
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
	doc, err := os.OpenFile("D:/share/doc.txt", os.O_CREATE|os.O_WRONLY, 06667) // /tmp/doc
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
