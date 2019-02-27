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
	GoodsId                 int64      `gorm:"column:goodsid;primary_key;AUTO_INCREMENT"`
	GoodsUUID               string     `gorm:"column:goodsuuid;index:goodsuuid;unique"`
	GoodsCode               string     `grom:"column:goodscode;index:goodscode;unique"`
	GoodsName               string     `gorm:"column:goodsname;index:goodsname"`
	GoodsDecs               string     `gorm:"column:goodsdesc;type:text"`
	GoodsPic                string     `gorm:"column:goodspic;type:text"`
	Price                   int64      `gorm:"column:price;"`
	QRCode                  string     `gorm:"column:qrcode;type:text"`
	GoodsTemplet            string     `gorm:"column:goodstemplet;type:text"`
	GoodsTempletLockContext string     `gorm:"column:goodstempletlockcontext;type:text"`
	GoodsTempletId          int64      `gorm:"column:goodstempletid;index:goodstempletid"`
	OrderId                 int64      `gorm:"column:orderid;index:orderid"`
	ProductId               int64      `gorm:"column:productid;index:productid"`
	ClassId                 int64      `gorm:"column:classid;index:classid"`
	CreatorId               int64      `gorm:"column:creatorid"`
	Status                  goodStatus `gorm:"column:status"`
	DataStatus              int64      `gorm:"column:datastatus"`
	CreateTime              int64      `gorm:"column:createtime"`
	UpdateTime              int64      `gorm:"column:updatetime"`
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

func (this *_GoodsDao) CreateCode() string {
	return time.Now().Format("20060102150405") + uuid.New().String()
}

func (this *_GoodsDao) GetByCode(code string) (*Goods, error) {
	obj := &Goods{}
	err := this.Db.Where("goodscode = ?", code).First(obj).Error
	return obj, err
}

func (this *_GoodsDao) ExitByCode(code string) (bool, error) {
	_, err := this.GetByCode(code)
	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil {
		return true, err
	}
	return true, nil
}

func (this *_GoodsDao) GetByUUID(uuid string) (*Goods, error) {
	obj := &Goods{}
	err := this.Db.Where("goodsuuid = ?", uuid).First(obj).Error
	return obj, err
}

func (this *_GoodsDao) ExitByUUID(uuid string) (bool, error) {
	_, err := this.GetByCode(uuid)
	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil {
		return true, err
	}
	return true, nil
}

func (this *_GoodsDao) Add(obj *Goods) (int64, error) {
	obj.CreateTime = time.Now().Unix()
	obj.UpdateTime = obj.CreateTime
	obj.DataStatus = 0
	obj.Status = GOODINIT
	err := this.Db.Create(obj).Error
	return obj.GoodsId, err
}
