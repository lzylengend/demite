package soft_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type DelSoftClassBankendRequest struct {
	Id int64 `json:"id"`
}

type DelSoftClassBankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type DelSoftClassBankendApi struct {
}

func (DelSoftClassBankendApi) GetRequest() interface{} {
	return &DelSoftClassBankendRequest{}
}

func (DelSoftClassBankendApi) GetResponse() interface{} {
	return &DelSoftClassBankendResponse{}
}

func (DelSoftClassBankendApi) GetApi() string {
	return "DelSoftClassBankend"
}

func (DelSoftClassBankendApi) GetDesc() string {
	return "删除资料分类"
}

func DelSoftClassBankend(c *gin.Context) {
	req := &DelSoftClassBankendRequest{}
	rsp := &DelSoftClassBankendResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	vd, err := model.SoftDao.List(req.Id, 100000000, 0)
	if err != nil {
		rsp.Status = my_error.ParamError("id")
		c.JSON(200, rsp)
		return
	}

	if len(vd) > 0 {
		rsp.Status = my_error.DrugClassDelError()
		c.JSON(200, rsp)
		return
	}

	vc, err := model.SoftClassDao.Get(req.Id)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	vc.DataStatus = time.Now().Unix()
	vc.UpdateTime = time.Now().Unix()

	err = model.SoftClassDao.Update(vc)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
