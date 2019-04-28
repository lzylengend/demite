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
	AuthList *authList             `json:"data"`
	Status   *my_error.ErrorCommon `json:"status"`
}

type authList struct {
	AuthDelGoods     bool `json:"authdelgoods"`
	AuthShieldWxUser bool `json:"authshieldwxuser"`
	AuthUserManage   bool `json:"authusermanage"`
	AuthDelStaff     bool `json:"authdelstaff"`
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

	g, err := model.UserGroupDao.Get(u.UserGroupId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
		return
	}

	aList := &authList{
		AuthDelGoods:     g.AuthDelGoods,
		AuthShieldWxUser: g.AuthShieldWxUser,
		AuthUserManage:   g.AuthUserManage,
		AuthDelStaff:     g.AuthDelStaff,
	}

	session := sessions.Default(c)
	session.Set(controller.SessionUserId, u.UserId)
	session.Save()
	rsp.AuthList = aList
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
