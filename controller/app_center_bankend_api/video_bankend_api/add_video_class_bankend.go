package video_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type AddVideoClassBankendRequest struct {
	Name string `json:"name"`
}

type AddVideoClassBankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type AddVideoClassBankendApi struct {
}

func (AddVideoClassBankendApi) GetRequest() interface{} {
	return &AddVideoClassBankendRequest{}
}

func (AddVideoClassBankendApi) GetResponse() interface{} {
	return &AddVideoClassBankendResponse{}
}

func (AddVideoClassBankendApi) GetApi() string {
	return "AddVideoClassBankend"
}

func (AddVideoClassBankendApi) GetDesc() string {
	return "新增视频分类"
}

func AddVideoClassBankend(c *gin.Context) {
	req := &AddVideoClassBankendRequest{}
	rsp := &AddVideoClassBankendResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	err = model.VideoClassDao.Add(&model.VideoClass{
		Name:       req.Name,
		DataStatus: 0,
		UpdateTime: time.Now().Unix(),
		CreateTime: time.Now().Unix(),
	})
	if err != nil {
		rsp.Status = my_error.IdExistError("名字")
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
