package soft_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type UpdateSoftClassBankendRequest struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type UpdateSoftClassBankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type UpdateSoftClassBankendApi struct {
}

func (UpdateSoftClassBankendApi) GetRequest() interface{} {
	return &UpdateSoftClassBankendRequest{}
}

func (UpdateSoftClassBankendApi) GetResponse() interface{} {
	return &UpdateSoftClassBankendResponse{}
}

func (UpdateSoftClassBankendApi) GetApi() string {
	return "UpdateSoftClassBankend"
}

func (UpdateSoftClassBankendApi) GetDesc() string {
	return "修改视频分类"
}

func UpdateSoftClassBankend(c *gin.Context) {
	req := &UpdateSoftClassBankendRequest{}
	rsp := &UpdateSoftClassBankendResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	vc, err := model.SoftClassDao.Get(req.Id)
	if err != nil {
		rsp.Status = my_error.ParamError("id")
		c.JSON(200, rsp)
		return
	}

	vc.Name = req.Name
	vc.UpdateTime = time.Now().Unix()

	err = model.SoftClassDao.Update(vc)
	if err != nil {
		rsp.Status = my_error.IdExistError("名字")
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
