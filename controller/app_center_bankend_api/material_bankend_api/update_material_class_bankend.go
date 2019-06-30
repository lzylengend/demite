package material_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type UpdateMaterialClassBankendRequest struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type UpdateMaterialClassBankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type UpdateMaterialClassBankendApi struct {
}

func (UpdateMaterialClassBankendApi) GetRequest() interface{} {
	return &UpdateMaterialClassBankendRequest{}
}

func (UpdateMaterialClassBankendApi) GetResponse() interface{} {
	return &UpdateMaterialClassBankendResponse{}
}

func (UpdateMaterialClassBankendApi) GetApi() string {
	return "UpdateMaterialClassBankend"
}

func (UpdateMaterialClassBankendApi) GetDesc() string {
	return "修改视频分类"
}

func UpdateMaterialClassBankend(c *gin.Context) {
	req := &UpdateMaterialClassBankendRequest{}
	rsp := &UpdateMaterialClassBankendResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	vc, err := model.MaterialClassDao.Get(req.Id)
	if err != nil {
		rsp.Status = my_error.ParamError("id")
		c.JSON(200, rsp)
		return
	}

	vc.Name = req.Name
	vc.UpdateTime = time.Now().Unix()

	err = model.MaterialClassDao.Update(vc)
	if err != nil {
		rsp.Status = my_error.IdExistError("名字")
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
