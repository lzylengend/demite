package item_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type GetItemRequest struct {
}

type GetItemResponse struct {
	Title   string                `json:"title"`
	Desc    string                `json:"desc"`
	Content string                `json:"content"`
	Status  *my_error.ErrorCommon `json:"status"`
}

type GetItemApi struct {
}

func (GetItemApi) GetRequest() interface{} {
	return &GetItemRequest{}
}

func (GetItemApi) GetResponse() interface{} {
	return &GetItemResponse{}
}

func (GetItemApi) GetApi() string {
	return "GetItem"
}

func (GetItemApi) GetDesc() string {
	return "获取科学项目"
}

func GetItem(c *gin.Context) {
	req := &GetItemRequest{}
	rsp := &GetItemResponse{}
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

	rsp.Content = in.Content
	rsp.Title = in.Title
	rsp.Desc = in.Desc
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
