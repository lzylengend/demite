package unlocck_apply_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type DealApplyRequest struct {
	Id    int64 `json:"id"`
	Agree bool  `json:"agree"`
}

type DealApplyResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type DealApplyApi struct {
}

func (DealApplyApi) GetRequest() interface{} {
	return &DealApplyRequest{}
}

func (DealApplyApi) GetResponse() interface{} {
	return &DealApplyResponse{}
}

func (DealApplyApi) GetApi() string {
	return "DealApply"
}

func (DealApplyApi) GetDesc() string {
	return "处理请求 agree 同意为true"
}

func DealApply(c *gin.Context) {
	req := &DealApplyRequest{}
	rsp := &DealApplyResponse{}
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

	err = model.UnlockApplyDao.DealApply(req.Id, req.Agree, userId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
