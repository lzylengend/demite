package goods_api

import (
	"demite/conf"
	"demite/model"
	"demite/my_error"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type GetGoodRequest struct {
	GoodsUUID string `json:"goodsuuid"`
}

type GetGoodResponse struct {
	data                    []*drugData           `json:"data"`
	Name                    string                `json:"name"`
	GoodsUUID               string                `json:"goodsuuid"`
	GoodsDecs               string                `json:"goodsdecs"`
	GoodsPic                string                `json:"goodspic"`
	GoodsTemplet            string                `json:"goodsteplet"`
	GoodsTempletLockContext string                `json:"goodstempletlockcontext"`
	CreateTime              int64                 `json:"createtime"`
	QRCode                  string                `json:"qrcode"`
	GoodsModel              string                `json:"goodmodel"`
	GuaranteeTime           int64                 `json:"guaranteetime"`
	GoodsPicData            string                `json:"goodpicdata"`
	Status                  *my_error.ErrorCommon `json:"status"`
}

type GetGoodApi struct {
}

func (GetGoodApi) GetRequest() interface{} {
	return &GetGoodRequest{}
}

func (GetGoodApi) GetResponse() interface{} {
	return &GetGoodResponse{}
}

func (GetGoodApi) GetApi() string {
	return "GetGood"
}

func (GetGoodApi) GetDesc() string {
	return "id获取货物"
}

func GetGood(c *gin.Context) {
	req := &GetGoodRequest{}
	rsp := &GetGoodResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	objList, err := model.GoodDrugsDao.GetByUUID(req.GoodsUUID)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	good, err := model.GoodsDao.GetByUUID(req.GoodsUUID)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	data, err := ioutil.ReadFile(conf.GetFilePath() + "/" + good.GoodsPic)
	if err != nil {
		data = []byte{}
		//rsp.Status = my_error.FileReadError(err.Error())
		//c.JSON(200, rsp)
	}

	rsp.Name = good.GoodsName
	rsp.GoodsDecs = good.GoodsDecs
	rsp.GoodsPic = good.GoodsPic
	rsp.GoodsTemplet = good.GoodsTemplet
	rsp.CreateTime = good.CreateTime
	rsp.QRCode = good.QRCode
	rsp.GoodsModel = good.GoodsModel
	rsp.GoodsPicData = base64.StdEncoding.EncodeToString(data)
	rsp.GuaranteeTime = good.GuaranteeTime
	rsp.GoodsUUID = good.GoodsUUID
	rsp.GoodsTempletLockContext = good.GoodsTempletLockContext

	for _, v := range objList {
		drug, err := model.DrugDao.Get(v.DrugId)
		if err != nil {
			rsp.Status = my_error.DbError(err.Error())
			c.JSON(200, rsp)
			return
		}

		class, err := model.DrugClassDao.Get(drug.DrugId)
		if err != nil {
			rsp.Status = my_error.DbError(err.Error())
			c.JSON(200, rsp)
			return
		}

		rsp.data = append(rsp.data, &drugData{
			Id:                    v.DrugId,
			ClassId:               drug.DrugClassId,
			ClassName:             class.ClassName,
			Name:                  drug.DrugName,
			Reagent:               drug.Reagent,
			ChromatographicColumn: drug.ChromatographicColumn,
			Controls:              drug.Controls,
			TestMethod:            drug.TestMethod,
		})

	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
