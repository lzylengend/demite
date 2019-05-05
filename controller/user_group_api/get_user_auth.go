package user_group_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"

	"github.com/gin-gonic/gin"
)

type GetUserAuthRequest struct {
}

type GetUserAuthResponse struct {
	Data   *auth                 `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type auth struct {
	AuthDelGoods     bool `json:"authdelgoods"`
	AuthShieldWxUser bool `json:"authshieldwxuser"`
	AuthUserManage   bool `json:"authusermanage"`
	AuthDelStaff     bool `json:"authdelstaff"`
}

type GetUserAuthApi struct {
}

func (GetUserAuthApi) GetRequest() interface{} {
	return &GetUserAuthRequest{}
}

func (GetUserAuthApi) GetResponse() interface{} {
	return &GetUserAuthResponse{}
}

func (GetUserAuthApi) GetApi() string {
	return "GetUserAuth"
}

func (GetUserAuthApi) GetDesc() string {
	return "获取用户权限"
}

func GetUserAuth(c *gin.Context) {
	req := &GetUserAuthRequest{}
	rsp := &GetUserAuthResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	id, err := controller.GetUserId(c)
	if err != nil {
		rsp.Status = my_error.NoLoginError()
		c.JSON(200, rsp)
		return
	}

	u, err := model.UserDao.GetById(id)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	g, err := model.UserGroupDao.Get(u.UserGroupId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	a := &auth{
		AuthDelGoods:     g.AuthDelGoods,
		AuthShieldWxUser: g.AuthShieldWxUser,
		AuthUserManage:   g.AuthUserManage,
		AuthDelStaff:     g.AuthDelStaff,
	}

	rsp.Data = a
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
