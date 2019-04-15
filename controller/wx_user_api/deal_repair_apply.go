package wx_user_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type DealRepairApplyRequest struct {
	Id int64 `json:"id"`
}

type DealRepairApplyResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type DealRepairApplyApi struct {
}

func (DealRepairApplyApi) GetRequest() interface{} {
	return &DealRepairApplyRequest{}
}

func (DealRepairApplyApi) GetResponse() interface{} {
	return &DealRepairApplyResponse{}
}

func (DealRepairApplyApi) GetApi() string {
	return "DealRepairApply"
}

func (DealRepairApplyApi) GetDesc() string {
	return "完成报修延保"
}

func DealRepairApply(c *gin.Context) {
	req := &DealRepairApplyRequest{}
	rsp := &DealRepairApplyResponse{}
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

	err = model.RepairDao.Finish(req.Id, wxId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
