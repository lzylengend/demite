package item_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type UpdateItemBankendRequest struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type UpdateItemBankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type UpdateItemBankendApi struct {
}

func (UpdateItemBankendApi) GetRequest() interface{} {
	return &UpdateItemBankendRequest{}
}

func (UpdateItemBankendApi) GetResponse() interface{} {
	return &UpdateItemBankendResponse{}
}

func (UpdateItemBankendApi) GetApi() string {
	return "UpdateItemBankend"
}

func (UpdateItemBankendApi) GetDesc() string {
	return "修改科学"
}

func UpdateItemBankend(c *gin.Context) {
	req := &UpdateItemBankendRequest{}
	rsp := &UpdateItemBankendResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	in, err := model.SchemeDao.GetItem()
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	in.Content = req.Content
	in.Title = req.Title
	in.Desc = req.Desc

	err = model.SchemeDao.Update(in)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
