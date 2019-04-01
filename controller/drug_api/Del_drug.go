package drug_api

import (
	"demite/model"
	"demite/my_error"

	"github.com/gin-gonic/gin"
)

type DelDrugRequest struct {
	Id int64 `json:"id"`
}

type DelDrugResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type DelDrugApi struct {
}

func (DelDrugApi) GetRequest() interface{} {
	return &DelDrugRequest{}
}

func (DelDrugApi) GetResponse() interface{} {
	return &DelDrugResponse{}
}

func (DelDrugApi) GetApi() string {
	return "DelDrug"
}

func (DelDrugApi) GetDesc() string {
	return "删除药品"
}

func DelDrug(c *gin.Context) {
	req := &DelDrugRequest{}
	rsp := &DelDrugResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	obj, err := model.DrugDao.Get(req.Id)
	if err != nil {
		rsp.Status = my_error.ParamError("id")
		c.JSON(200, rsp)
		return
	}

	obj.DataStatus = 1

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
