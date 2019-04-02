package wx_user_api

import (
	"demite/conf"
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type GetGoodRequest struct {
	GoodUUID string `json:"gooduuid"`
}

type GetGoodResponse struct {
	Status                  *my_error.ErrorCommon `json:"status"`
	Name                    string                `json:"name"`
	GoodsDecs               string                `json:"goodsdecs"`
	GoodsPic                string                `json:"goodspic"`
	GoodsTemplet            string                `json:"goodsteplet"`
	GoodsTempletLockContext string                `json:"goodstempletlockcontext"`
	LockStatus              string                `json:"lockstatus"`
	GoodsPicData            string                `json:"goodpicdata"`
}

type GetGoodApi struct {
}

func (GetGoodApi) GetRequest() interface{} {
	return &GetGoodRequest{}
}

func (GetGoodApi) GetResponse() interface{} {
	return &GetGoodResponse{}
}

func (GetGoodApi) GetApi() string {
	return "GetGood"
}

func (GetGoodApi) GetDesc() string {
	return "获取设备详情"
}

func GetGood(c *gin.Context) {
	req := &GetGoodRequest{}
	rsp := &GetGoodResponse{}
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
		rsp.Status = my_error.NotBindError()
		c.JSON(200, rsp)
		return
	}

	data, err := ioutil.ReadFile(conf.GetFilePath() + "/" + good.GoodsPic)
	if err != nil {
		data = []byte{}
	}

	rsp.GoodsDecs = good.GoodsDecs
	rsp.GoodsPic = good.GoodsPic
	rsp.GoodsTemplet = good.GoodsTemplet
	rsp.Name = good.GoodsName
	rsp.GoodsTempletLockContext = ""
	rsp.GoodsPicData = base64.StdEncoding.EncodeToString(data)
	rsp.LockStatus = string(gwObj.Status)

	if gwObj.Status == model.GOODSWXUSERUNLOCK {
		rsp.GoodsTempletLockContext = good.GoodsTempletLockContext
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
