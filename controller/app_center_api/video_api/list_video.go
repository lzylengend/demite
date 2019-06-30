package video_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListVideoRequest struct {
	ClassId int64 `json:"classid"`
	Limit   int64 `json:"limit"`
	Offset  int64 `json:"offset"`
}

type ListVideoResponse struct {
	Data   []*ListVideoData      `json:"data"`
	Count  int64                 `json:"count"`
	Status *my_error.ErrorCommon `json:"status"`
}

type ListVideoData struct {
	Id      int64  `json:"id"`
	FileId  string `json:"fileid"`
	ClassId int64  `json:"classid"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
}

type ListVideoApi struct {
}

func (ListVideoApi) GetRequest() interface{} {
	return &ListVideoRequest{}
}

func (ListVideoApi) GetResponse() interface{} {
	return &ListVideoResponse{Data: []*ListVideoData{&ListVideoData{}}}
}

func (ListVideoApi) GetApi() string {
	return "ListVideo"
}

func (ListVideoApi) GetDesc() string {
	return "列出视频"
}

func ListVideo(c *gin.Context) {
	req := &ListVideoRequest{}
	rsp := &ListVideoResponse{Data: []*ListVideoData{}}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	res, err := model.VideoDao.List(req.ClassId, req.Limit, req.Offset)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	count, err := model.VideoDao.Count(req.ClassId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	for _, v := range res {
		rsp.Data = append(rsp.Data, &ListVideoData{
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
