package user_api

import (
	"demite/controller"
	"demite/my_error"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LogoutResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

func Logout(c *gin.Context) {
	rsp := &LogoutResponse{}
	session := sessions.Default(c)
	userId := session.Get(controller.SessionUserId)
	if userId == nil {
		rsp.Status = my_error.NoLoginError()
		c.JSON(200, rsp)
		return
	}

	session.Delete(controller.SessionUserId)
	session.Save()

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)

}
