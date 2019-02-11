package model

import "github.com/jinzhu/gorm"

type Device struct {
	deviceId   int64  `gorm:"column:deviceid;primary_key;AUTO_INCREMENT"`
	deviceName string `gorm:"column:devicename;index:devicename"`
	CreatorId  int64  `gorm:"column:creatorid"`
	DataStatus int64  `gorm:"column:datastatus"`
	CreateTime int64  `gorm:"column:createtime"`
	UpdateTime int64  `gorm:"column:updatetime"`
}
type _DeviceDao struct {
	Db *gorm.DB
}

func (Device) TableName() string {
	return "device"
}

func newDeviceDao(db *gorm.DB) *_DeviceDao {
	db.AutoMigrate(&Device{})

	return &_DeviceDao{Db: db.Model(&Device{})}
}
