package drug_class_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListDrugClassRequest struct {
}

type ListDrugClassResponse struct {
	Data   []*drugClassData      `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type drugClassData struct {
	Name string `json:"name"`
}

type ListDrugClassApi struct {
}

func (ListDrugClassApi) GetRequest() interface{} {
	return &ListDrugClassRequest{}
}

func (ListDrugClassApi) GetResponse() interface{} {
	return &ListDrugClassResponse{}
}

func (ListDrugClassApi) GetApi() string {
	return "ListDrugClass"
}

func (ListDrugClassApi) GetDesc() string {
	return "列出药品分类"
}

func ListDrugClass(c *gin.Context) {
	req := &ListDrugClassRequest{}
	rsp := &ListDrugClassResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	res, err := model.DrugClassDao.List()
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	dataList := make([]*drugClassData, 0)
	for _, v := range res {
		dataList = append(dataList, &drugClassData{
			Name: v.ClassName,
		})
	}
	rsp.Data = dataList
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
