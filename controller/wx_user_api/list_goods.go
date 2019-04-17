package wx_user_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"

	"github.com/gin-gonic/gin"
)

type ListGoodsRequest struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type ListGoodsResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
	Data   []*goodData           `json:"data"`
	Count  int64                 `json:"count"`
}

type goodData struct {
	UUID   string `json:"uuid"`
	Name   string `json:"name"`
	FileId string `json:"fileid"`
}

type ListGoodsApi struct {
}

func (ListGoodsApi) GetRequest() interface{} {
	return &ListGoodsRequest{}
}

func (ListGoodsApi) GetResponse() interface{} {
	return &ListGoodsResponse{}
}

func (ListGoodsApi) GetApi() string {
	return "ListGoods"
}

func (ListGoodsApi) GetDesc() string {
	return "列出已经绑定的设备"
}

func ListGoods(c *gin.Context) {
	req := &ListGoodsRequest{}
	rsp := &ListGoodsResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	wxId, err := controller.GetWxUserId(c)
	if err != nil {
		rsp.Status = my_error.NoLoginError()
		c.JSON(200, rsp)
		return
	}

	gwObj, err := model.GoodsWXUserDao.ListByWXId(wxId, req.Limit, req.Offset)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	count, err := model.GoodsWXUserDao.CountByWXId(wxId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Count = count

	for _, v := range gwObj {
		good, err := model.GoodsDao.GetByUUID(v.GoodsUUID)
		if err != nil {
			rsp.Status = my_error.DbError(err.Error())
			c.JSON(200, rsp)
			return
		}

		rsp.Data = append(rsp.Data, &goodData{
			UUID:   v.GoodsUUID,
			Name:   good.GoodsName,
			FileId: good.GoodsPic,
		})
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
