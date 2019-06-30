package qa_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type AddQABankendRequest struct {
	Content string `json:"content"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
}

type AddQABankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type AddQABankendApi struct {
}

func (AddQABankendApi) GetRequest() interface{} {
	return &AddQABankendRequest{}
}

func (AddQABankendApi) GetResponse() interface{} {
	return &AddQABankendResponse{}
}

func (AddQABankendApi) GetApi() string {
	return "AddQABankend"
}

func (AddQABankendApi) GetDesc() string {
	return "新增问答"
}

func AddQABankend(c *gin.Context) {
	req := &AddQABankendRequest{}
	rsp := &AddQABankendResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	err = model.QADao.Add(&model.QA{
		Content:    req.Content,
		DataStatus: 0,
		UpdateTime: time.Now().Unix(),
		CreateTime: time.Now().Unix(),
		Title:      req.Title,
		Desc:       req.Desc,
	})
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
