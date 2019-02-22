package product_api

import (
	"demite/conf"
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type AddProductRequest struct {
	Name        string `json:"name"`
	ProductDecs string `json:"productdecs"`
	ProductPic  string `json:"productpic"`
	Price       int64  `json:"price"`
	SortId      int64  `json:"sortid"`
	ClassId     int64  `json:"classid"`
}

type AddProductResponse struct {
	Id     int64                 `json:"id"`
	Status *my_error.ErrorCommon `json:"status"`
}

type AddProductApi struct {
}

func (AddProductApi) GetRequest() interface{} {
	return &AddProductRequest{}
}

func (AddProductApi) GetResponse() interface{} {
	return &AddProductResponse{}
}

func (AddProductApi) GetApi() string {
	return "AddProduct"
}

func (AddProductApi) GetDesc() string {
	return "新增产品"
}

func AddProduct(c *gin.Context) {
	req := &AddProductRequest{}
	rsp := &AddProductResponse{}
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

	uid, err := controller.GetUserId(c)
	if err != nil {
		rsp.Status = my_error.NoLoginError()
		c.JSON(200, rsp)
		return
	}

	p, err := model.ProduceDao.Insert(req.Name, req.ProductDecs, req.ProductPic, req.Price, req.SortId, req.ClassId, uid)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Id = p.ProductId
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
