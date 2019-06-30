package qa_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListQABankendRequest struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type ListQABankendResponse struct {
	Data   []*ListQAData         `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type ListQAData struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type ListQABankendApi struct {
}

func (ListQABankendApi) GetRequest() interface{} {
	return &ListQABankendRequest{}
}

func (ListQABankendApi) GetResponse() interface{} {
	return &ListQABankendResponse{Data: []*ListQAData{&ListQAData{}}}
}

func (ListQABankendApi) GetApi() string {
	return "ListQABankend"
}

func (ListQABankendApi) GetDesc() string {
	return "列出"
}

func ListQABankend(c *gin.Context) {
	req := &ListQABankendRequest{}
	rsp := &ListQABankendResponse{Data: []*ListQAData{}}
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

	for _, v := range res {
		rsp.Data = append(rsp.Data, &ListQAData{
			Id:      v.Id,
			Desc:    v.Desc,
			Title:   v.Title,
			Content: v.Content,
		})
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
