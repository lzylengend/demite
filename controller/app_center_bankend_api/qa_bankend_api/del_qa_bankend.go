package qa_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type DelQABankendRequest struct {
	Id int64 `json:"id"`
}

type DelQABankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type DelQABankendApi struct {
}

func (DelQABankendApi) GetRequest() interface{} {
	return &DelQABankendRequest{}
}

func (DelQABankendApi) GetResponse() interface{} {
	return &DelQABankendResponse{}
}

func (DelQABankendApi) GetApi() string {
	return "DelQABankend"
}

func (DelQABankendApi) GetDesc() string {
	return "删除问答"
}

func DelQABankend(c *gin.Context) {
	req := &DelQABankendRequest{}
	rsp := &DelQABankendResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	s, err := model.QADao.Get(req.Id)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	s.DataStatus = time.Now().Unix()

	err = model.QADao.Update(s)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
