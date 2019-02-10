package user_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"demite/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

type LoginResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
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

	u, err := model.UserDao.GetByName(req.Name)
	if err != nil {
		rsp.Status = my_error.UserNameError(err.Error())
		c.JSON(200, rsp)
		return
	}

	pwd, err := model.UserPasswordDao.GetById(u.UserId)
	if err != nil {
		rsp.Status = my_error.PwdIdError(err.Error())
		c.JSON(200, rsp)
		return
	}

	if util.Md5(model.UserPasswordDao.PwdCombine(req.Pwd, pwd.Salt)) != pwd.Pwd {
		rsp.Status = my_error.PwdFailError("")
		c.JSON(200, rsp)
		return
	}

	session := sessions.Default(c)
	session.Set(controller.SessionUserId, u.UserId)
	session.Save()
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
