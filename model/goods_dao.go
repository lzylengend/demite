package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
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
	GoodsModel              string     `gorm:"column:goodmodel;index:goodmodel"`
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
	Province                string     `gorm:"column:province;index:province"`
	ProvinceId              int64      `gorm:"column:provinceid;index:provinceid"`
	Hospital                string     `gorm:"column:hospital;index:hospital"`
	GuaranteeTime           int64      `gorm:"column:guaranteetime"`
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
	err := this.Db.Where("goodsuuid = ?", code).First(obj).Error
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

func (this *_GoodsDao) Get(id int64) (*Goods, error) {
	obj := &Goods{}
	err := this.Db.Where("goodsid = ?", id).First(obj).Error
	return obj, err
}

func (this *_GoodsDao) Set(obj *Goods) error {
	err := this.Db.Save(obj).Error
	return err
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

func (this *_GoodsDao) ListByQRCode(key string, limit, offset int64, hospital string, province string) ([]*Goods, error) {
	objList := make([]*Goods, 0)
	sql := `datastatus  = ?`
	args := make([]interface{}, 0)
	args = append(args, 0)

	if key != "" {
		sql = sql + ` and goodsname like ? `
		args = append(args, "%"+key+"%")
	}

	if hospital != "" {
		sql = sql + ` and hospital like ? `
		args = append(args, "%"+hospital+"%")
	}

	if province != "" {
		sql = sql + ` and province = ? `
		args = append(args, province)
	}

	err := this.Db.Where(sql, args...).Offset(offset).Limit(limit).Order("createtime desc").Find(&objList).Error

	return objList, err
}

func (this *_GoodsDao) CountByQRCode(key string, hospital string, province string) (int64, error) {
	var n int64
	sql := `datastatus  = ?`
	args := make([]interface{}, 0)
	args = append(args, 0)

	if key != "" {
		sql = sql + ` and goodsname like ? `
		args = append(args, "%"+key+"%")
	}

	if hospital != "" {
		sql = sql + ` and hospital like ? `
		args = append(args, "%"+hospital+"%")
	}

	if province != "" {
		sql = sql + ` and province = ? `
		args = append(args, province)
	}

	err := this.Db.Where(sql, args...).Count(&n).Error

	return n, err
}
