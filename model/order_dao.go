package model

import (
	"github.com/jinzhu/gorm"
	"math/rand"
	"strconv"
	"time"
)

type orderStatus string

const (
	ORDERNOTPAYMENT = "ORDERNOTPAYMENT"
	ORDERCOMPLETE   = "ORDERCOMPLETE"
)

type Order struct {
	OrderId       int64  `gorm:"column:orderid;primary_key;AUTO_INCREMENT"`
	OrderCode     string `gorm:"column:ordercode;index:ordercode;unique"`
	CreateTime    int64  `gorm:"column:createtime;index:createtime;"`
	CreateId      int64  `gorm:"column:createid;"`
	Status        string `gorm:"column:status"`
	CouponId      int64  `gorm:"column:couponid"`
	OriginalPrice int64  `gorm:"column:originalprice"`
	TotalPrice    int64  `gorm:"column:totalprice"`
}

type _OrdertDao struct {
	Db *gorm.DB
}

func (Order) TableName() string {
	return "order"
}

func newOrderDao(db *gorm.DB) *_OrdertDao {
	db.AutoMigrate(&Order{})

	return &_OrdertDao{Db: db.Model(&Order{})}
}

func (this *_OrdertDao) CreateOrder(createId, couponId, originalPrice, totalPrice int64, gList []*Goods) (*Order, error) {
	obj := &Order{
		OrderCode:     this.CreateCode(createId),
		CreateTime:    time.Now().Unix(),
		CreateId:      createId,
		Status:        ORDERNOTPAYMENT,
		CouponId:      couponId,
		OriginalPrice: originalPrice,
		TotalPrice:    totalPrice,
	}

	objLog := &OrderLog{
		OrderCode:     obj.OrderCode,
		CreateTime:    obj.CreateTime,
		CreateId:      obj.CreateId,
		Status:        obj.Status,
		CouponId:      obj.CouponId,
		OriginalPrice: obj.OriginalPrice,
		TotalPrice:    obj.TotalPrice,
	}

	tx := this.Db.Begin()
	err := tx.Create(obj).Error
	if err != nil {
		tx.Rollback()
		return obj, err
	}

	objLog.OrderId = obj.OrderId
	err = tx.Create(objLog).Error
	if err != nil {
		tx.Rollback()
		return obj, err
	}

	for i, _ := range gList {
		gList[i].OrderId = obj.OrderId
		err = tx.Create(gList[i]).Error
		if err != nil {
			tx.Rollback()
			return obj, err
		}
	}

	tx.Commit()
	return obj, nil
}

func (this *_OrdertDao) CreateCode(createId int64) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	return time.Now().Format("20060102150405") + strconv.FormatInt(createId, 10) + strconv.FormatInt(r.Int63n(90)+10, 10)
}
