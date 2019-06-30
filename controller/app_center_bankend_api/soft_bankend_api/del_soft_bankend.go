package soft_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type DelSoftBankendRequest struct {
	Id int64 `json:"id"`
}

type DelSoftBankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type DelSoftBankendApi struct {
}

func (DelSoftBankendApi) GetRequest() interface{} {
	return &DelSoftBankendRequest{}
}

func (DelSoftBankendApi) GetResponse() interface{} {
	return &DelSoftBankendResponse{}
}

func (DelSoftBankendApi) GetApi() string {
	return "DelSoftBankend"
}

func (DelSoftBankendApi) GetDesc() string {
	return "删除资料"
}

func DelSoftBankend(c *gin.Context) {
	req := &DelSoftBankendRequest{}
	rsp := &DelSoftBankendResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	vd, err := model.SoftDao.Get(req.Id)
	if err != nil {
		rsp.Status = my_error.ParamError("id")
		c.JSON(200, rsp)
		return
	}

	vd.UpdateTime = time.Now().Unix()
	vd.DataStatus = time.Now().Unix()

	err = model.SoftDao.Update(vd)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
