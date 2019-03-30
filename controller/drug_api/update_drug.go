package drug_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type UpdateDrugRequest struct {
	Id                    int64  `json:"id"`
	Name                  string `json:"name"`
	ClassId               int64  `json:"classid"`
	Reagent               string `json:"reagent"`               //试剂
	ChromatographicColumn string `json:"chromatographiccolumn"` //色谱柱
	Controls              string `json:"controls"`              //质控品
	TestMethod            string `json:"testmethod"`            //检测方法
}

type UpdateDrugResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type UpdateDrugApi struct {
}

func (UpdateDrugApi) GetRequest() interface{} {
	return &UpdateDrugRequest{}
}

func (UpdateDrugApi) GetResponse() interface{} {
	return &UpdateDrugResponse{}
}

func (UpdateDrugApi) GetApi() string {
	return "UpdateDrug"
}

func (UpdateDrugApi) GetDesc() string {
	return "修改药品  reagent 试剂 chromatographiccolumn 色谱柱 controls 质控品 testmethod  检测方法"
}

func UpdateDrug(c *gin.Context) {
	req := &UpdateDrugRequest{}
	rsp := &UpdateDrugResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	if req.Name == "" {
		rsp.Status = my_error.NotNilError("药品名")
		c.JSON(200, rsp)
		return
	}

	_, err = model.DrugClassDao.Get(req.ClassId)
	if err != nil {
		rsp.Status = my_error.ParamError("classid")
		c.JSON(200, rsp)
		return
	}

	obj, err := model.DrugDao.Get(req.Id)
	if err != nil {
		rsp.Status = my_error.ParamError("id")
		c.JSON(200, rsp)
		return
	}

	obj.DrugName = req.Name
	obj.DrugClassId = req.ClassId
	obj.TestMethod = req.TestMethod
	obj.Controls = req.Controls
	obj.ChromatographicColumn = req.ChromatographicColumn
	obj.Reagent = req.Reagent

	err = model.DrugDao.Set(obj)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
