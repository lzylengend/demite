package qa_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type UpdateQABankendRequest struct {
	Id      int64  `json:"id"`
	Content string `json:"content"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
}

type UpdateQABankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type UpdateQABankendApi struct {
}

func (UpdateQABankendApi) GetRequest() interface{} {
	return &UpdateQABankendRequest{}
}

func (UpdateQABankendApi) GetResponse() interface{} {
	return &UpdateQABankendResponse{}
}

func (UpdateQABankendApi) GetApi() string {
	return "UpdateQABankend"
}

func (UpdateQABankendApi) GetDesc() string {
	return "修改数据"
}

func UpdateQABankend(c *gin.Context) {
	req := &UpdateQABankendRequest{}
	rsp := &UpdateQABankendResponse{}
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

	s.Content = req.Content
	s.Desc = req.Desc
	s.Title = req.Title

	err = model.QADao.Update(s)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
