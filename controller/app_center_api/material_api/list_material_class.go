package material_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListMaterialClassRequest struct {
}

type ListMaterialClassResponse struct {
	Data   []*ListMaterialClassData `json:"data"`
	Status *my_error.ErrorCommon    `json:"status"`
}

type ListMaterialClassData struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type ListMaterialClassApi struct {
}

func (ListMaterialClassApi) GetRequest() interface{} {
	return &ListMaterialClassRequest{}
}

func (ListMaterialClassApi) GetResponse() interface{} {
	return &ListMaterialClassResponse{Data: []*ListMaterialClassData{&ListMaterialClassData{}}}
}

func (ListMaterialClassApi) GetApi() string {
	return "ListMaterialClass"
}

func (ListMaterialClassApi) GetDesc() string {
	return "列出资料分类"
}

func ListMaterialClass(c *gin.Context) {
	req := &ListMaterialClassRequest{}
	rsp := &ListMaterialClassResponse{Data: []*ListMaterialClassData{}}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	res, err := model.MaterialClassDao.List()
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	for _, v := range res {
		rsp.Data = append(rsp.Data, &ListMaterialClassData{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
