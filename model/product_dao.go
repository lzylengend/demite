package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Product struct {
	ProductId   int64  `gorm:"column:productid;primary_key;AUTO_INCREMENT"`
	ProductName string `gorm:"column:productname;index:productname"`
	ProductDecs string `gorm:"column:productdesc;type:text"`
	ProductPic  string `gorm:"column:productpic;type:text"`
	Price       int64  `gorm:"column:price;"`
	SortId      int64  `gorm:"column:sortid;"`
	Show        int64  `gorm:"column:show;"`
	Num         int64  `gorm:"column:num;"`
	ClassId     int64  `gorm:"column:classid;index:classid"`
	CreatorId   int64  `gorm:"column:creatorid"`
	DataStatus  int64  `gorm:"column:datastatus"`
	CreateTime  int64  `gorm:"column:createtime"`
	UpdateTime  int64  `gorm:"column:updatetime"`
}

type _ProductDao struct {
	Db *gorm.DB
}

func (Product) TableName() string {
	return "product"
}

func newProductDao(db *gorm.DB) *_ProductDao {
	db.AutoMigrate(&Product{})

	return &_ProductDao{Db: db.Model(&Product{})}
}

func (this *_ProductDao) Insert(productName, productDecs, productPic string, price, sortId, classId, creatorId, num int64) (*Product, error) {
	obj := &Product{
		ProductName: productName,
		ProductDecs: productDecs,
		ProductPic:  productPic,
		Price:       price,
		SortId:      sortId,
		ClassId:     classId,
		CreatorId:   creatorId,
		Show:        time.Now().Unix(),
		Num:         num,
		DataStatus:  0,
		CreateTime:  time.Now().Unix(),
		UpdateTime:  time.Now().Unix(),
	}

	err := this.Db.Create(obj).Error
	return obj, err
}

func (this *_ProductDao) ListByCreateTime(key string, limit, offset int64) ([]*Product, error) {
	objList := make([]*Product, 0)

	key = "%" + key + "%"
	err := this.Db.Where("productname like ? and datastatus = ?", key, 0).Offset(offset).Limit(limit).Order("createtime").Find(&objList).Error
	return objList, err
}

func (this *_ProductDao) CountByKey(key string) (int64, error) {
	n := 0

	key = "%" + key + "%"
	err := this.Db.Where("productname like ? and datastatus = ?", key, 0).Count(&n).Error
	return int64(n), err
}

func (this *_ProductDao) GetById(id int64) (*Product, error) {
	obj := &Product{}
	err := this.Db.Where("productid = ?", id).First(obj).Error
	return obj, err
}

func (this *_ProductDao) Set(obj *Product) error {
	obj.UpdateTime = time.Now().Unix()
	return this.Db.Save(obj).Error
}
