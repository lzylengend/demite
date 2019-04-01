package wx_user_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"

	"github.com/gin-gonic/gin"
)

type ListGoodsRequest struct {
}

type ListGoodsResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
	Data   []*goodData           `json:"data"`
}

type goodData struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
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

	gwObj, err := model.GoodsWXUserDao.ListByWXId(wxId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	for _, v := range gwObj {
		good, err := model.GoodsDao.GetByUUID(v.GoodsUUID)
		if err != nil {
			rsp.Status = my_error.DbError(err.Error())
			c.JSON(200, rsp)
			return
		}

		rsp.Data = append(rsp.Data, &goodData{
			UUID: v.GoodsUUID,
			Name: good.GoodsName,
		})
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
