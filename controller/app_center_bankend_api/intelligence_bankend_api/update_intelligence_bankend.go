package intelligence_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type UpdateInterlligenceBankendRequest struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type UpdateInterlligenceBankendResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type UpdateInterlligenceBankendApi struct {
}

func (UpdateInterlligenceBankendApi) GetRequest() interface{} {
	return &UpdateInterlligenceBankendRequest{}
}

func (UpdateInterlligenceBankendApi) GetResponse() interface{} {
	return &UpdateInterlligenceBankendResponse{}
}

func (UpdateInterlligenceBankendApi) GetApi() string {
	return "UpdateInterlligenceBankend"
}

func (UpdateInterlligenceBankendApi) GetDesc() string {
	return "修改资质"
}

func UpdateInterlligenceBankend(c *gin.Context) {
	req := &UpdateInterlligenceBankendRequest{}
	rsp := &UpdateInterlligenceBankendResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	in, err := model.SchemeDao.GetIntelligence()
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	in.Content = req.Content
	in.Title = req.Title
	in.Desc = req.Desc

	err = model.SchemeDao.Update(in)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
