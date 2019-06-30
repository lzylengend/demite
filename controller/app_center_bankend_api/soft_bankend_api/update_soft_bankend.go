package soft_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type UpdateSoftBankendRequest struct {
	Id      int64  `json:"id"`
	ClassId int64  `json:"classid"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type UpdateSoftBankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type UpdateSoftBankendApi struct {
}

func (UpdateSoftBankendApi) GetRequest() interface{} {
	return &UpdateSoftBankendRequest{}
}

func (UpdateSoftBankendApi) GetResponse() interface{} {
	return &UpdateSoftBankendResponse{}
}

func (UpdateSoftBankendApi) GetApi() string {
	return "UpdateSoftBankend"
}

func (UpdateSoftBankendApi) GetDesc() string {
	return "修改软件下载"
}

func UpdateSoftBankend(c *gin.Context) {
	req := &UpdateSoftBankendRequest{}
	rsp := &UpdateSoftBankendResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	vd, err := model.SoftDao.Get(req.Id)
	if err != nil {
		rsp.Status = my_error.ParamError("id")
		c.JSON(200, rsp)
		return
	}

	vd.ClassId = req.ClassId
	vd.UpdateTime = time.Now().Unix()
	vd.Title = req.Title
	vd.Desc = req.Desc
	vd.Content = req.Content

	err = model.SoftDao.Update(vd)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
