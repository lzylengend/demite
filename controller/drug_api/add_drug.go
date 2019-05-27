package drug_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type AddDrugRequest struct {
	Name                  string `json:"name"`
	ClassId               int64  `json:"classid"`
	Reagent               string `json:"reagent"`               //试剂
	ChromatographicColumn string `json:"chromatographiccolumn"` //色谱柱
	Controls              string `json:"controls"`              //质控品
	TestMethod            string `json:"testmethod"`            //检测方法
	Preprocessing         string `json:"preprocessing"`         //样品预处理
	PotencyRange          string `json:"potencyrange"`          //浓度范围
}

type AddDrugResponse struct {
	Id     int64                 `json:"id"`
	Status *my_error.ErrorCommon `json:"status"`
}

type AddDrugApi struct {
}

func (AddDrugApi) GetRequest() interface{} {
	return &AddDrugRequest{}
}

func (AddDrugApi) GetResponse() interface{} {
	return &AddDrugResponse{}
}

func (AddDrugApi) GetApi() string {
	return "AddDrug"
}

func (AddDrugApi) GetDesc() string {
	return "新增药品  reagent 试剂 chromatographiccolumn 色谱柱 controls 质控品 testmethod  检测方法"
}

func AddDrug(c *gin.Context) {
	req := &AddDrugRequest{}
	rsp := &AddDrugResponse{}
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

	drug, err := model.DrugDao.AddDrug(req.Name, req.ClassId, req.Reagent, req.ChromatographicColumn, req.Controls,
		req.TestMethod, req.Preprocessing, req.PotencyRange)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Id = drug.DrugId
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
