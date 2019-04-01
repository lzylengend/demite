package goods_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type GoodsUpdateRequest struct {
	UUId                    string  `json:"uuid"`
	Name                    string  `json:"name"`
	GoodsDecs               string  `json:"goodsdecs"`
	GoodsPic                string  `json:"goodspic"`
	DrugList                []int64 `json:"druglist"`
	GoodsTemplet            string  `json:"goodsteplet"`
	GoodsTempletLockContext string  `json:"goodstempletlockcontext"`
	GoodsModel              string  `json:"goodmodel"`
	GuaranteeTime           int64   `json:"guaranteetime"`
	ClassId                 int64   `json:"classid"`
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

	g, err := model.GoodsDao.GetByUUID(req.UUId)
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

	g.GuaranteeTime = req.GuaranteeTime
	g.GoodsModel = req.GoodsModel
	g.UpdateTime = time.Now().Unix()
	g.ClassId = req.ClassId
	g.GoodsName = req.Name
	g.GoodsDecs = req.GoodsDecs
	g.GoodsPic = req.GoodsPic
	g.GoodsTemplet = req.GoodsTemplet
	g.GoodsTempletLockContext = req.GoodsTempletLockContext

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
