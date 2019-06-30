package soft_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListSoftClassBankendRequest struct {
}

type ListSoftClassBankendResponse struct {
	Data   []*ListSoftClassData  `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type ListSoftClassData struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type ListSoftClassBankendApi struct {
}

func (ListSoftClassBankendApi) GetRequest() interface{} {
	return &ListSoftClassBankendRequest{}
}

func (ListSoftClassBankendApi) GetResponse() interface{} {
	return &ListSoftClassBankendResponse{Data: []*ListSoftClassData{&ListSoftClassData{}}}
}

func (ListSoftClassBankendApi) GetApi() string {
	return "ListSoftClassBankend"
}

func (ListSoftClassBankendApi) GetDesc() string {
	return "列出视频分类"
}

func ListSoftClassBankend(c *gin.Context) {
	req := &ListSoftClassBankendRequest{}
	rsp := &ListSoftClassBankendResponse{Data: []*ListSoftClassData{}}
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
