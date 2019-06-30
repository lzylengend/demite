package soft_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type AddSoftClassBankendRequest struct {
	Name string `json:"name"`
}

type AddSoftClassBankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type AddSoftClassBankendApi struct {
}

func (AddSoftClassBankendApi) GetRequest() interface{} {
	return &AddSoftClassBankendRequest{}
}

func (AddSoftClassBankendApi) GetResponse() interface{} {
	return &AddSoftClassBankendResponse{}
}

func (AddSoftClassBankendApi) GetApi() string {
	return "AddSoftClassBankend"
}

func (AddSoftClassBankendApi) GetDesc() string {
	return "新增资料分类"
}

func AddSoftClassBankend(c *gin.Context) {
	req := &AddSoftClassBankendRequest{}
	rsp := &AddSoftClassBankendResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	err = model.SoftClassDao.Add(&model.SoftClass{
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
