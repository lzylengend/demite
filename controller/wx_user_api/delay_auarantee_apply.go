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

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
