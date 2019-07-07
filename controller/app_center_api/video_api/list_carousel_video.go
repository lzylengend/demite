package video_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListCarouselVideoRequest struct {
}

type ListCarouselVideoResponse struct {
	Data   []*ListCarouselVideoData `json:"data"`
	Status *my_error.ErrorCommon    `json:"status"`
}

type ListCarouselVideoData struct {
	Id     int64  `json:"id"`
	FileId string `json:"fileid"`
	PicId  string `json:"picid"`
}

type ListCarouselVideoApi struct {
}

func (ListCarouselVideoApi) GetRequest() interface{} {
	return &ListCarouselVideoRequest{}
}

func (ListCarouselVideoApi) GetResponse() interface{} {
	return &ListCarouselVideoResponse{Data: []*ListCarouselVideoData{&ListCarouselVideoData{}}}
}

func (ListCarouselVideoApi) GetApi() string {
	return "ListCarouselVideo"
}

func (ListCarouselVideoApi) GetDesc() string {
	return "列出轮播视频"
}

func ListCarouselVideo(c *gin.Context) {
	req := &ListCarouselVideoRequest{}
	rsp := &ListCarouselVideoResponse{Data: []*ListCarouselVideoData{}}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	res, err := model.VideoDao.ListCarousel()
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	for _, v := range res {
		rsp.Data = append(rsp.Data, &ListCarouselVideoData{
			Id:     v.Id,
			FileId: v.FileId,
			PicId:  v.PicId,
		})
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
