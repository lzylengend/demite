package wx_user_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type GetRepairApplyRequest struct {
	Id int64 `json:"id"`
}

type GetRepairApplyResponse struct {
	Data   *repair               `json:"data"`
	List   []*getRepairApplyData `json:"list"`
	Status *my_error.ErrorCommon `json:"status"`
}

type getRepairApplyData struct {
	UserName   string `json:"username"`
	WxUserName string `json:"wxusername"`
	StaffName  string `json:"staffname"`
	StaffPhone string `json:"staffphone"`
	RepairTime int64  `json:"repairtime"`
	CreateTime int64  `json:"createtime"`
	Status     string `json:"currentstatus"`
}

type GetRepairApplyApi struct {
}

func (GetRepairApplyApi) GetRequest() interface{} {
	return &GetRepairApplyRequest{}
}

func (GetRepairApplyApi) GetResponse() interface{} {
	return &GetRepairApplyResponse{}
}

func (GetRepairApplyApi) GetApi() string {
	return "GetRepairApply"
}

func (GetRepairApplyApi) GetDesc() string {
	return "获取报修延保"
}

func GetRepairApply(c *gin.Context) {
	req := &GetRepairApplyRequest{}
	rsp := &GetRepairApplyResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	wxId, err := controller.GetWxUserId(c)
	if err != nil {
		rsp.Status = my_error.NoLoginError()
		c.JSON(200, rsp)
		return
	}

	r, err := model.RepairDao.GetByWIdAndxUserId(wxId, req.Id)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	good, err := model.GoodsDao.GetByUUID(r.GoodUUID)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Data = &repair{
		GoodUUID:   good.GoodsUUID,
		GoodModel:  good.GoodsModel,
		GoodName:   good.GoodsName,
		CreateTime: r.CreateTime,
	}

	objList, err := model.RepairScheduleDao.ListByRepairId(req.Id)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	data := make([]*getRepairApplyData, 0)
	for _, v := range objList {
		data = append(data, &getRepairApplyData{
			UserName:   "",
			WxUserName: "",
			StaffName:  "",
			StaffPhone: "",
			RepairTime: v.RepairTime,
			CreateTime: v.CreateTime,
			Status:     string(v.Status),
		})
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
