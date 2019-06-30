package material_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type AddMaterialClassBankendRequest struct {
	Name string `json:"name"`
}

type AddMaterialClassBankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type AddMaterialClassBankendApi struct {
}

func (AddMaterialClassBankendApi) GetRequest() interface{} {
	return &AddMaterialClassBankendRequest{}
}

func (AddMaterialClassBankendApi) GetResponse() interface{} {
	return &AddMaterialClassBankendResponse{}
}

func (AddMaterialClassBankendApi) GetApi() string {
	return "AddMaterialClassBankend"
}

func (AddMaterialClassBankendApi) GetDesc() string {
	return "新增资料分类"
}

func AddMaterialClassBankend(c *gin.Context) {
	req := &AddMaterialClassBankendRequest{}
	rsp := &AddMaterialClassBankendResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	err = model.MaterialClassDao.Add(&model.MaterialClass{
		Name:       req.Name,
		DataStatus: 0,
		UpdateTime: time.Now().Unix(),
		CreateTime: time.Now().Unix(),
	})
	if err != nil {
		rsp.Status = my_error.IdExistError("名字")
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
