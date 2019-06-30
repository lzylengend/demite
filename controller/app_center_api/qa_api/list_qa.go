package qa_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListQARequest struct {
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
	Key    string `json:"key"`
}

type ListQAResponse struct {
	Data   []*ListQAData         `json:"data"`
	Count  int64                 `json:"count"`
	Status *my_error.ErrorCommon `json:"status"`
}

type ListQAData struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type ListQAApi struct {
}

func (ListQAApi) GetRequest() interface{} {
	return &ListQARequest{}
}

func (ListQAApi) GetResponse() interface{} {
	return &ListQAResponse{Data: []*ListQAData{&ListQAData{}}}
}

func (ListQAApi) GetApi() string {
	return "ListQA"
}

func (ListQAApi) GetDesc() string {
	return "列出"
}

func ListQA(c *gin.Context) {
	req := &ListQARequest{}
	rsp := &ListQAResponse{Data: []*ListQAData{}}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	res, err := model.QADao.List(0, req.Limit, req.Offset)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	count, err := model.QADao.Count(0)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	for _, v := range res {
		rsp.Data = append(rsp.Data, &ListQAData{
			Id:    v.Id,
			Desc:  v.Desc,
			Title: v.Title,
		})
	}

	rsp.Count = count
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
