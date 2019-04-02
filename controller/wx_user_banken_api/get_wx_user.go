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
	OpenId    string            `json:"openid"`
	NickName  string            `json:"nickname"`
	Gender    string            `json:"gender"`
	City      string            `json:"city"`
	Province  string            `json:"province"`
	AvatarUrl string            `json:"avatarUrl"`
	Country   string            `json:"country"`
	Data      []*getWxUserGoods `json:"data"`
}

type getWxUserGoods struct {
	Name                    string `json:"name"`
	GoodsUUID               string `json:"goodsuuid"`
	GoodsDecs               string `json:"goodsdecs"`
	GoodsTemplet            string `json:"goodsteplet"`
	GoodsTempletLockContext string `json:"goodstempletlockcontext"`
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
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	gwList, err := model.GoodsWXUserDao.ListByWXId(req.Id, 9999, 0)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	gwData := make([]*getWxUserGoods, 0)
	for _, v := range gwList {
		good, err := model.GoodsDao.GetByUUID(v.GoodsUUID)
		if err != nil {
			rsp.Status = my_error.DbError(err.Error())
			c.JSON(200, rsp)
			return
		}

		gwData = append(gwData, &getWxUserGoods{
			Name:                    good.GoodsName,
			GoodsUUID:               good.GoodsUUID,
			GoodsDecs:               good.GoodsDecs,
			GoodsTemplet:            good.GoodsTemplet,
			GoodsTempletLockContext: good.GoodsTempletLockContext,
		})
	}

	rsp.Data = &getWxData{
		OpenId:    obj.OpenId,
		NickName:  obj.NickName,
		Gender:    obj.Gender,
		City:      obj.City,
		Province:  obj.Province,
		AvatarUrl: obj.AvatarUrl,
		Country:   obj.Country,
		Data:      gwData,
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
