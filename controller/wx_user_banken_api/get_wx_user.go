package wx_user_banken_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type GetWxUserRequest struct {
	Id int64 `json:"id"`
}

type GetWxUserResponse struct {
	Data   *getWxData            `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type getWxData struct {
	Id        int64  `json:"id"`
	OpenId    string `json:"openid"`
	NickName  string `json:"nickname"`
	Gender    string `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	AvatarUrl string `json:"avatarUrl"`
	Country   string `json:"country"`
}

type GetWxUserApi struct {
}

func (GetWxUserApi) GetRequest() interface{} {
	return &GetWxUserRequest{}
}

func (GetWxUserApi) GetResponse() interface{} {
	return &GetWxUserResponse{
		Data: &getWxData{},
	}
}

func (GetWxUserApi) GetApi() string {
	return "GetWxUser"
}

func (GetWxUserApi) GetDesc() string {
	return "获取关注用户"
}

func GetWxUser(c *gin.Context) {
	req := &GetWxUserRequest{}
	rsp := &GetWxUserResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	obj, err := model.WxUserDao.GetById(req.Id)

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
