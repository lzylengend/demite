package model

import (
	"github.com/jinzhu/gorm"
)

type DelayGuaranteeAppl struct {
	Id         int64              `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	GoodsUUID  string             `gorm:"column:goodsuuid;index:goodsuuid"`
	WXUserId   int64              `gorm:"column:wxuserid;index:wxuserid"`
	Status     goodsWXUserSatatus `gorm:"column:status"`
	Creater    int64              `gorm:"column:creater"`
	DataStatus int64              `gorm:"column:datastatus"`
	CreateTime int64              `gorm:"column:createtime"`
	UpdateTime int64              `gorm:"column:updatetime"`
}
type _DelayGuaranteeApplDao struct {
	Db *gorm.DB
}

func (DelayGuaranteeAppl) TableName() string {
	return "delpayguranteeapply"
}

func newDelayGuaranteeApplDao(db *gorm.DB) *_DelayGuaranteeApplDao {
	db.AutoMigrate(&DelayGuaranteeAppl{})

	return &_DelayGuaranteeApplDao{Db: db.Model(&DelayGuaranteeAppl{})}
}

func (this *_DelayGuaranteeApplDao) Add(goodUUID string, wxUserId int64) (*DelayGuaranteeAppl, error) {

}
