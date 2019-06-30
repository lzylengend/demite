package material_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type DelMaterialClassBankendRequest struct {
	Id int64 `json:"id"`
}

type DelMaterialClassBankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type DelMaterialClassBankendApi struct {
}

func (DelMaterialClassBankendApi) GetRequest() interface{} {
	return &DelMaterialClassBankendRequest{}
}

func (DelMaterialClassBankendApi) GetResponse() interface{} {
	return &DelMaterialClassBankendResponse{}
}

func (DelMaterialClassBankendApi) GetApi() string {
	return "DelMaterialClassBankend"
}

func (DelMaterialClassBankendApi) GetDesc() string {
	return "删除资料分类"
}

func DelMaterialClassBankend(c *gin.Context) {
	req := &DelMaterialClassBankendRequest{}
	rsp := &DelMaterialClassBankendResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	vd, err := model.MaterialDao.List(req.Id, 100000000, 0)
	if err != nil {
		rsp.Status = my_error.ParamError("id")
		c.JSON(200, rsp)
		return
	}

	if len(vd) > 0 {
		rsp.Status = my_error.DrugClassDelError()
		c.JSON(200, rsp)
		return
	}

	vc, err := model.MaterialClassDao.Get(req.Id)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	vc.DataStatus = time.Now().Unix()
	vc.UpdateTime = time.Now().Unix()

	err = model.MaterialClassDao.Update(vc)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
