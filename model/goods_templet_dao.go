package model

import "github.com/jinzhu/gorm"

type GoodsTemplet struct {
	GoodsTempletId          int64  `gorm:"column:goodstempletid;primary_key;AUTO_INCREMENT"`
	GoodsTempletName        string `gorm:"column:goodstempletname;index:goodstempletname"`
	GoodsTempletContext     string `gorm:"column:goodstempletcontext;type:text"`
	GoodsTempletLockContext string `gorm:"column:goodstempletlockcontext;type:text"`
	CreatorId               int64  `gorm:"column:creatorid"`
	DataStatus              int64  `gorm:"column:datastatus"`
	CreateTime              int64  `gorm:"column:createtime"`
	UpdateTime              int64  `gorm:"column:updatetime"`
}

type _GoodsTempletDao struct {
	Db *gorm.DB
}

func (GoodsTemplet) TableName() string {
	return "goodstemplet"
}

func newGoodsTempletDao(db *gorm.DB) *_GoodsTempletDao {
	db.AutoMigrate(&GoodsTemplet{})

	return &_GoodsTempletDao{Db: db.Model(&GoodsTemplet{})}
}
