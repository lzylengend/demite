package scheme_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type UpdateSchemeBankendRequest struct {
	Id      int64  `json:"id"`
	Content string `json:"content"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
}

type UpdateSchemeBankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type UpdateSchemeBankendApi struct {
}

func (UpdateSchemeBankendApi) GetRequest() interface{} {
	return &UpdateSchemeBankendRequest{}
}

func (UpdateSchemeBankendApi) GetResponse() interface{} {
	return &UpdateSchemeBankendResponse{}
}

func (UpdateSchemeBankendApi) GetApi() string {
	return "UpdateSchemeBankend"
}

func (UpdateSchemeBankendApi) GetDesc() string {
	return "修改数据"
}

func UpdateSchemeBankend(c *gin.Context) {
	req := &UpdateSchemeBankendRequest{}
	rsp := &UpdateSchemeBankendResponse{}
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

	s.Content = req.Content
	s.Desc = req.Desc
	s.Title = req.Title

	err = model.SchemeDao.Update(s)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
