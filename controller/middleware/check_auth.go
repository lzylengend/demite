package middleware

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

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
