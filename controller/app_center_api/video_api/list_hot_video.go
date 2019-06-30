package video_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListHotVideoRequest struct {
}

type ListHotVideoResponse struct {
	Data   []*ListHotVideoData   `json:"data"`
	Count  int64                 `json:"count"`
	Status *my_error.ErrorCommon `json:"status"`
}

type ListHotVideoData struct {
	Id     int64  `json:"id"`
	FileId string `json:"fileid"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
}

type ListHotVideoApi struct {
}

func (ListHotVideoApi) GetRequest() interface{} {
	return &ListHotVideoRequest{}
}

func (ListHotVideoApi) GetResponse() interface{} {
	return &ListHotVideoResponse{Data: []*ListHotVideoData{&ListHotVideoData{}}}
}

func (ListHotVideoApi) GetApi() string {
	return "ListHotVideo"
}

func (ListHotVideoApi) GetDesc() string {
	return "列出热门视频"
}

func ListHotVideo(c *gin.Context) {
	req := &ListHotVideoRequest{}
	rsp := &ListHotVideoResponse{Data: []*ListHotVideoData{}}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	res, err := model.VideoDao.ListHot()
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	count, err := model.VideoDao.CountHot()
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	for _, v := range res {
		rsp.Data = append(rsp.Data, &ListHotVideoData{
			Id:     v.Id,
			FileId: v.FileId,
			Title:  v.Title,
			Desc:   v.Desc,
		})
	}

	rsp.Count = count
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
