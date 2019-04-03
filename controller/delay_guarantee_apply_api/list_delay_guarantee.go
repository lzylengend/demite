package delay_guarantee_apply_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListDelayGuaranteeApplyRequest struct {
	Limit       int64  `json:"limit"`
	Offset      int64  `json:"offset"`
	GoodName    string `json:"goodname"`
	NickName    string `json:"nickname"`
	ApplyStatus string `json:"applystatus"`
}

type ListDelayGuaranteeApplyResponse struct {
	Data   []*listDelayGuaranteeApplyData `json:"data"`
	Count  int64                          `json:"count"`
	Status *my_error.ErrorCommon          `json:"status"`
}

type listDelayGuaranteeApplyData struct {
	Id          int64  `json:"id"`
	GoodName    string `json:"goodname"`
	NickName    string `json:"nickname"`
	UserName    string `json:"username"`
	CreateTime  int64  `json:"createtime"`
	UpdateTime  int64  `json:"updatetime"`
	ApplyStatus string `json:"applystatus"`
}

type ListDelayGuaranteeApplyApi struct {
}

func (ListDelayGuaranteeApplyApi) GetRequest() interface{} {
	return &ListDelayGuaranteeApplyRequest{}
}

func (ListDelayGuaranteeApplyApi) GetResponse() interface{} {
	return &ListDelayGuaranteeApplyResponse{}
}

func (ListDelayGuaranteeApplyApi) GetApi() string {
	return "ListDelayGuaranteeApply"
}

func (ListDelayGuaranteeApplyApi) GetDesc() string {
	return "获取列表 applystatus:applying,comfirm,refuse"
}

func ListDelayGuaranteeApply(c *gin.Context) {
	req := &ListDelayGuaranteeApplyRequest{}
	rsp := &ListDelayGuaranteeApplyResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	objList, err := model.DelayGuaranteeApplyDao.ListByGoodUUIdWxUserIdStatus(req.GoodName, req.NickName, req.Limit, req.Offset, req.ApplyStatus)
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

		rsp.Data = append(rsp.Data, &listDelayGuaranteeApplyData{
			Id:          v.Id,
			GoodName:    good.GoodsName,
			NickName:    wxUser.NickName,
			CreateTime:  v.CreateTime,
			UpdateTime:  v.UpdateTime,
			UserName:    name,
			ApplyStatus: string(v.Status),
		})
	}

	count, err := model.DelayGuaranteeApplyDao.CountByGoodUUIdWxUserIdStatus(req.GoodName, req.NickName, req.ApplyStatus)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Count = count
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
