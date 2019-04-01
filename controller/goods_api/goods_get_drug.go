package goods_api

import (
	"demite/model"
	"demite/my_error"

	"github.com/gin-gonic/gin"
)

type GoodsGetDrugRequest struct {
	GoodsUUID string `json:"goodsuuid"`
}

type GoodsGetDrugResponse struct {
	data   []*drugData           `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type drugData struct {
	Id                    int64  `json:"id"`
	ClassId               int64  `json:"classid"`
	ClassName             string `json:"classname"`
	Name                  string `json:"name"`
	Reagent               string `json:"reagent"`               //试剂
	ChromatographicColumn string `json:"chromatographiccolumn"` //色谱柱
	Controls              string `json:"controls"`              //质控品
	TestMethod            string `json:"testmethod"`            //检测方法
}

type GoodsGetDrugApi struct {
}

func (GoodsGetDrugApi) GetRequest() interface{} {
	return &GoodsGetDrugRequest{}
}

func (GoodsGetDrugApi) GetResponse() interface{} {
	return &GoodsGetDrugResponse{}
}

func (GoodsGetDrugApi) GetApi() string {
	return "GoodsGetDrug"
}

func (GoodsGetDrugApi) GetDesc() string {
	return "货物id获取药品"
}

func GoodsGetDrug(c *gin.Context) {
	req := &GoodsGetDrugRequest{}
	rsp := &GoodsGetDrugResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	objList, err := model.GoodDrugsDao.GetByUUID(req.GoodsUUID)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	for _, v := range objList {
		drug, err := model.DrugDao.Get(v.DrugId)
		if err != nil {
			rsp.Status = my_error.DbError(err.Error())
			c.JSON(200, rsp)
			return
		}

		class, err := model.DrugClassDao.Get(drug.DrugId)
		if err != nil {
			rsp.Status = my_error.DbError(err.Error())
			c.JSON(200, rsp)
			return
		}

		rsp.data = append(rsp.data, &drugData{
			Id:                    v.DrugId,
			ClassId:               drug.DrugClassId,
			ClassName:             class.ClassName,
			Name:                  drug.DrugName,
			Reagent:               drug.Reagent,
			ChromatographicColumn: drug.ChromatographicColumn,
			Controls:              drug.Controls,
			TestMethod:            drug.TestMethod,
		})
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
