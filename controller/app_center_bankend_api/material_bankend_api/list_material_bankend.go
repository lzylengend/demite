package material_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListMaterialBankendRequest struct {
	ClassId int64 `json:"classid"`
	Limit   int64 `json:"limit"`
	Offset  int64 `json:"offset"`
}

type ListMaterialBankendResponse struct {
	Data   []*ListMaterialData   `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type ListMaterialData struct {
	Id        int64  `json:"id"`
	FileId    string `json:"fileid"`
	ClassId   int64  `json:"classid"`
	ClassName string `json:"classname"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
}

type ListMaterialBankendApi struct {
}

func (ListMaterialBankendApi) GetRequest() interface{} {
	return &ListMaterialBankendRequest{}
}

func (ListMaterialBankendApi) GetResponse() interface{} {
	return &ListMaterialBankendResponse{Data: []*ListMaterialData{&ListMaterialData{}}}
}

func (ListMaterialBankendApi) GetApi() string {
	return "ListMaterialBankend"
}

func (ListMaterialBankendApi) GetDesc() string {
	return "列出视频"
}

func ListMaterialBankend(c *gin.Context) {
	req := &ListMaterialBankendRequest{}
	rsp := &ListMaterialBankendResponse{Data: []*ListMaterialData{}}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	res, err := model.MaterialDao.List(req.ClassId, req.Limit, req.Offset)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	for _, v := range res {
		name := ""

		cla, err := model.MaterialClassDao.Get(v.ClassId)
		if err == nil {
			name = cla.Name
		}

		rsp.Data = append(rsp.Data, &ListMaterialData{
			Id:        v.Id,
			FileId:    v.FileId,
			ClassId:   v.ClassId,
			ClassName: name,
			Desc:      v.Desc,
			Title:     v.Title,
		})
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
