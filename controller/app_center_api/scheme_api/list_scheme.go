package scheme_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListSchemeRequest struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type ListSchemeResponse struct {
	Data   []*ListSchemeData     `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type ListSchemeData struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type ListSchemeApi struct {
}

func (ListSchemeApi) GetRequest() interface{} {
	return &ListSchemeRequest{}
}

func (ListSchemeApi) GetResponse() interface{} {
	return &ListSchemeResponse{Data: []*ListSchemeData{&ListSchemeData{}}}
}

func (ListSchemeApi) GetApi() string {
	return "ListScheme"
}

func (ListSchemeApi) GetDesc() string {
	return "列出"
}

func ListScheme(c *gin.Context) {
	req := &ListSchemeRequest{}
	rsp := &ListSchemeResponse{Data: []*ListSchemeData{}}
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

	for _, v := range res {
		rsp.Data = append(rsp.Data, &ListSchemeData{
			Id:    v.Id,
			Desc:  v.Desc,
			Title: v.Title,
		})
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
