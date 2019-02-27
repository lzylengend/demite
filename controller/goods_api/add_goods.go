package goods_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type GoodsAddRequest struct {
	Name                    string `json:"name"`
	GoodsUUID               string `json:"goodsuuid"`
	GoodsDecs               string `json:"goodsdecs"`
	GoodsPic                string `json:"goodspic"`
	GoodsTemplet            string `json:"goodsteplet"`
	GoodsTempletLockContext string `json:"goodstempletlockcontext"`
}

type GoodsAddResponse struct {
	Id     int64                 `json:"id"`
	Status *my_error.ErrorCommon `json:"status"`
}

type GoodsAddApi struct {
}

func (GoodsAddApi) GetRequest() interface{} {
	return &GoodsAddRequest{}
}

func (GoodsAddApi) GetResponse() interface{} {
	return &GoodsAddResponse{}
}

func (GoodsAddApi) GetApi() string {
	return "GoodsAdd"
}

func (GoodsAddApi) GetDesc() string {
	return "新增货物"
}

func GoodsAdd(c *gin.Context) {
	req := &GoodsAddRequest{}
	rsp := &GoodsAddResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	b, err := model.GoodsDao.ExitByUUID(req.GoodsUUID)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	if b {
		rsp.Status = my_error.GoodCodeExistError()
		c.JSON(200, rsp)
		return
	}

	var uId int64 = 0
	uId, err = controller.GetUserId(c)
	if err != nil {
		rsp.Status = my_error.NoLoginError()
		c.JSON(200, rsp)
		return
	}

	g := &model.Goods{
		GoodsUUID:               req.GoodsUUID,
		GoodsCode:               model.GoodsDao.CreateCode(),
		GoodsName:               req.Name,
		GoodsDecs:               req.GoodsDecs,
		GoodsPic:                req.GoodsPic,
		GoodsTemplet:            req.GoodsTemplet,
		GoodsTempletLockContext: req.GoodsTempletLockContext,
		Status:                  model.GOODINIT,
		DataStatus:              0,
		CreateTime:              time.Now().Unix(),
		UpdateTime:              time.Now().Unix(),
		CreatorId:               uId,
	}

	rsp.Id, err = model.GoodsDao.Add(g)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
