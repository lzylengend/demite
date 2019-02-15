package model

import "github.com/jinzhu/gorm"

type Product struct {
	ProductId   int64  `gorm:"column:productid;primary_key;AUTO_INCREMENT"`
	ProductName string `gorm:"column:productname;index:productname"`
	ProductDecs string `gorm:"column:productdesc;type:text"`
	ProductPic  string `gorm:"column:productpic;type:text"`
	Price       int64  `gorm:"column:price;"`
	SortId      int64  `gorm:"column:sortid;"`
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
