package wx_user_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type RepairApplyRequest struct {
	GoodUUID  string `json:"gooduuid"`
	GoodModel string `json:"gooduuid"`
	Phone     string `json:"phone"`
	Name      string `json:"name"`
	Hospital  string `json:"hospital"`
	Office    string `json:"office"`
	Faultdesc string `json:"faultdesc"`
	Faulttype string `json:"faulttype"`
	Fileid1   string `json:"fileid1"`
	Fileid2   string `json:"fileid2"`
}

type RepairApplyResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type RepairApplyApi struct {
}

func (RepairApplyApi) GetRequest() interface{} {
	return &RepairApplyRequest{}
}

func (RepairApplyApi) GetResponse() interface{} {
	return &RepairApplyResponse{}
}

func (RepairApplyApi) GetApi() string {
	return "RepairApply"
}

func (RepairApplyApi) GetDesc() string {
	return "报修延保"
}

func RepairApply(c *gin.Context) {
	req := &RepairApplyRequest{}
	rsp := &RepairApplyResponse{}
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

	exist, _, err := model.GoodsWXUserDao.GetAndExist(req.GoodUUID, wxId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	if !exist {
		rsp.Status = my_error.NotBindError()
		c.JSON(200, rsp)
		return
	}

	_, err = model.RepairDao.Apply(req.GoodUUID, req.GoodModel, req.Phone, req.Name,
		req.Hospital, req.Office, req.Faultdesc, req.Faulttype, req.Fileid1, req.Fileid2, wxId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
