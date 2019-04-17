package wx_user_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListRepairApplyRequest struct {
	GoodUUID string `json:"gooduuid"`
	Limit    int64  `json:"limit"`
	Offset   int64  `json:"offset"`
}

type ListRepairApplyResponse struct {
	Data   []*repair             `json:"data"`
	Count  int64                 `json:"count"`
	Status *my_error.ErrorCommon `json:"status"`
}

type repair struct {
	Id          int64  `json:"id"`
	GoodModel   string `json:"gooduuid"`
	GoodName    string `json:"goodname"`
	ApplyStatus string `json:"applystatus"`
	CreateTime  int64  `json:"createtime"`
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

	objList, err := model.RepairDao.ListByWxUserIdAndGoodUUID(wxId, req.Limit, req.Offset, req.GoodUUID)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	good, err := model.GoodsDao.GetByUUID(req.GoodUUID)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	count, err := model.RepairDao.CountByWxUserIdAndGoodUUID(wxId, req.GoodUUID)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	data := make([]*repair, 0)
	for _, v := range objList {
		data = append(data, &repair{
			GoodModel:   good.GoodsModel,
			GoodName:    good.GoodsName,
			CreateTime:  v.CreateTime,
			ApplyStatus: string(v.Status),
			Id:          v.RepairId,
		})
	}

	rsp.Data = data
	rsp.Count = count
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
