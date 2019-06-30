package video_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type UpdateVideoClassBankendRequest struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type UpdateVideoClassBankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type UpdateVideoClassBankendApi struct {
}

func (UpdateVideoClassBankendApi) GetRequest() interface{} {
	return &UpdateVideoClassBankendRequest{}
}

func (UpdateVideoClassBankendApi) GetResponse() interface{} {
	return &UpdateVideoClassBankendResponse{}
}

func (UpdateVideoClassBankendApi) GetApi() string {
	return "UpdateVideoClassBankend"
}

func (UpdateVideoClassBankendApi) GetDesc() string {
	return "修改视频分类"
}

func UpdateVideoClassBankend(c *gin.Context) {
	req := &UpdateVideoClassBankendRequest{}
	rsp := &UpdateVideoClassBankendResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	vc, err := model.VideoClassDao.Get(req.Id)
	if err != nil {
		rsp.Status = my_error.ParamError("id")
		c.JSON(200, rsp)
		return
	}

	vc.Name = req.Name
	vc.UpdateTime = time.Now().Unix()

	err = model.VideoClassDao.Update(vc)
	if err != nil {
		rsp.Status = my_error.IdExistError("名字")
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
