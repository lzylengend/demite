package video_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type DelVideoClassBankendRequest struct {
	Id int64 `json:"id"`
}

type DelVideoClassBankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type DelVideoClassBankendApi struct {
}

func (DelVideoClassBankendApi) GetRequest() interface{} {
	return &DelVideoClassBankendRequest{}
}

func (DelVideoClassBankendApi) GetResponse() interface{} {
	return &DelVideoClassBankendResponse{}
}

func (DelVideoClassBankendApi) GetApi() string {
	return "DelVideoClassBankend"
}

func (DelVideoClassBankendApi) GetDesc() string {
	return "删除视频分类"
}

func DelVideoClassBankend(c *gin.Context) {
	req := &DelVideoClassBankendRequest{}
	rsp := &DelVideoClassBankendResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	vd, err := model.VideoDao.List(req.Id, 100000000, 0)
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

	vc, err := model.VideoClassDao.Get(req.Id)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	vc.DataStatus = time.Now().Unix()
	vc.UpdateTime = time.Now().Unix()

	err = model.VideoClassDao.Update(vc)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
