package apply_api

import (
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListApplyRequest struct {
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
	Key    string `json:"key"`
}

type ListApplyResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type ListApplyApi struct {
}

func (ListApplyApi) GetRequest() interface{} {
	return &ListApplyRequest{}
}

func (ListApplyApi) GetResponse() interface{} {
	return &ListApplyResponse{}
}

func (ListApplyApi) GetApi() string {
	return "ListApply"
}

func (ListApplyApi) GetDesc() string {
	return "获取列表"
}

func ListApply(c *gin.Context) {
	req := &ListApplyRequest{}
	rsp := &ListApplyResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
