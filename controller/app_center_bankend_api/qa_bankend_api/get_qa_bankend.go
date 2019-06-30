package qa_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type GetQABankendRequest struct {
	Id int64 `json:"id"`
}

type GetQABankendResponse struct {
	Data   *GetQAData            `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type GetQAData struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type GetQABankendApi struct {
}

func (GetQABankendApi) GetRequest() interface{} {
	return &GetQABankendRequest{}
}

func (GetQABankendApi) GetResponse() interface{} {
	return &GetQABankendResponse{Data: &GetQAData{}}
}

func (GetQABankendApi) GetApi() string {
	return "GetQABankend"
}

func (GetQABankendApi) GetDesc() string {
	return "获取问答"
}

func GetQABankend(c *gin.Context) {
	req := &GetQABankendRequest{}
	rsp := &GetQABankendResponse{}
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
		Id:      res.Id,
		Desc:    res.Desc,
		Title:   res.Title,
		Content: res.Content,
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
