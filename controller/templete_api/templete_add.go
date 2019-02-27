package templete_api

import (
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type TempleteAddRequest struct {
}

type TempleteAddResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type TempleteAddApi struct {
}

func (TempleteAddApi) GetRequest() interface{} {
	return &TempleteAddRequest{}
}

func (TempleteAddApi) GetResponse() interface{} {
	return &TempleteAddResponse{}
}

func (TempleteAddApi) GetApi() string {
	return "TempleteAdd"
}

func (TempleteAddApi) GetDesc() string {
	return "新增模板"
}

func TempleteAdd(c *gin.Context) {
	req := &TempleteAddRequest{}
	rsp := &TempleteAddResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}
}
