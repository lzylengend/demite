package intelligence_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type GetInterlligenceRequest struct {
}

type GetInterlligenceResponse struct {
	Title   string                `json:"title"`
	Desc    string                `json:"desc"`
	Content string                `json:"content"`
	Status  *my_error.ErrorCommon `json:"status"`
}

type GetInterlligenceApi struct {
}

func (GetInterlligenceApi) GetRequest() interface{} {
	return &GetInterlligenceRequest{}
}

func (GetInterlligenceApi) GetResponse() interface{} {
	return &GetInterlligenceResponse{}
}

func (GetInterlligenceApi) GetApi() string {
	return "GetInterlligence"
}

func (GetInterlligenceApi) GetDesc() string {
	return "获取资质"
}

func GetInterlligence(c *gin.Context) {
	req := &GetInterlligenceRequest{}
	rsp := &GetInterlligenceResponse{}
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

	rsp.Content = in.Content
	rsp.Title = in.Title
	rsp.Desc = in.Desc
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
