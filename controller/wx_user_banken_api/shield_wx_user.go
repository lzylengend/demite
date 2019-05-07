package wx_user_banken_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type ShieldWxUserRequest struct {
	Id     int64 `json:"id"`
	Shield bool  `json:"shield"`
}

type ShieldWxUserResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type ShieldWxUserApi struct {
}

func (ShieldWxUserApi) GetRequest() interface{} {
	return &ShieldWxUserRequest{}
}

func (ShieldWxUserApi) GetResponse() interface{} {
	return &ShieldWxUserResponse{}
}

func (ShieldWxUserApi) GetApi() string {
	return "ShieldWxUser"
}

func (ShieldWxUserApi) GetDesc() string {
	return "屏蔽用户"
}

func ShieldWxUser(c *gin.Context) {
	req := &ShieldWxUserRequest{}
	rsp := &ShieldWxUserResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	obj, err := model.WxUserDao.GetById(req.Id)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	var shield int64 = 0
	if !req.Shield {
		shield = time.Now().Unix()
	}

	obj.Shield = shield
	obj.UpdateTime = time.Now().Unix()
	err = model.WxUserDao.Set(obj)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
