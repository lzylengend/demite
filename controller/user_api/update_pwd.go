package user_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"demite/util"
	"github.com/gin-gonic/gin"
)

type UpdatePwdRequest struct {
	Pwd string `json:"pwd"`
}

type UpdatePwdResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

func UpdatePwd(c *gin.Context) {
	req := &UpdatePwdRequest{}
	rsp := &UpdatePwdResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	uid, err := controller.GetUserId(c)
	if err != nil {
		rsp.Status = my_error.NoLoginError()
		c.JSON(200, rsp)
		return
	}

	p, err := model.UserPasswordDao.GetById(uid)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	p.Pwd = util.Md5(model.UserPasswordDao.PwdCombine(req.Pwd, p.Salt))
	err = model.UserPasswordDao.Set(p)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
