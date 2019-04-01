package drug_class_api

import (
	"demite/model"
	"demite/my_error"

	"github.com/gin-gonic/gin"
)

type DelDrugClassRequest struct {
	ClassId int64 `json:"classid"`
}

type DelDrugClassResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type DelDrugClassApi struct {
}

func (DelDrugClassApi) GetRequest() interface{} {
	return &DelDrugClassRequest{}
}

func (DelDrugClassApi) GetResponse() interface{} {
	return &DelDrugClassResponse{}
}

func (DelDrugClassApi) GetApi() string {
	return "DelDrugClass"
}

func (DelDrugClassApi) GetDesc() string {
	return "删除药品分类"
}

func DelDrugClass(c *gin.Context) {
	req := &DelDrugClassRequest{}
	rsp := &DelDrugClassResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	drugList, err := model.DrugDao.GetByClassId(req.ClassId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	if len(drugList) > 0 {
		rsp.Status = my_error.DrugClassDelError()
		c.JSON(200, rsp)
		return
	}

	class, err := model.DrugClassDao.Get(req.ClassId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	class.DataStatus = 1

	err = model.DrugClassDao.Set(class)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
