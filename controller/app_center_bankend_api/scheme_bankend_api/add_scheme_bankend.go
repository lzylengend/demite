package scheme_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type AddSchemeBankendRequest struct {
	Content string `json:"content"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	FileId  string `json:"fileid"`
}

type AddSchemeBankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type AddSchemeBankendApi struct {
}

func (AddSchemeBankendApi) GetRequest() interface{} {
	return &AddSchemeBankendRequest{}
}

func (AddSchemeBankendApi) GetResponse() interface{} {
	return &AddSchemeBankendResponse{}
}

func (AddSchemeBankendApi) GetApi() string {
	return "AddSchemeBankend"
}

func (AddSchemeBankendApi) GetDesc() string {
	return "新增解决方案"
}

func AddSchemeBankend(c *gin.Context) {
	req := &AddSchemeBankendRequest{}
	rsp := &AddSchemeBankendResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	err = model.SchemeDao.Add(&model.Scheme{
		Content:    req.Content,
		DataStatus: 0,
		UpdateTime: time.Now().Unix(),
		CreateTime: time.Now().Unix(),
		Title:      req.Title,
		Desc:       req.Desc,
		FileId:     req.FileId,
	})
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
