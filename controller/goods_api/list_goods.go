package goods_api

import (
	"demite/my_error"
	"github.com/gin-gonic/gin"
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
}
