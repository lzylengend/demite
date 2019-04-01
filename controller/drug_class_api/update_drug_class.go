package drug_class_api

import (
	"demite/model"
	"demite/my_error"

	"github.com/gin-gonic/gin"
)

type UpdateDrugClassRequest struct {
	ClassId int64  `json:"classid"`
	Name    string `json:"name"`
}

type UpdateDrugClassResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type UpdateDrugClassApi struct {
}

func (UpdateDrugClassApi) GetRequest() interface{} {
	return &UpdateDrugClassRequest{}
}

func (UpdateDrugClassApi) GetResponse() interface{} {
	return &UpdateDrugClassResponse{}
}

func (UpdateDrugClassApi) GetApi() string {
	return "UpdateDrugClass"
}

func (UpdateDrugClassApi) GetDesc() string {
	return "修改药品分类"
}

func UpdateDrugClass(c *gin.Context) {
	req := &UpdateDrugClassRequest{}
	rsp := &UpdateDrugClassResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	if req.Name == "" {
		rsp.Status = my_error.NotNilError("分组名")
		c.JSON(200, rsp)
		return
	}

	class, err := model.DrugClassDao.Get(req.ClassId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	class.ClassName = req.Name

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
