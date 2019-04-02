package unlocck_apply_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListApplyRequest struct {
	Limit       int64  `json:"limit"`
	Offset      int64  `json:"offset"`
	GoodName    string `json:"goodname"`
	NickName    string `json:"nickname"`
	ApplyStatus string `json:"applystatus"`
}

type ListApplyResponse struct {
	Data   []*listApplyData      `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type listApplyData struct {
	Id          int64  `json:"id"`
	GoodName    string `json:"goodname"`
	NickName    string `json:"nickname"`
	UserName    string `json:"username"`
	CreateTime  int64  `json:"createtime"`
	UpdateTime  int64  `json:"updatetime"`
	ApplyStatus string `json:"applystatus"`
}

type ListApplyApi struct {
}

func (ListApplyApi) GetRequest() interface{} {
	return &ListApplyRequest{}
}

func (ListApplyApi) GetResponse() interface{} {
	return &ListApplyResponse{}
}

func (ListApplyApi) GetApi() string {
	return "ListApply"
}

func (ListApplyApi) GetDesc() string {
	return "获取列表 applystatus:lock,applying,unlock,refuse"
}

func ListApply(c *gin.Context) {
	req := &ListApplyRequest{}
	rsp := &ListApplyResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	objList, err := model.UnlockApplyDao.ListByGoodUUIdWxUserIdStatus(req.GoodName, req.NickName, req.Limit, req.Offset, req.ApplyStatus)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	for _, v := range objList {
		wxUser, err := model.WxUserDao.GetById(v.WXUserId)
		if err != nil {
			rsp.Status = my_error.DbError(err.Error())
			c.JSON(200, rsp)
			return
		}

		good, err := model.GoodsDao.GetByUUID(v.GoodsUUID)
		if err != nil {
			rsp.Status = my_error.DbError(err.Error())
			c.JSON(200, rsp)
			return
		}

		name := ""
		if v.Creater > 0 {
			user, err := model.UserDao.GetById(v.Creater)
			if err != nil {
				rsp.Status = my_error.DbError(err.Error())
				c.JSON(200, rsp)
				return
			}
			name = user.UserName
		}

		rsp.Data = append(rsp.Data, &listApplyData{
			Id:          v.Id,
			GoodName:    good.GoodsName,
			NickName:    wxUser.NickName,
			CreateTime:  v.CreateTime,
			UpdateTime:  v.UpdateTime,
			UserName:    name,
			ApplyStatus: string(v.Status),
		})
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
