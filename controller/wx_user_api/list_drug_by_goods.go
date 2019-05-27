package wx_user_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"

	"github.com/gin-gonic/gin"
)

type ListDrugByGoodsRequest struct {
	GoodUUID string `json:"gooduuid"`
}

type ListDrugByGoodsResponse struct {
	Status   *my_error.ErrorCommon `json:"status"`
	GoodName string                `json:"goodname"`
	Data     []*drugData           `json:"data"`
}

type ListDrugByGoodsApi struct {
}

func (ListDrugByGoodsApi) GetRequest() interface{} {
	return &ListDrugByGoodsRequest{}
}

func (ListDrugByGoodsApi) GetResponse() interface{} {
	return &ListDrugByGoodsResponse{}
}

func (ListDrugByGoodsApi) GetApi() string {
	return "ListDrugByGoods"
}

func (ListDrugByGoodsApi) GetDesc() string {
	return "列出设备绑定的药品"
}

type drugData struct {
	Id                    int64  `json:"id"`
	ClassId               int64  `json:"classid"`
	ClassName             string `json:"classname"`
	Name                  string `json:"name"`
	Reagent               string `json:"reagent"`               //试剂
	ChromatographicColumn string `json:"chromatographiccolumn"` //色谱柱
	Controls              string `json:"controls"`              //质控品
	TestMethod            string `json:"testmethod"`            //检测方法
	Preprocessing         string `json:"preprocessing"`         //样品预处理
	PotencyRange          string `json:"range"`                 //浓度范围
}

func ListDrugByGoods(c *gin.Context) {
	req := &ListDrugByGoodsRequest{}
	rsp := &ListDrugByGoodsResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	_, err = controller.GetWxUserId(c)
	if err != nil {
		rsp.Status = my_error.NoLoginError()
		c.JSON(200, rsp)
		return
	}

	gdList, err := model.GoodDrugsDao.GetByUUID(req.GoodUUID)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	g, err := model.GoodsDao.GetByUUID(req.GoodUUID)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.GoodName = g.GoodsName

	for _, v := range gdList {
		drug, err := model.DrugDao.Get(v.DrugId)
		if err != nil {
			rsp.Status = my_error.DbError(err.Error())
			c.JSON(200, rsp)
			return
		}

		class, err := model.DrugClassDao.Get(drug.DrugClassId)
		if err != nil {
			rsp.Status = my_error.DbError(err.Error())
			c.JSON(200, rsp)
			return
		}

		rsp.Data = append(rsp.Data, &drugData{
			Id:                    drug.DrugId,
			ClassId:               drug.DrugClassId,
			Name:                  drug.DrugName,
			Reagent:               drug.Reagent,
			ChromatographicColumn: drug.ChromatographicColumn,
			Controls:              drug.Controls,
			TestMethod:            drug.TestMethod,
			ClassName:             class.ClassName,
			Preprocessing:         drug.Preprocessing,
			PotencyRange:          drug.PotencyRange,
		})
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
