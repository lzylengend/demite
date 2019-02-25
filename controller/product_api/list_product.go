package product_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListProductRequest struct {
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
	Key    string `json:"key"`
}

type ListProductResponse struct {
	Data   []*productData        `json:"data"`
	Count  int64                 `json:"count"`
	Status *my_error.ErrorCommon `json:"status"`
}

type productData struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	ProductDecs string `json:"productdecs"`
	ProductPic  string `json:"productpic"`
	Price       int64  `json:"price"`
	SortId      int64  `json:"sortid"`
	ClassId     int64  `json:"classid"`
	Show        bool   `json:"show"`
	Num         int64  `json:"num"`
}

type ListProductApi struct {
}

func (ListProductApi) GetRequest() interface{} {
	return &ListProductRequest{}
}

func (ListProductApi) GetResponse() interface{} {
	return &ListProductResponse{
		Data: []*productData{
			&productData{},
		},
	}
}

func (ListProductApi) GetApi() string {
	return "ListProduct"
}

func (ListProductApi) GetDesc() string {
	return "列出产品"
}

func ListProduct(c *gin.Context) {
	req := &ListProductRequest{}
	rsp := &ListProductResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	data, err := model.ProduceDao.ListByCreateTime(req.Key, req.Limit, req.Offset)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	count, err := model.ProduceDao.CountByKey(req.Key)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Data = make([]*productData, 0)
	for _, v := range data {
		show := true
		if v.Show > 0 {
			show = false
		}
		rsp.Data = append(rsp.Data, &productData{
			Id:          v.ProductId,
			Name:        v.ProductName,
			ProductDecs: v.ProductDecs,
			ProductPic:  v.ProductPic,
			Price:       v.Price,
			SortId:      v.SortId,
			ClassId:     v.ClassId,
			Show:        show,
			Num:         v.Num,
		})
	}
	rsp.Count = count
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
