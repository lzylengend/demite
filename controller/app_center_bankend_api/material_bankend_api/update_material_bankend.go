package material_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type UpdateMaterialBankendRequest struct {
	Id      int64  `json:"id"`
	ClassId int64  `json:"classid"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	FileId  string `json:"fileid"`
	PicId   string `json:"picid"`
}

type UpdateMaterialBankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type UpdateMaterialBankendApi struct {
}

func (UpdateMaterialBankendApi) GetRequest() interface{} {
	return &UpdateMaterialBankendRequest{}
}

func (UpdateMaterialBankendApi) GetResponse() interface{} {
	return &UpdateMaterialBankendResponse{}
}

func (UpdateMaterialBankendApi) GetApi() string {
	return "UpdateMaterialBankend"
}

func (UpdateMaterialBankendApi) GetDesc() string {
	return "修改视频"
}

func UpdateMaterialBankend(c *gin.Context) {
	req := &UpdateMaterialBankendRequest{}
	rsp := &UpdateMaterialBankendResponse{}
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

	vd.ClassId = req.ClassId
	vd.UpdateTime = time.Now().Unix()
	vd.Title = req.Title
	vd.Desc = req.Desc
	vd.FileId = req.FileId
	vd.PicId = req.PicId

	err = model.MaterialDao.Update(vd)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
