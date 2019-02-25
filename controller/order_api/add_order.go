package order_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type AddOrderRequest struct {
	Goods []*orderGoods `json:"goods"`
}

type orderGoods struct {
	ProductId int64 `json:"productid"`
	Num       int64 `json:"num"`
}

type AddOrderResponse struct {
	OrderCode string                `json:"ordercode"`
	Status    *my_error.ErrorCommon `json:"status"`
}

type AddOrderApi struct {
}

func (AddOrderApi) GetRequest() interface{} {
	return &AddOrderRequest{
		Goods: []*orderGoods{
			&orderGoods{
				ProductId: 0,
				Num:       0,
			},
		},
	}
}

func (AddOrderApi) GetResponse() interface{} {
	return &AddOrderResponse{}
}

func (AddOrderApi) GetApi() string {
	return "AddOrder"
}

func (AddOrderApi) GetDesc() string {
	return "新增订单"
}

func AddOrder(c *gin.Context) {
	req := &AddOrderRequest{}
	rsp := &AddOrderResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	//FIXME
	var uId int64 = 0
	//uId, err := controller.GetWxUserId(c)
	//if err != nil {
	//	rsp.Status = my_error.NoLoginError()
	//	c.JSON(200, rsp)
	//	return
	//}

	if len(req.Goods) == 0 {
		rsp.Status = my_error.ParamError("goods")
		c.JSON(200, rsp)
		return
	}

	gList := make([]*model.Goods, 0)
	var originalPrice int64 = 0
	var totalPrice int64 = 0
	for _, v := range req.Goods {
		if v.Num == 0 {
			rsp.Status = my_error.ParamError("num")
			c.JSON(200, rsp)
			return
		}

		p, err := model.ProduceDao.GetById(v.ProductId)
		if err != nil {
			rsp.Status = my_error.ParamError("id")
			c.JSON(200, rsp)
			return
		}

		if v.Num < p.Num {
			rsp.Status = my_error.GoodNotEnoughError(p.ProductName)
			c.JSON(200, rsp)
			return
		}

		for i := 0; i < int(v.Num); i++ {
			gList = append(gList, &model.Goods{
				GoodsCode:      model.GoodsDao.CreateCode(uId),
				GoodsName:      p.ProductName,
				GoodsDecs:      p.ProductDecs,
				GoodsPic:       p.ProductPic,
				Price:          p.Price,
				GoodsTemplet:   "",
				GoodsTempletId: 0,
				ProductId:      p.ProductId,
				ClassId:        p.ClassId,
				CreatorId:      uId,
				Status:         model.GOODINIT,
				DataStatus:     0,
				CreateTime:     time.Now().Unix(),
				UpdateTime:     time.Now().Unix(),
			})
		}

		totalPrice = totalPrice + (p.Price * v.Num)
	}

	//TODO 优惠价
	originalPrice = totalPrice

	o, err := model.OrdertDao.CreateOrder(uId, 0, originalPrice, totalPrice, gList)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	rsp.OrderCode = o.OrderCode
	c.JSON(200, rsp)
	return
}
