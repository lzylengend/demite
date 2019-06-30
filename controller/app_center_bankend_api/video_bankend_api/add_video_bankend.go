package video_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type AddVideoBankendRequest struct {
	FileId   string `json:"fileid"`
	ClassId  int64  `json:"classid"`
	Hot      bool   `json:"hot"`
	Carousel bool   `json:"carousel"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
}

type AddVideoBankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type AddVideoBankendApi struct {
}

func (AddVideoBankendApi) GetRequest() interface{} {
	return &AddVideoBankendRequest{}
}

func (AddVideoBankendApi) GetResponse() interface{} {
	return &AddVideoBankendResponse{}
}

func (AddVideoBankendApi) GetApi() string {
	return "AddVideoBankend"
}

func (AddVideoBankendApi) GetDesc() string {
	return "新增视频"
}

func AddVideoBankend(c *gin.Context) {
	req := &AddVideoBankendRequest{}
	rsp := &AddVideoBankendResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	err = model.VideoDao.Add(&model.Video{
		FileId:     req.FileId,
		ClassId:    req.ClassId,
		Hot:        req.Hot,
		Carousel:   req.Carousel,
		DataStatus: 0,
		UpdateTime: time.Now().Unix(),
		CreateTime: time.Now().Unix(),
		Title:      req.Title,
		Desc:       req.Desc,
	})
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
