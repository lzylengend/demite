package video_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type UpdateVideoBankendRequest struct {
	Id       int64  `json:"id"`
	ClassId  int64  `json:"classid"`
	Hot      bool   `json:"hot"`
	Carousel bool   `json:"carousel"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	FileId   string `json:"fileid"`
	PicId    string `json:"picid"`
}

type UpdateVideoBankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type UpdateVideoBankendApi struct {
}

func (UpdateVideoBankendApi) GetRequest() interface{} {
	return &UpdateVideoBankendRequest{}
}

func (UpdateVideoBankendApi) GetResponse() interface{} {
	return &UpdateVideoBankendResponse{}
}

func (UpdateVideoBankendApi) GetApi() string {
	return "UpdateVideoBankend"
}

func (UpdateVideoBankendApi) GetDesc() string {
	return "修改视频"
}

func UpdateVideoBankend(c *gin.Context) {
	req := &UpdateVideoBankendRequest{}
	rsp := &UpdateVideoBankendResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	vd, err := model.VideoDao.Get(req.Id)
	if err != nil {
		rsp.Status = my_error.ParamError("id")
		c.JSON(200, rsp)
		return
	}

	vd.ClassId = req.ClassId
	vd.UpdateTime = time.Now().Unix()
	vd.Hot = req.Hot
	vd.Carousel = req.Carousel
	vd.Title = req.Title
	vd.Desc = req.Desc
	vd.FileId = req.FileId
	vd.PicId = req.PicId

	err = model.VideoDao.Update(vd)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
