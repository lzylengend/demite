package item_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type GetItemBankendRequest struct {
}

type GetItemBankendResponse struct {
	Data   *GetItemData          `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type GetItemData struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type GetItemBankendApi struct {
}

func (GetItemBankendApi) GetRequest() interface{} {
	return &GetItemBankendRequest{}
}

func (GetItemBankendApi) GetResponse() interface{} {
	return &GetItemBankendResponse{}
}

func (GetItemBankendApi) GetApi() string {
	return "GetItemBankend"
}

func (GetItemBankendApi) GetDesc() string {
	return "获取科学项目"
}

func GetItemBankend(c *gin.Context) {
	req := &GetItemBankendRequest{}
	rsp := &GetItemBankendResponse{}
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

	rsp.Data = &GetItemData{
		Content: in.Content,
		Title:   in.Title,
		Desc:    in.Desc,
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
