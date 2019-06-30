package soft_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type AddSoftBankendRequest struct {
	Content string `json:"content"`
	ClassId int64  `json:"classid"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
}

type AddSoftBankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type AddSoftBankendApi struct {
}

func (AddSoftBankendApi) GetRequest() interface{} {
	return &AddSoftBankendRequest{}
}

func (AddSoftBankendApi) GetResponse() interface{} {
	return &AddSoftBankendResponse{}
}

func (AddSoftBankendApi) GetApi() string {
	return "AddSoftBankend"
}

func (AddSoftBankendApi) GetDesc() string {
	return "新增软件下载"
}

func AddSoftBankend(c *gin.Context) {
	req := &AddSoftBankendRequest{}
	rsp := &AddSoftBankendResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	err = model.SoftDao.Add(&model.Soft{
		Content:    req.Content,
		ClassId:    req.ClassId,
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
