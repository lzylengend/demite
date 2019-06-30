package intelligence_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type GetInterlligenceBankendRequest struct {
}

type GetInterlligenceBankendResponse struct {
	Data   *GetInterlligenceData `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type GetInterlligenceData struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type GetInterlligenceBankendApi struct {
}

func (GetInterlligenceBankendApi) GetRequest() interface{} {
	return &GetInterlligenceBankendRequest{}
}

func (GetInterlligenceBankendApi) GetResponse() interface{} {
	return &GetInterlligenceBankendResponse{}
}

func (GetInterlligenceBankendApi) GetApi() string {
	return "GetInterlligenceBankend"
}

func (GetInterlligenceBankendApi) GetDesc() string {
	return "获取资质"
}

func GetInterlligenceBankend(c *gin.Context) {
	req := &GetInterlligenceBankendRequest{}
	rsp := &GetInterlligenceBankendResponse{}
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

	rsp.Data = &GetInterlligenceData{
		Content: in.Content,
		Title:   in.Title,
		Desc:    in.Desc,
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
