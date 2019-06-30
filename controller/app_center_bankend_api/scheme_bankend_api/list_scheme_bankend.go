package scheme_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListSchemeBankendRequest struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type ListSchemeBankendResponse struct {
	Data   []*ListSchemeData     `json:"data"`
	Count  int64                 `json:"count"`
	Status *my_error.ErrorCommon `json:"status"`
}

type ListSchemeData struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	FileId string `json:"fileid"`
}

type ListSchemeBankendApi struct {
}

func (ListSchemeBankendApi) GetRequest() interface{} {
	return &ListSchemeBankendRequest{}
}

func (ListSchemeBankendApi) GetResponse() interface{} {
	return &ListSchemeBankendResponse{Data: []*ListSchemeData{&ListSchemeData{}}}
}

func (ListSchemeBankendApi) GetApi() string {
	return "ListSchemeBankend"
}

func (ListSchemeBankendApi) GetDesc() string {
	return "列出"
}

func ListSchemeBankend(c *gin.Context) {
	req := &ListSchemeBankendRequest{}
	rsp := &ListSchemeBankendResponse{Data: []*ListSchemeData{}}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	res, err := model.SchemeDao.List(req.Limit, req.Offset)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	count, err := model.SchemeDao.Count()
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	for _, v := range res {
		rsp.Data = append(rsp.Data, &ListSchemeData{
			Id:     v.Id,
			Desc:   v.Desc,
			Title:  v.Title,
			FileId: v.FileId,
		})
	}

	rsp.Count = count
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
