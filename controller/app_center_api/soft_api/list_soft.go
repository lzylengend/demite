package soft_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListSoftRequest struct {
	ClassId int64 `json:"classid"`
	Limit   int64 `json:"limit"`
	Offset  int64 `json:"offset"`
}

type ListSoftResponse struct {
	Data   []*ListSoftData       `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type ListSoftData struct {
	Id      int64  `json:"id"`
	Content string `json:"content"`
	ClassId int64  `json:"classid"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
}

type ListSoftApi struct {
}

func (ListSoftApi) GetRequest() interface{} {
	return &ListSoftRequest{}
}

func (ListSoftApi) GetResponse() interface{} {
	return &ListSoftResponse{Data: []*ListSoftData{&ListSoftData{}}}
}

func (ListSoftApi) GetApi() string {
	return "ListSoft"
}

func (ListSoftApi) GetDesc() string {
	return "列出软件"
}

func ListSoft(c *gin.Context) {
	req := &ListSoftRequest{}
	rsp := &ListSoftResponse{Data: []*ListSoftData{}}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	res, err := model.SoftDao.List(req.ClassId, req.Limit, req.Offset)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	for _, v := range res {
		rsp.Data = append(rsp.Data, &ListSoftData{
			Id:      v.Id,
			Content: v.Content,
			ClassId: v.ClassId,
			Title:   v.Title,
			Desc:    v.Desc,
		})
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
