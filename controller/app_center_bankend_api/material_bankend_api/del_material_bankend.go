package material_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type DelMaterialBankendRequest struct {
	Id int64 `json:"id"`
}

type DelMaterialBankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type DelMaterialBankendApi struct {
}

func (DelMaterialBankendApi) GetRequest() interface{} {
	return &DelMaterialBankendRequest{}
}

func (DelMaterialBankendApi) GetResponse() interface{} {
	return &DelMaterialBankendResponse{}
}

func (DelMaterialBankendApi) GetApi() string {
	return "DelMaterialBankend"
}

func (DelMaterialBankendApi) GetDesc() string {
	return "删除资料"
}

func DelMaterialBankend(c *gin.Context) {
	req := &DelMaterialBankendRequest{}
	rsp := &DelMaterialBankendResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	vd, err := model.MaterialDao.Get(req.Id)
	if err != nil {
		rsp.Status = my_error.ParamError("id")
		c.JSON(200, rsp)
		return
	}

	vd.UpdateTime = time.Now().Unix()
	vd.DataStatus = time.Now().Unix()

	err = model.MaterialDao.Update(vd)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
