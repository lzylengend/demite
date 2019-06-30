package soft_bankend_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListSoftBankendRequest struct {
	ClassId int64 `json:"classid"`
	Limit   int64 `json:"limit"`
	Offset  int64 `json:"offset"`
}

type ListSoftBankendResponse struct {
	Data   []*ListSoftData       `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type ListSoftData struct {
	Id        int64  `json:"id"`
	Content   string `json:"content"`
	ClassId   int64  `json:"classid"`
	ClassName string `json:"classname"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
}

type ListSoftBankendApi struct {
}

func (ListSoftBankendApi) GetRequest() interface{} {
	return &ListSoftBankendRequest{}
}

func (ListSoftBankendApi) GetResponse() interface{} {
	return &ListSoftBankendResponse{Data: []*ListSoftData{&ListSoftData{}}}
}

func (ListSoftBankendApi) GetApi() string {
	return "ListSoftBankend"
}

func (ListSoftBankendApi) GetDesc() string {
	return "列出软件下砸"
}

func ListSoftBankend(c *gin.Context) {
	req := &ListSoftBankendRequest{}
	rsp := &ListSoftBankendResponse{Data: []*ListSoftData{}}
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
		name := ""

		cla, err := model.SoftClassDao.Get(v.ClassId)
		if err == nil {
			name = cla.Name
		}

		rsp.Data = append(rsp.Data, &ListSoftData{
			Id:        v.Id,
			Content:   v.Content,
			ClassId:   v.ClassId,
			ClassName: name,
			Desc:      v.Desc,
			Title:     v.Title,
		})
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
