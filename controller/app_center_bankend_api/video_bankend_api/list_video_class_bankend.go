package video_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListVideoClassBankendRequest struct {
}

type ListVideoClassBankendResponse struct {
	Data   []*ListVideoClassData `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type ListVideoClassData struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type ListVideoClassBankendApi struct {
}

func (ListVideoClassBankendApi) GetRequest() interface{} {
	return &ListVideoClassBankendRequest{}
}

func (ListVideoClassBankendApi) GetResponse() interface{} {
	return &ListVideoClassBankendResponse{Data: []*ListVideoClassData{&ListVideoClassData{}}}
}

func (ListVideoClassBankendApi) GetApi() string {
	return "ListVideoClassBankend"
}

func (ListVideoClassBankendApi) GetDesc() string {
	return "列出视频分类"
}

func ListVideoClassBankend(c *gin.Context) {
	req := &ListVideoClassBankendRequest{}
	rsp := &ListVideoClassBankendResponse{Data: []*ListVideoClassData{}}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	res, err := model.VideoClassDao.List()
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	for _, v := range res {
		rsp.Data = append(rsp.Data, &ListVideoClassData{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
