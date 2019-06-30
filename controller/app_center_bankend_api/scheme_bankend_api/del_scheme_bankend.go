package scheme_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type DelSchemeBankendRequest struct {
	Id int64 `json:"id"`
}

type DelSchemeBankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type DelSchemeBankendApi struct {
}

func (DelSchemeBankendApi) GetRequest() interface{} {
	return &DelSchemeBankendRequest{}
}

func (DelSchemeBankendApi) GetResponse() interface{} {
	return &DelSchemeBankendResponse{}
}

func (DelSchemeBankendApi) GetApi() string {
	return "DelSchemeBankend"
}

func (DelSchemeBankendApi) GetDesc() string {
	return "删除数据"
}

func DelSchemeBankend(c *gin.Context) {
	req := &DelSchemeBankendRequest{}
	rsp := &DelSchemeBankendResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	s, err := model.SchemeDao.Get(req.Id)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	s.DataStatus = time.Now().Unix()

	err = model.SchemeDao.Update(s)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
