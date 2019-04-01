package goods_api

import (
	"demite/conf"
	"demite/model"
	"demite/my_error"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type GoodsListRequest struct {
	Limit        int64  `json:"limit"`
	Offset       int64  `json:"offset"`
	Key          string `json:"key"`
	CreateQRCode string `json:"createqrcode"`
}

type GoodsListResponse struct {
	Data   []*good               `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type good struct {
	Name                    string `json:"name"`
	GoodsUUID               string `json:"goodsuuid"`
	GoodsDecs               string `json:"goodsdecs"`
	GoodsPic                string `json:"goodspic"`
	GoodsTemplet            string `json:"goodsteplet"`
	GoodsTempletLockContext string `json:"goodstempletlockcontext"`
	CreateTime              int64  `json:"createtime"`
	QRCode                  string `json:"qrcode"`
	GoodsModel              string `json:"goodmodel"`
	GuaranteeTime           int64  `json:"guaranteetime"`
	GoodsPicData            string `json:"goodpicdata"`
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

	objList, err := model.GoodsDao.ListByQRCode(req.Key, req.Limit, req.Offset)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	for _, v := range objList {
		data, err := ioutil.ReadFile(conf.GetFilePath() + "/" + v.GoodsPic)
		if err != nil {
			data = []byte{}
			//rsp.Status = my_error.FileReadError(err.Error())
			//c.JSON(200, rsp)
			continue
		}

		rsp.Data = append(rsp.Data, &good{
			Name:                    v.GoodsName,
			GoodsUUID:               v.GoodsUUID,
			GoodsDecs:               v.GoodsDecs,
			GoodsPic:                v.GoodsPic,
			GoodsTemplet:            v.GoodsTemplet,
			GoodsTempletLockContext: v.GoodsTempletLockContext,
			CreateTime:              v.CreateTime,
			QRCode:                  v.QRCode,
			GoodsModel:              v.GoodsModel,
			GuaranteeTime:           v.GuaranteeTime,
			GoodsPicData:            base64.StdEncoding.EncodeToString(data),
		})
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
