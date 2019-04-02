package wx_user_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type UnlockApplyRequest struct {
	GoodUUID string `json:"gooduuid"`
}

type UnlockApplyResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type UnlockApplyApi struct {
}

func (UnlockApplyApi) GetRequest() interface{} {
	return &UnlockApplyRequest{}
}

func (UnlockApplyApi) GetResponse() interface{} {
	return &UnlockApplyResponse{}
}

func (UnlockApplyApi) GetApi() string {
	return "UnlockApply"
}

func (UnlockApplyApi) GetDesc() string {
	return "申请解锁"
}

func UnlockApply(c *gin.Context) {
	req := &UnlockApplyRequest{}
	rsp := &UnlockApplyResponse{}
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

	exist, gw, err := model.GoodsWXUserDao.GetAndExist(req.GoodUUID, wxId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	if !exist {
		rsp.Status = my_error.NotBindError()
		c.JSON(200, rsp)
		return
	}

	exist, _, err = model.UnlockApplyDao.GetByStatusAndExit(req.GoodUUID, wxId, model.GOODSWXUSERAPPLYING)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	if exist {
		rsp.Status = my_error.ExistApplyError()
		c.JSON(200, rsp)
		return
	}

	err = model.UnlockApplyDao.Apply(req.GoodUUID, wxId, gw)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
