package material_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type AddMaterialBankendRequest struct {
	FileId  string `json:"fileid"`
	ClassId int64  `json:"classid"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	PicId   string `json:"picid"`
}

type AddMaterialBankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type AddMaterialBankendApi struct {
}

func (AddMaterialBankendApi) GetRequest() interface{} {
	return &AddMaterialBankendRequest{}
}

func (AddMaterialBankendApi) GetResponse() interface{} {
	return &AddMaterialBankendResponse{}
}

func (AddMaterialBankendApi) GetApi() string {
	return "AddMaterialBankend"
}

func (AddMaterialBankendApi) GetDesc() string {
	return "新增资料"
}

func AddMaterialBankend(c *gin.Context) {
	req := &AddMaterialBankendRequest{}
	rsp := &AddMaterialBankendResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	err = model.MaterialDao.Add(&model.Material{
		FileId:     req.FileId,
		ClassId:    req.ClassId,
		DataStatus: 0,
		UpdateTime: time.Now().Unix(),
		CreateTime: time.Now().Unix(),
		Title:      req.Title,
		Desc:       req.Desc,
		PicId:      req.PicId,
	})
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
