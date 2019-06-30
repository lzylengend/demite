package material_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListMaterialClassBankendRequest struct {
}

type ListMaterialClassBankendResponse struct {
	Data   []*ListMaterialClassData `json:"data"`
	Status *my_error.ErrorCommon    `json:"status"`
}

type ListMaterialClassData struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type ListMaterialClassBankendApi struct {
}

func (ListMaterialClassBankendApi) GetRequest() interface{} {
	return &ListMaterialClassBankendRequest{}
}

func (ListMaterialClassBankendApi) GetResponse() interface{} {
	return &ListMaterialClassBankendResponse{Data: []*ListMaterialClassData{&ListMaterialClassData{}}}
}

func (ListMaterialClassBankendApi) GetApi() string {
	return "ListMaterialClassBankend"
}

func (ListMaterialClassBankendApi) GetDesc() string {
	return "列出视频分类"
}

func ListMaterialClassBankend(c *gin.Context) {
	req := &ListMaterialClassBankendRequest{}
	rsp := &ListMaterialClassBankendResponse{Data: []*ListMaterialClassData{}}
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
