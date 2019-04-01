package wx_user_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"demite/wx_api"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Code string `json:"code"`
}

type LoginResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
	Openid string                `json:"openid"`
}

func Login(c *gin.Context) {
	req := &LoginRequest{}
	rsp := &LoginResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	sessionInfo, err := wx_api.CodeToSession(req.Code)
	if err != nil {
		rsp.Status = my_error.WxError(err.Error())
		c.JSON(200, rsp)
		return
	}

	obj, b, err := model.WxUserDao.ExistOpenid(sessionInfo.Openid)
	if err != nil {
		rsp.Status = my_error.WxError(err.Error())
		c.JSON(200, rsp)
		return
	}

	if b {
		obj.SessionKey = sessionInfo.SessionKey
		err = model.WxUserDao.Set(obj)
		if err != nil {
			rsp.Status = my_error.DbError(err.Error())
			c.JSON(200, rsp)
			return
		}
	} else {
		obj, err = model.WxUserDao.AddWxUser(sessionInfo.Openid, sessionInfo.SessionKey, "", "", "", "", "", "")
		if err != nil {
			rsp.Status = my_error.DbError(err.Error())
			c.JSON(200, rsp)
			return
		}
	}

	session := sessions.Default(c)
	session.Set(controller.SessionWxUserId, obj.WxUserId)
	session.Save()
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
