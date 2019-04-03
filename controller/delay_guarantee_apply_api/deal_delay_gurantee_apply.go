package delay_guarantee_apply_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type DealDelayGuaranteeApplyRequest struct {
	Id        int64 `json:"id"`
	Agree     bool  `json:"agree"`
	DelayTime int64 `json:"delaytime"`
}

type DealDelayGuaranteeApplyResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type DealDelayGuaranteeApplyApi struct {
}

func (DealDelayGuaranteeApplyApi) GetRequest() interface{} {
	return &DealDelayGuaranteeApplyRequest{}
}

func (DealDelayGuaranteeApplyApi) GetResponse() interface{} {
	return &DealDelayGuaranteeApplyResponse{}
}

func (DealDelayGuaranteeApplyApi) GetApi() string {
	return "DealDelayGuaranteeApply"
}

func (DealDelayGuaranteeApplyApi) GetDesc() string {
	return "处理请求 agree 同意为true"
}

func DealDelayGuaranteeApply(c *gin.Context) {
	req := &DealDelayGuaranteeApplyRequest{}
	rsp := &DealDelayGuaranteeApplyResponse{}
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

	err = model.DelayGuaranteeApplyDao.DealApply(req.Id, req.Agree, userId, req.DelayTime)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
