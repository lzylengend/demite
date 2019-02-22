package product_api

import (
	"demite/conf"
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
)

type UpdateProductRequest struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	ProductDecs string `json:"productdecs"`
	ProductPic  string `json:"productpic"`
	Price       int64  `json:"price"`
	SortId      int64  `json:"sortid"`
	ClassId     int64  `json:"classid"`
	Show        bool   `json:"show"`
}

type UpdateProductResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type UpdateProductApi struct {
}

func (UpdateProductApi) GetRequest() interface{} {
	return &UpdateProductRequest{}
}

func (UpdateProductApi) GetResponse() interface{} {
	return &UpdateProductResponse{}
}

func (UpdateProductApi) GetApi() string {
	return "UpdateProduct"
}

func (UpdateProductApi) GetDesc() string {
	return "修改产品"
}

func UpdateProduct(c *gin.Context) {
	req := &UpdateProductRequest{}
	rsp := &UpdateProductResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	if req.Name == "" {
		rsp.Status = my_error.NotNilError("产品名字")
		c.JSON(200, rsp)
		return
	}

	if req.Price <= 0 {
		rsp.Status = my_error.ParamError("价格")
		c.JSON(200, rsp)
		return
	}

	_, b, err := model.ClassDao.ExistId(req.ClassId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}
	if !b {
		rsp.Status = my_error.ParamError("classid")
		c.JSON(200, rsp)
		return
	}

	_, err = ioutil.ReadFile(conf.GetFilePath() + "/" + req.ProductPic)
	if err != nil {
		rsp.Status = my_error.FileReadError(err.Error())
		c.JSON(200, rsp)
		return
	}

	p, err := model.ProduceDao.GetById(req.Id)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	if req.Show {
		p.Show = 0
	} else {
		p.Show = time.Now().Unix()
	}
	p.ProductName = req.Name
	p.ProductDecs = req.ProductDecs
	p.ProductPic = req.ProductPic
	p.Price = req.Price
	p.SortId = req.SortId
	p.ClassId = req.ClassId

	err = model.ProduceDao.Set(p)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
