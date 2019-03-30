package goods_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type GoodsUpdateRequest struct {
	Id                      int64   `json:"id"`
	Name                    string  `json:"name"`
	GoodsDecs               string  `json:"goodsdecs"`
	GoodsPic                string  `json:"goodspic"`
	DrugList                []int64 `json:"druglist"`
	GoodsTemplet            string  `json:"goodsteplet"`
	GoodsTempletLockContext string  `json:"goodstempletlockcontext"`
}

type GoodsUpdateResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type GoodsUpdateApi struct {
}

func (GoodsUpdateApi) GetRequest() interface{} {
	return &GoodsUpdateRequest{}
}

func (GoodsUpdateApi) GetResponse() interface{} {
	return &GoodsUpdateResponse{}
}

func (GoodsUpdateApi) GetApi() string {
	return "GoodsUpdate"
}

func (GoodsUpdateApi) GetDesc() string {
	return "修改货物"
}

func GoodsUpdate(c *gin.Context) {
	req := &GoodsUpdateRequest{}
	rsp := &GoodsUpdateResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	if req.Name == "" {
		rsp.Status = my_error.NotNilError("name")
		c.JSON(200, rsp)
		return
	}

	g, err := model.GoodsDao.Get(req.Id)
	if err != nil {
		rsp.Status = my_error.ParamError("id")
		c.JSON(200, rsp)
		return
	}

	err = model.GoodDrugsDao.Update(req.DrugList, g.GoodsUUID)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

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
