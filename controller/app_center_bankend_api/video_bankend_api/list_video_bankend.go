package video_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListVideoBankendRequest struct {
	ClassId int64 `json:"classid"`
	Limit   int64 `json:"limit"`
	Offset  int64 `json:"offset"`
}

type ListVideoBankendResponse struct {
	Data   []*ListVideoData      `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type ListVideoData struct {
	Id        int64  `json:"id"`
	FileId    string `json:"fileid"`
	ClassId   int64  `json:"classid"`
	ClassName string `json:"classname"`
	Hot       bool   `json:"hot"`
	Carousel  bool   `json:"carousel"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	FileUrl   string `json:"fileurl"`
}

type ListVideoBankendApi struct {
}

func (ListVideoBankendApi) GetRequest() interface{} {
	return &ListVideoBankendRequest{}
}

func (ListVideoBankendApi) GetResponse() interface{} {
	return &ListVideoBankendResponse{Data: []*ListVideoData{&ListVideoData{}}}
}

func (ListVideoBankendApi) GetApi() string {
	return "ListVideoBankend"
}

func (ListVideoBankendApi) GetDesc() string {
	return "列出视频"
}

func ListVideoBankend(c *gin.Context) {
	req := &ListVideoBankendRequest{}
	rsp := &ListVideoBankendResponse{Data: []*ListVideoData{}}
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

	for _, v := range res {
		name := ""
		cla, err := model.VideoClassDao.Get(v.ClassId)
		if err == nil {
			name = cla.Name
		}

		rsp.Data = append(rsp.Data, &ListVideoData{
			Id:        v.Id,
			FileId:    v.FileId,
			ClassId:   v.ClassId,
			ClassName: name,
			Hot:       v.Hot,
			Carousel:  v.Carousel,
			Desc:      v.Desc,
			Title:     v.Title,
		})
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
