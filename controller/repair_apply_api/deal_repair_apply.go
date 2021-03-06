package repair_apply_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type DealRepairApplyRequest struct {
	Id         int64  `json:"id"`
	StaffId    int64  `json:"staffid"`
	RepairTime int64  `json:"repairtime"`
	Agree      bool   `json:"agree"`
	Reason     string `json:"reason"`
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
	return "处理请求"
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

	userId, err := controller.GetUserId(c)
	if err != nil {
		rsp.Status = my_error.NoLoginError()
		c.JSON(200, rsp)
		return
	}

	if req.Agree {
		_, err = model.StaffDao.Get(req.StaffId)
		if err != nil {
			rsp.Status = my_error.DbError(err.Error())
			c.JSON(200, rsp)
			return
		}

		err = model.RepairDao.Deal(req.Id, userId, req.StaffId, req.RepairTime)
		if err != nil {
			rsp.Status = my_error.DbError(err.Error())
			c.JSON(200, rsp)
			return
		}
	} else {
		err = model.RepairDao.Refuse(req.Id, userId, req.Reason)
		if err != nil {
			rsp.Status = my_error.DbError(err.Error())
			c.JSON(200, rsp)
			return
		}
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
