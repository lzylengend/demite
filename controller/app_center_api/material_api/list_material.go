package material_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListMaterialRequest struct {
	ClassId int64 `json:"classid"`
	Limit   int64 `json:"limit"`
	Offset  int64 `json:"offset"`
}

type ListMaterialResponse struct {
	Data   []*ListMaterialData   `json:"data"`
	Count  int64                 `json:"count"`
	Status *my_error.ErrorCommon `json:"status"`
}

type ListMaterialData struct {
	Id      int64  `json:"id"`
	FileId  string `json:"fileid"`
	ClassId int64  `json:"classid"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
}

type ListMaterialApi struct {
}

func (ListMaterialApi) GetRequest() interface{} {
	return &ListMaterialRequest{}
}

func (ListMaterialApi) GetResponse() interface{} {
	return &ListMaterialResponse{Data: []*ListMaterialData{&ListMaterialData{}}}
}

func (ListMaterialApi) GetApi() string {
	return "ListMaterial"
}

func (ListMaterialApi) GetDesc() string {
	return "列出资料"
}

func ListMaterial(c *gin.Context) {
	req := &ListMaterialRequest{}
	rsp := &ListMaterialResponse{Data: []*ListMaterialData{}}
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

	count, err := model.MaterialDao.Count(req.ClassId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	for _, v := range res {
		rsp.Data = append(rsp.Data, &ListMaterialData{
			Id:      v.Id,
			FileId:  v.FileId,
			ClassId: v.ClassId,
			Title:   v.Title,
			Desc:    v.Desc,
		})
	}

	rsp.Count = count
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
