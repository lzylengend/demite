package model

import "github.com/jinzhu/gorm"

type OrderLog struct {
	OrderId       int64  `gorm:"column:orderid;index:orderid;"`
	OrderCode     string `gorm:"column:ordercode;index:ordercode;"`
	CreateTime    int64  `gorm:"column:createtime;index:createtime;"`
	CreateId      int64  `gorm:"column:createid;"`
	Status        string `gorm:"column:status"`
	CouponId      int64  `gorm:"column:couponid"`
	OriginalPrice int64  `gorm:"column:originalprice"`
	TotalPrice    int64  `gorm:"column:totalprice"`
}

type _OrdertLogDao struct {
	Db *gorm.DB
}

func (OrderLog) TableName() string {
	return "orderlog"
}

func newOrderLogDao(db *gorm.DB) *_OrdertLogDao {
	db.AutoMigrate(&OrderLog{})

	return &_OrdertLogDao{Db: db.Model(&OrderLog{})}
}
