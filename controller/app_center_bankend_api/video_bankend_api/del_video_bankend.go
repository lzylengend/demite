package video_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type DelVideoBankendRequest struct {
	Id int64 `json:"id"`
}

type DelVideoBankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type DelVideoBankendApi struct {
}

func (DelVideoBankendApi) GetRequest() interface{} {
	return &DelVideoBankendRequest{}
}

func (DelVideoBankendApi) GetResponse() interface{} {
	return &DelVideoBankendResponse{}
}

func (DelVideoBankendApi) GetApi() string {
	return "DelVideoBankend"
}

func (DelVideoBankendApi) GetDesc() string {
	return "删除视频"
}

func DelVideoBankend(c *gin.Context) {
	req := &DelVideoBankendRequest{}
	rsp := &DelVideoBankendResponse{}
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

	vd.UpdateTime = time.Now().Unix()
	vd.DataStatus = time.Now().Unix()

	err = model.VideoDao.Update(vd)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
