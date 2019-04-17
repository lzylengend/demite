package remote_apply_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type DealRemoteApplyRequest struct {
	Id       int64 `json:"id"`
	StaffId  int64 `json:"staffid"`
	DealTime int64 `json:"dealtime"`
}

type DealRemoteApplyResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type DealRemoteApplyApi struct {
}

func (DealRemoteApplyApi) GetRequest() interface{} {
	return &DealRemoteApplyRequest{}
}

func (DealRemoteApplyApi) GetResponse() interface{} {
	return &DealRemoteApplyResponse{}
}

func (DealRemoteApplyApi) GetApi() string {
	return "DealRemoteApply"
}

func (DealRemoteApplyApi) GetDesc() string {
	return "处理请求"
}

func DealRemoteApply(c *gin.Context) {
	req := &DealRemoteApplyRequest{}
	rsp := &DealRemoteApplyResponse{}
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

	_, err = model.StaffDao.Get(req.StaffId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	err = model.RemoteDao.Deal(req.Id, userId, req.StaffId, req.DealTime)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
