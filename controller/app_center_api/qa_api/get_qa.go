package qa_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type GetQARequest struct {
	Id int64 `json:"id"`
}

type GetQAResponse struct {
	Data   *GetQAData            `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type GetQAData struct {
	Content string `json:"content"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
}

type GetQAApi struct {
}

func (GetQAApi) GetRequest() interface{} {
	return &GetQARequest{}
}

func (GetQAApi) GetResponse() interface{} {
	return &GetQAResponse{Data: &GetQAData{}}
}

func (GetQAApi) GetApi() string {
	return "GetQA"
}

func (GetQAApi) GetDesc() string {
	return "列出"
}

func GetQA(c *gin.Context) {
	req := &GetQARequest{}
	rsp := &GetQAResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	res, err := model.QADao.Get(req.Id)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Data = &GetQAData{
		Content: res.Content,
		Desc:    res.Desc,
		Title:   res.Title,
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
