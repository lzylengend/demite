package router

import (
	"demite/controller/delay_guarantee_apply_api"
	"demite/controller/drug_api"
	"demite/controller/drug_class_api"
	"demite/controller/file_api"
	"demite/controller/goods_api"
	"demite/controller/middleware"
	"demite/controller/place_api"
	"demite/controller/remote_apply_api"
	"demite/controller/repair_apply_api"
	"demite/controller/staff_api"
	"demite/controller/unlocck_apply_api"
	"demite/controller/user_api"
	"demite/controller/user_group_api"
	"demite/controller/wx_user_api"
	"demite/controller/wx_user_banken_api"
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

func Init(g *gin.Engine, filePath string) {
	routerMap = make(map[string]MyRouter)

	g.GET("t est", file_api.Test)
	g.Static("/file", filePath)

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

		//class := manage.Group("/class", middleware.CheckSession)
		//{
		//	MyRouterPost(class, "/add", class_api.AddClassApi{}, class_api.AddClass)
		//	MyRouterPost(class, "/list", class_api.ListClassApi{}, class_api.ListClass)
		//	MyRouterPost(class, "/update", class_api.UpdateClassApi{}, class_api.UpdateClass)
		//}

		place := manage.Group("/place", middleware.CheckSession)
		{
			MyRouterPost(place, "/list", place_api.ListPlaceApi{}, place_api.ListPlace)
		}

		file := manage.Group("/file", middleware.CheckSession)
		{
			MyRouterPost(file, "/upload", file_api.UploadFileApi{}, file_api.UploadFile)
			MyRouterPost(file, "/download", file_api.DownloadFileApi{}, file_api.DownloadFile)
		}

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
			MyRouterPost(goods, "/goodgetdrug", goods_api.GoodsGetDrugApi{}, goods_api.GoodsGetDrug)
			MyRouterPost(goods, "/getgood", goods_api.GetGoodApi{}, goods_api.GetGood)
			MyRouterPost(goods, "/del", goods_api.GoodsDelApi{}, middleware.CheckUserAuth, goods_api.GoodsDel)
		}

		wxuUerBanken := manage.Group("/wxuser", middleware.CheckSession)
		{
			MyRouterPost(wxuUerBanken, "/list", wx_user_banken_api.ListWxUserApi{}, wx_user_banken_api.ListWxUser)
			MyRouterPost(wxuUerBanken, "/getwxuser", wx_user_banken_api.GetWxUserApi{}, wx_user_banken_api.GetWxUser)
			MyRouterPost(wxuUerBanken, "/shield", wx_user_banken_api.ShieldWxUserApi{}, wx_user_banken_api.ShieldWxUser)
			MyRouterPost(wxuUerBanken, "/update", wx_user_banken_api.UpdateWxUserApi{}, middleware.CheckUserAuth, wx_user_banken_api.UpdateWxUser)
		}

		unlockApply := manage.Group("/unlockapply", middleware.CheckSession)
		{
			MyRouterPost(unlockApply, "/list", unlocck_apply_api.ListApplyApi{}, unlocck_apply_api.ListApply)
			MyRouterPost(unlockApply, "/dealapply", unlocck_apply_api.DealApplyApi{}, unlocck_apply_api.DealApply)
		}

		delay_guarantee_apply := manage.Group("/delayguaranteeapply", middleware.CheckSession)
		{
			MyRouterPost(delay_guarantee_apply, "/list", delay_guarantee_apply_api.ListDelayGuaranteeApplyApi{}, delay_guarantee_apply_api.ListDelayGuaranteeApply)
			MyRouterPost(delay_guarantee_apply, "/dealapply", delay_guarantee_apply_api.DealDelayGuaranteeApplyApi{}, delay_guarantee_apply_api.DealDelayGuaranteeApply)
		}

		staff := manage.Group("/staff", middleware.CheckSession)
		{
			MyRouterPost(staff, "/add", staff_api.AddStaffApi{}, staff_api.AddStaff)
			MyRouterPost(staff, "/list", staff_api.ListStaffApi{}, staff_api.ListStaff)
			MyRouterPost(staff, "/update", staff_api.UpdateStaffApi{}, staff_api.UpdateStaff)
			MyRouterPost(staff, "/del", staff_api.DelStaffApi{}, middleware.CheckUserAuth, staff_api.DelStaff)
		}

		repairApply := manage.Group("/repairapply", middleware.CheckSession)
		{
			MyRouterPost(repairApply, "/list", repair_apply_api.ListRepairApi{}, repair_apply_api.ListRepair)
			MyRouterPost(repairApply, "/get", repair_apply_api.GetRepairApplyApi{}, repair_apply_api.GetRepairApply)
			MyRouterPost(repairApply, "/deal", repair_apply_api.DealRepairApplyApi{}, repair_apply_api.DealRepairApply)
		}

		remoteApply := manage.Group("/remoteapply", middleware.CheckSession)
		{
			MyRouterPost(remoteApply, "/list", remote_apply_api.ListRemoteApi{}, remote_apply_api.ListRemote)
			MyRouterPost(remoteApply, "/get", remote_apply_api.GetRemoteApplyApi{}, remote_apply_api.GetRemoteApply)
			MyRouterPost(remoteApply, "/deal", remote_apply_api.DealRemoteApplyApi{}, remote_apply_api.DealRemoteApply)
		}

		userGroup := manage.Group("/usergroup", middleware.CheckSession)
		{
			MyRouterPost(userGroup, "/list", user_group_api.ListUserGroupApi{}, user_group_api.ListUserGroup)
			MyRouterPost(userGroup, "/getauth", user_group_api.GetUserAuthApi{}, user_group_api.GetUserAuth)
			//MyRouterPost(userGroup, "/getauth", user_group_api.GetUserAuthApi{}, user_group_api.GetUserAuth)
		}
		//produce := manage.Group("/product", middleware.CheckSession)
		//{
		//	MyRouterPost(produce, "/add", product_api.AddProductApi{}, product_api.AddProduct)
		//	MyRouterPost(produce, "/list", product_api.ListProductApi{}, product_api.ListProduct)
		//	MyRouterPost(produce, "/update", product_api.UpdateProductApi{}, product_api.UpdateProduct)
		//}

	}

	mini := g.Group("/api", middleware.LogReq)
	{
		//MyRouterPost(mini, "/login", wx_user_api.LoginApi{}, wx_user_api.Login)
		mini.POST("/login", wx_user_api.Login)
		wxUser := mini.Group("/wxuser", middleware.CheckWxSession, middleware.CheckWxAuth)
		{
			MyRouterPost(wxUser, "/bindgood", wx_user_api.BindGoodsApi{}, wx_user_api.BindGoods)
			MyRouterPost(wxUser, "/listgoods", wx_user_api.ListGoodsApi{}, wx_user_api.ListGoods)
			MyRouterPost(wxUser, "/getgoods", wx_user_api.GetGoodApi{}, wx_user_api.GetGood)
			MyRouterPost(wxUser, "/listdrugbygood", wx_user_api.ListDrugByGoodsApi{}, wx_user_api.ListDrugByGoods)
			MyRouterPost(wxUser, "/unlockapply", wx_user_api.UnlockApplyApi{}, wx_user_api.UnlockApply)
			MyRouterPost(wxUser, "/delaygauaranteeapply", wx_user_api.DelayAuaranteeApplyApi{}, wx_user_api.DelayAuaranteeApply)
			MyRouterPost(wxUser, "/repairapply", wx_user_api.RepairApplyApi{}, wx_user_api.RepairApply)
			MyRouterPost(wxUser, "/lsitrepairapply", wx_user_api.ListRepairApplyApi{}, wx_user_api.ListRepairApply)
			MyRouterPost(wxUser, "/getrepairapply", wx_user_api.GetRepairApplyApi{}, wx_user_api.GetRepairApply)
			MyRouterPost(wxUser, "/dealrepairapply", wx_user_api.DealRepairApplyApi{}, wx_user_api.DealRepairApply)
			MyRouterPost(wxUser, "/remoteapply", wx_user_api.RemoteApplyApi{}, wx_user_api.RemoteApply)
			MyRouterPost(wxUser, "/lsitremoteapply", wx_user_api.ListRemoteApplyApi{}, wx_user_api.ListRemoteApply)
			MyRouterPost(wxUser, "/getremoteapply", wx_user_api.GetRemoteApplyApi{}, wx_user_api.GetRemoteApply)
			MyRouterPost(wxUser, "/dealremoteapply", wx_user_api.DealRemoteApplyApi{}, wx_user_api.DealRemoteApply)
			MyRouterPost(wxUser, "/uploaduserinfo", wx_user_api.UploadUserInfoApi{}, wx_user_api.UploadUserInfo)
			MyRouterPost(wxUser, "/getduserinfo", wx_user_api.GetUserInfoApi{}, wx_user_api.GetUserInfo)
		}
		wxUserEx := mini.Group("/wxuserex")
		{
			MyRouterPost(wxUserEx, "/uploadfile", wx_user_api.UploadFileApi{}, wx_user_api.UploadFile)
		}
	}
}

func MyRouterPost(group *gin.RouterGroup, path string, r MyRouter, handleFun ...gin.HandlerFunc) {
	group.POST(path, handleFun...)
	routerMap[reflect.TypeOf(r).PkgPath()+"."+r.GetApi()] = r
}

func DoDoc(g *gin.Engine, filePath string) error {
	doc, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 06667) // /tmp/doc
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
