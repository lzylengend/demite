package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type goodStatus string

const (
	GOODINIT     goodStatus = "GOODINIT"
	GOODCOMPLETE goodStatus = "GOODCOMPLETE"
)

type Goods struct {
	GoodsId        int64      `gorm:"column:goodsid;primary_key;AUTO_INCREMENT"`
	GoodsCode      string     `grom:"column:goodscode;index:goodscode;unique"`
	GoodsName      string     `gorm:"column:goodsname;index:goodsname"`
	GoodsDecs      string     `gorm:"column:goodsdesc;type:text"`
	GoodsPic       string     `gorm:"column:goodspic;type:text"`
	Price          int64      `gorm:"column:price;"`
	GoodsTemplet   string     `gorm:"column:goodstemplet"`
	GoodsTempletId int64      `gorm:"column:goodstempletid;index:goodstempletid"`
	OrderId        int64      `gorm:"column:orderid;index:orderid"`
	ProductId      int64      `gorm:"column:productid;index:productid"`
	ClassId        int64      `gorm:"column:classid;index:classid"`
	CreatorId      int64      `gorm:"column:creatorid"`
	Status         goodStatus `gorm:"column:status"`
	DataStatus     int64      `gorm:"column:datastatus"`
	CreateTime     int64      `gorm:"column:createtime"`
	UpdateTime     int64      `gorm:"column:updatetime"`
}
type _GoodsDao struct {
	Db *gorm.DB
}

func (Goods) TableName() string {
	return "goods"
}

func newGoodsDao(db *gorm.DB) *_GoodsDao {
	db.AutoMigrate(&Goods{})

	return &_GoodsDao{Db: db.Model(&Goods{})}
}

func (this *_GoodsDao) CreateCode(createId int64) string {
	return time.Now().Format("20060102150405") + uuid.New().String()
}
