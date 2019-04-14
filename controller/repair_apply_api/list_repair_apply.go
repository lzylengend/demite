package repair_apply_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListRepairRequest struct {
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
	Name   string `json:"name"`
}

type ListRepairResponse struct {
	Data   []*repair             `json:"data"`
	Count  int64                 `json:"count"`
	Status *my_error.ErrorCommon `json:"status"`
}

type repair struct {
	Id        int64  `json:"id"`
	GoodName  string `json:"goodname"`
	GoodModel string `json:"goodmodel"`
	Hospital  string `json:"hospital"`
	Office    string `json:"office"`
	Phone     string `json:"phone"`
	Name      string `json:"name"`
	FaultDesc string `json:"faultdesc"`
	FaultType string `json:"faulttype"`
	FileId1   string `json:"fileid1"`
	FileId2   string `json:"fileid2"`
}

type ListRepairApi struct {
}

func (ListRepairApi) GetRequest() interface{} {
	return &ListRepairRequest{}
}

func (ListRepairApi) GetResponse() interface{} {
	return &ListRepairResponse{}
}

func (ListRepairApi) GetApi() string {
	return "ListRepair"
}

func (ListRepairApi) GetDesc() string {
	return "列出请求"
}

func ListRepair(c *gin.Context) {
	req := &ListRepairRequest{}
	rsp := &ListRepairResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	objList, err := model.RepairDao.List(req.Name, req.Limit, req.Offset)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	count, err := model.RepairDao.Count(req.Name)
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
			Id:        v.RepairId,
			GoodName:  good.GoodsName,
			GoodModel: v.GoodModel,
			Hospital:  v.Hospital,
			Office:    v.Office,
			Phone:     v.Phone,
			Name:      v.Name,
			FaultDesc: v.FaultDesc,
			FaultType: v.FaultType,
			FileId1:   v.FileId1,
			FileId2:   v.FileId2,
		})
	}

	rsp.Data = data
	rsp.Count = count
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
