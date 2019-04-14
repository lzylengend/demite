package wx_user_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListRepairApplyRequest struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type ListRepairApplyResponse struct {
	Data   []*repair             `json:"data"`
	Count  int64                 `json:"count"`
	Status *my_error.ErrorCommon `json:"status"`
}

type repair struct {
	GoodUUID   string `json:"gooduuid"`
	GoodModel  string `json:"gooduuid"`
	GoodName   string `json:"goodname"`
	CreateTime int64  `json:"createtime"`
}

type ListRepairApplyApi struct {
}

func (ListRepairApplyApi) GetRequest() interface{} {
	return &ListRepairApplyRequest{}
}

func (ListRepairApplyApi) GetResponse() interface{} {
	return &ListRepairApplyResponse{}
}

func (ListRepairApplyApi) GetApi() string {
	return "ListRepairApply"
}

func (ListRepairApplyApi) GetDesc() string {
	return "列出报修延保"
}

func ListRepairApply(c *gin.Context) {
	req := &ListRepairApplyRequest{}
	rsp := &ListRepairApplyResponse{}
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

	objList, err := model.RepairDao.ListByWxUserId(wxId, req.Limit, req.Offset)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	count, err := model.RepairDao.CountByWxUserId(wxId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	data := make([]*repair, 0)
	for _, v := range objList {
		good, err := model.GoodsDao.GetByUUID(v.GoodUUID)
		if err != nil {
			rsp.Status = my_error.DbError(err.Error())
			c.JSON(200, rsp)
			return
		}

		data = append(data, &repair{
			GoodUUID:   good.GoodsUUID,
			GoodModel:  good.GoodsModel,
			GoodName:   good.GoodsName,
			CreateTime: v.CreateTime,
		})
	}

	rsp.Data = data
	rsp.Count = count
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
