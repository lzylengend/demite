package middleware

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

func CheckUserAuth(c *gin.Context) {
	rsp := &commonRespose{}
	id, err := controller.GetUserId(c)
	if err != nil {
		rsp.Status = my_error.NoLoginError()
		c.JSON(200, rsp)
		c.Abort()
		return
	}

	u, err := model.UserDao.GetById(id)
	if err != nil {
		rsp.Status = my_error.NoLoginError()
		c.JSON(200, rsp)
		c.Abort()
		return
	}

	if u.UserGroupId != 1 {
		rsp.Status = my_error.NoAuthError()
		c.JSON(200, rsp)
		c.Abort()
		return
	}

	c.Next()
}

func CheckWxAuth(c *gin.Context) {
	rsp := &commonRespose{}
	id, err := controller.GetWxUserId(c)

	if err != nil {
		rsp.Status = my_error.NoLoginError()
		c.JSON(200, rsp)
		c.Abort()
		return
	}

	wx, err := model.WxUserDao.GetById(id)
	if err != nil {
		rsp.Status = my_error.NoLoginError()
		c.JSON(200, rsp)
		c.Abort()
		return
	}

	if wx.Shield != 0 {
		rsp.Status = my_error.NoAuthError()
		c.JSON(200, rsp)
		c.Abort()
		return
	}

	c.Next()
}
