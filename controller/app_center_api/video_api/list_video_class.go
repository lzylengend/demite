package video_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListVideoClassRequest struct {
}

type ListVideoClassResponse struct {
	Data   []*ListVideoClassData `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type ListVideoClassData struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type ListVideoClassApi struct {
}

func (ListVideoClassApi) GetRequest() interface{} {
	return &ListVideoClassRequest{}
}

func (ListVideoClassApi) GetResponse() interface{} {
	return &ListVideoClassResponse{Data: []*ListVideoClassData{&ListVideoClassData{}}}
}

func (ListVideoClassApi) GetApi() string {
	return "ListVideoClass"
}

func (ListVideoClassApi) GetDesc() string {
	return "列出视频分类"
}

func ListVideoClass(c *gin.Context) {
	req := &ListVideoClassRequest{}
	rsp := &ListVideoClassResponse{Data: []*ListVideoClassData{}}
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
