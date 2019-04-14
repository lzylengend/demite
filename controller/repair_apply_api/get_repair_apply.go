package repair_apply_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type GetRepairApplyRequest struct {
	Id int64 `json:"id"`
}

type GetRepairApplyResponse struct {
	Data   []*repairDetal        `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type repairDetal struct {
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
	return "列出请求"
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

	objList, err := model.RepairScheduleDao.ListByRepairId(req.Id)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	data := make([]*repairDetal, 0)
	for _, v := range objList {
		data = append(data, &repairDetal{
			UserName:   "",
			WxUserName: "",
			StaffName:  "",
			StaffPhone: "",
			RepairTime: v.RepairTime,
			CreateTime: v.CreateTime,
			Status:     string(v.Status),
		})
	}

	rsp.Data = data
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
