package drug_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListDrugRequest struct {
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
	Key    string `json:"key"`
}

type ListDrugResponse struct {
	Data   []*drugData           `json:"data"`
	Count  int64                 `json:"count"`
	Status *my_error.ErrorCommon `json:"status"`
}

type drugData struct {
	Id                    int64  `json:"id"`
	ClassName             string `json:"classname"`
	Name                  string `json:"name"`
	Reagent               string `json:"reagent"`               //试剂
	ChromatographicColumn string `json:"chromatographiccolumn"` //色谱柱
	Controls              string `json:"controls"`              //质控品
	TestMethod            string `json:"testmethod"`            //检测方法
}

type ListDrugApi struct {
}

func (ListDrugApi) GetRequest() interface{} {
	return &ListDrugRequest{}
}

func (ListDrugApi) GetResponse() interface{} {
	return &ListDrugResponse{
		Data: []*drugData{
			&drugData{},
		},
	}
}

func (ListDrugApi) GetApi() string {
	return "ListDrug"
}

func (ListDrugApi) GetDesc() string {
	return "列出药品"
}

func ListDrug(c *gin.Context) {
	req := &ListDrugRequest{}
	rsp := &ListDrugResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	data, err := model.DrugDao.ListByCreateTime(req.Key, req.Limit, req.Offset)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	count, err := model.DrugDao.CountByKey(req.Key)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Data = make([]*drugData, 0)
	for _, v := range data {
		class, err := model.DrugClassDao.Get(v.DrugClassId)
		if err != nil {
			rsp.Status = my_error.DbError(err.Error())
			c.JSON(200, rsp)
			return
		}

		rsp.Data = append(rsp.Data, &drugData{
			Name:                  v.DrugName,
			Reagent:               v.Reagent,
			ChromatographicColumn: v.ChromatographicColumn,
			Controls:              v.Controls,
			TestMethod:            v.TestMethod,
			ClassName:             class.ClassName,
		})
	}
	rsp.Count = count
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}