package wx_user_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type GetUserInfoRequest struct {
}

type GetUserInfoResponse struct {
	NickName  string                `json:"nickname"`
	Gender    string                `json:"gender"`
	City      string                `json:"city"`
	Province  string                `json:"province"`
	AvatarUrl string                `json:"avatarurl"`
	Country   string                `json:"country"`
	Status    *my_error.ErrorCommon `json:"status"`
}

type GetUserInfoApi struct {
}

func (GetUserInfoApi) GetRequest() interface{} {
	return &GetUserInfoRequest{}
}

func (GetUserInfoApi) GetResponse() interface{} {
	return &GetUserInfoResponse{}
}

func (GetUserInfoApi) GetApi() string {
	return "GetUserInfo"
}

func (GetUserInfoApi) GetDesc() string {
	return "h获取用户信息"
}

func GetUserInfo(c *gin.Context) {
	req := &GetUserInfoRequest{}
	rsp := &GetUserInfoResponse{}
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

	wxuser, exist, err := model.WxUserDao.GetAndExistById(wxId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	if !exist {
		rsp.Status = my_error.NotUserInfoError()
		c.JSON(200, rsp)
		return
	}

	rsp.AvatarUrl = wxuser.AvatarUrl
	rsp.NickName = wxuser.NickName
	rsp.Gender = wxuser.Gender

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
