package goods_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type GoodsDelRequest struct {
	UUId string `json:"uuid"`
}

type GoodsDelResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type GoodsDelApi struct {
}

func (GoodsDelApi) GetRequest() interface{} {
	return &GoodsDelRequest{}
}

func (GoodsDelApi) GetResponse() interface{} {
	return &GoodsDelResponse{}
}

func (GoodsDelApi) GetApi() string {
	return "GoodsDel"
}

func (GoodsDelApi) GetDesc() string {
	return "删除货物"
}

func GoodsDel(c *gin.Context) {
	req := &GoodsDelRequest{}
	rsp := &GoodsDelResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	g, err := model.GoodsDao.GetByUUID(req.UUId)
	if err != nil {
		rsp.Status = my_error.ParamError("id")
		c.JSON(200, rsp)
		return
	}

	g.DataStatus = time.Now().Unix()
	err = model.GoodsDao.Set(g)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
