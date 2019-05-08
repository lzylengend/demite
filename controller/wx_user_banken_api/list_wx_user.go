package wx_user_banken_api

import (
	"demite/model"
	"demite/my_error"

	"github.com/gin-gonic/gin"
)

type ListWxUserRequest struct {
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
	Key    string `json:"key"`
}

type ListWxUserResponse struct {
	Data   []*wxData             `json:"data"`
	Count  int64                 `json:"count"`
	Status *my_error.ErrorCommon `json:"status"`
}

type wxData struct {
	Id        int64  `json:"id"`
	OpenId    string `json:"openid"`
	NickName  string `json:"nickname"`
	Gender    string `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	AvatarUrl string `json:"avatarUrl"`
	Country   string `json:"country"`
	Shield    bool   `json:"shield"`
}

type ListWxUserApi struct {
}

func (ListWxUserApi) GetRequest() interface{} {
	return &ListWxUserRequest{}
}

func (ListWxUserApi) GetResponse() interface{} {
	return &ListWxUserResponse{
		Data: []*wxData{
			&wxData{},
		},
	}
}

func (ListWxUserApi) GetApi() string {
	return "ListWxUser"
}

func (ListWxUserApi) GetDesc() string {
	return "列出已经关注用户"
}

func ListWxUser(c *gin.Context) {
	req := &ListWxUserRequest{}
	rsp := &ListWxUserResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	objList, err := model.WxUserDao.List(req.Key, req.Limit, req.Offset)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	count, err := model.WxUserDao.Count(req.Key)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	for _, v := range objList {
		rsp.Data = append(rsp.Data, &wxData{
			Id:        v.WxUserId,
			OpenId:    v.OpenId,
			NickName:  v.NickName,
			Gender:    v.Gender,
			City:      v.City,
			Province:  v.Province,
			AvatarUrl: v.AvatarUrl,
			Country:   v.Country,
			Shield:    v.Shield != 0,
		})
	}

	rsp.Count = count
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
