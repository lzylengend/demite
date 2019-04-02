package wx_user_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"

	"github.com/gin-gonic/gin"
)

type BindGoodsRequest struct {
	GoodUUID string `json:"gooduuid"`
}

type BindGoodsResponse struct {
	Status                  *my_error.ErrorCommon `json:"status"`
	Name                    string                `json:"name"`
	GoodsDecs               string                `json:"goodsdecs"`
	GoodsPic                string                `json:"goodspic"`
	GoodsTemplet            string                `json:"goodsteplet"`
	GoodsTempletLockContext string                `json:"goodstempletlockcontext"`
	Lock                    bool                  `json:"lock"`
}

type BindGoodsApi struct {
}

func (BindGoodsApi) GetRequest() interface{} {
	return &BindGoodsRequest{}
}

func (BindGoodsApi) GetResponse() interface{} {
	return &BindGoodsResponse{}
}

func (BindGoodsApi) GetApi() string {
	return "BindGoods"
}

func (BindGoodsApi) GetDesc() string {
	return "绑定"
}

func BindGoods(c *gin.Context) {
	req := &BindGoodsRequest{}
	rsp := &BindGoodsResponse{}
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

	b, gwObj, err := model.GoodsWXUserDao.GetAndExist(req.GoodUUID, wxId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	good, err := model.GoodsDao.GetByUUID(req.GoodUUID)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	if !b {
		err = model.GoodsWXUserDao.Add(req.GoodUUID, wxId)
		if err != nil {
			rsp.Status = my_error.DbError(err.Error())
			c.JSON(200, rsp)
			return
		}
	}

	rsp.GoodsDecs = good.GoodsDecs
	rsp.GoodsPic = good.GoodsPic
	rsp.GoodsTemplet = good.GoodsTemplet
	rsp.Name = good.GoodsName
	rsp.Lock = true
	rsp.GoodsTempletLockContext = ""

	if gwObj.Status == model.GOODSWXUSERUNLOCK {
		rsp.GoodsTempletLockContext = good.GoodsTempletLockContext
		rsp.Lock = false
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
