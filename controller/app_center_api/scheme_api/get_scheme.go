package scheme_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type GetSchemeRequest struct {
	Id int64 `json:"id"`
}

type GetSchemeResponse struct {
	Data   *GetSchemeData        `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type GetSchemeData struct {
	Content string `json:"content"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
}

type GetSchemeApi struct {
}

func (GetSchemeApi) GetRequest() interface{} {
	return &GetSchemeRequest{}
}

func (GetSchemeApi) GetResponse() interface{} {
	return &GetSchemeResponse{Data: &GetSchemeData{}}
}

func (GetSchemeApi) GetApi() string {
	return "GetScheme"
}

func (GetSchemeApi) GetDesc() string {
	return "列出"
}

func GetScheme(c *gin.Context) {
	req := &GetSchemeRequest{}
	rsp := &GetSchemeResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	res, err := model.SchemeDao.Get(req.Id)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Data = &GetSchemeData{
		Content: res.Content,
		Desc:    res.Desc,
		Title:   res.Title,
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
