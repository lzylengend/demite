package scheme_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type GetSchemeBankendRequest struct {
	Id int64 `json:"id"`
}

type GetSchemeBankendResponse struct {
	Data   *GetSchemeData        `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type GetSchemeData struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
	FileId  string `json:"fileid"`
}

type GetSchemeBankendApi struct {
}

func (GetSchemeBankendApi) GetRequest() interface{} {
	return &GetSchemeBankendRequest{}
}

func (GetSchemeBankendApi) GetResponse() interface{} {
	return &GetSchemeBankendResponse{Data: &GetSchemeData{}}
}

func (GetSchemeBankendApi) GetApi() string {
	return "GetSchemeBankend"
}

func (GetSchemeBankendApi) GetDesc() string {
	return "获取"
}

func GetSchemeBankend(c *gin.Context) {
	req := &GetSchemeBankendRequest{}
	rsp := &GetSchemeBankendResponse{}
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
		Id:      res.Id,
		Desc:    res.Desc,
		Title:   res.Title,
		Content: res.Content,
		FileId:  res.FileId,
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
