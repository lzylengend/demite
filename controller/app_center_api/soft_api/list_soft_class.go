package soft_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListSoftClassRequest struct {
}

type ListSoftClassResponse struct {
	Data   []*ListSoftClassData  `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type ListSoftClassData struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type ListSoftClassApi struct {
}

func (ListSoftClassApi) GetRequest() interface{} {
	return &ListSoftClassRequest{}
}

func (ListSoftClassApi) GetResponse() interface{} {
	return &ListSoftClassResponse{Data: []*ListSoftClassData{&ListSoftClassData{}}}
}

func (ListSoftClassApi) GetApi() string {
	return "ListSoftClass"
}

func (ListSoftClassApi) GetDesc() string {
	return "列出资料分类"
}

func ListSoftClass(c *gin.Context) {
	req := &ListSoftClassRequest{}
	rsp := &ListSoftClassResponse{Data: []*ListSoftClassData{}}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	res, err := model.SoftClassDao.List()
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	for _, v := range res {
		rsp.Data = append(rsp.Data, &ListSoftClassData{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
