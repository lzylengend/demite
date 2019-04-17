package wx_user_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type UploadUserInfoRequest struct {
	NickName  string `json:"nickname"`
	Gender    string `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	AvatarUrl string `json:"avatarurl"`
	Country   string `json:"country"`
}

type UploadUserInfoResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type UploadUserInfoApi struct {
}

func (UploadUserInfoApi) GetRequest() interface{} {
	return &UploadUserInfoRequest{}
}

func (UploadUserInfoApi) GetResponse() interface{} {
	return &UploadUserInfoResponse{}
}

func (UploadUserInfoApi) GetApi() string {
	return "UploadUserInfo"
}

func (UploadUserInfoApi) GetDesc() string {
	return "上传用户信息"
}

func UploadUserInfo(c *gin.Context) {
	req := &UploadUserInfoRequest{}
	rsp := &UploadUserInfoResponse{}
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

	wxuser, err := model.WxUserDao.GetById(wxId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	wxuser.NickName = req.NickName
	wxuser.AvatarUrl = req.AvatarUrl
	wxuser.City = req.City
	wxuser.Province = req.Province
	wxuser.Country = req.Country
	wxuser.Gender = req.Gender

	err = model.WxUserDao.Set(wxuser)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
