package goods_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type GoodsListRequest struct {
	Limit        int64  `json:"limit"`
	Offset       int64  `json:"offset"`
	Key          string `json:"key"`
	Hospital     string `json:"hospital"`
	Province     string `json:"province"`
	CreateQRCode string `json:"createqrcode"`
}

type GoodsListResponse struct {
	Data   []*good               `json:"data"`
	Count  int64                 `json:"count"`
	Status *my_error.ErrorCommon `json:"status"`
}

type good struct {
	Name                    string `json:"name"`
	GoodsUUID               string `json:"goodsuuid"`
	GoodsDecs               string `json:"goodsdecs"`
	GoodsTemplet            string `json:"goodsteplet"`
	GoodsTempletLockContext string `json:"goodstempletlockcontext"`
	CreateTime              int64  `json:"createtime"`
	QRCode                  string `json:"qrcode"`
	GoodsModel              string `json:"goodmodel"`
	GuaranteeTime           int64  `json:"guaranteetime"`
	Hospital                string `json:"hospital"`
	Province                string `json:"province"`
}

type GoodsListApi struct {
}

func (GoodsListApi) GetRequest() interface{} {
	return &GoodsListRequest{}
}

func (GoodsListApi) GetResponse() interface{} {
	return &GoodsListResponse{}
}

func (GoodsListApi) GetApi() string {
	return "GoodsList"
}

func (GoodsListApi) GetDesc() string {
	return "列出货物"
}

func GoodsList(c *gin.Context) {
	req := &GoodsListRequest{}
	rsp := &GoodsListResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	objList, err := model.GoodsDao.ListByQRCode(req.Key, req.Limit, req.Offset, req.Hospital, req.Province)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	count, err := model.GoodsDao.CountByQRCode(req.Key, req.Hospital, req.Province)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	for _, v := range objList {
		rsp.Data = append(rsp.Data, &good{
			Name:                    v.GoodsName,
			GoodsUUID:               v.GoodsUUID,
			GoodsDecs:               v.GoodsDecs,
			GoodsTemplet:            v.GoodsTemplet,
			GoodsTempletLockContext: v.GoodsTempletLockContext,
			CreateTime:              v.CreateTime,
			QRCode:                  v.QRCode,
			GoodsModel:              v.GoodsModel,
			GuaranteeTime:           v.GuaranteeTime,
			Province:                v.Province,
			Hospital:                v.Hospital,
		})
	}

	rsp.Count = count
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
