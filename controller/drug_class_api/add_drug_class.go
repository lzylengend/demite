package drug_class_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type AddDrugClassRequest struct {
	Name      string `json:"name"`
	UpClassId int64  `json:"upclassid"`
}

type AddDrugClassResponse struct {
	ClassId int64                 `json:"classid"`
	Status  *my_error.ErrorCommon `json:"status"`
}

type AddDrugClassApi struct {
}

func (AddDrugClassApi) GetRequest() interface{} {
	return &AddDrugClassRequest{}
}

func (AddDrugClassApi) GetResponse() interface{} {
	return &AddDrugClassResponse{}
}

func (AddDrugClassApi) GetApi() string {
	return "AddDrugClass"
}

func (AddDrugClassApi) GetDesc() string {
	return "新增药品分类"
}

func AddDrugClass(c *gin.Context) {
	req := &AddDrugClassRequest{}
	rsp := &AddDrugClassResponse{}
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

	path := ""
	if req.UpClassId != 0 {
		upClass, err := model.DrugClassDao.Get(req.UpClassId)
		if err != nil {
			rsp.Status = my_error.DbError(err.Error())
			c.JSON(200, rsp)
			return
		}
		path = upClass.Path
	}

	class, err := model.DrugClassDao.AddDrugClass(req.Name, req.UpClassId, path)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.ClassId = class.ClassId
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
