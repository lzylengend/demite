package wx_user_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type DelayAuaranteeApplyRequest struct {
	GoodUUID string `json:"gooduuid"`
}

type DelayAuaranteeApplyResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type DelayAuaranteeApplyApi struct {
}

func (DelayAuaranteeApplyApi) GetRequest() interface{} {
	return &DelayAuaranteeApplyRequest{}
}

func (DelayAuaranteeApplyApi) GetResponse() interface{} {
	return &DelayAuaranteeApplyResponse{}
}

func (DelayAuaranteeApplyApi) GetApi() string {
	return "DelayAuaranteeApply"
}

func (DelayAuaranteeApplyApi) GetDesc() string {
	return "申请延保"
}

func DelayAuaranteeApply(c *gin.Context) {
	req := &DelayAuaranteeApplyRequest{}
	rsp := &DelayAuaranteeApplyResponse{}
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

	n, err := model.DelayGuaranteeApplyDao.CountByGoodUUIDWXUserIdStatus(req.GoodUUID, wxId, model.DELAYGUARANTEEAPPLYNG)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	if n > 0 {
		rsp.Status = my_error.ExistApplyError()
		c.JSON(200, rsp)
		return
	}

	_, err = model.DelayGuaranteeApplyDao.Add(req.GoodUUID, wxId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
