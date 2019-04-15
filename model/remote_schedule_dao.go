package model

import "github.com/jinzhu/gorm"

type RemoteSchedule struct {
	RemoteScheduleId int64        `gorm:"column:RemoteScheduleid;primary_key;AUTO_INCREMENT"`
	RemoteId         int64        `gorm:"column:repairid;index:repairid"`
	CreateId         int64        `gorm:"column:createid;"`
	WxUserId         int64        `gorm:"column:wxuserid;"`
	StaffId          int64        `gorm:"column:staffid"`
	RemoteTime       int64        `gorm:"column:repairtime"`
	CreateTime       int64        `gorm:"column:createtime"`
	UpdateTime       int64        `gorm:"column:updatetime"`
	Status           repairStatus `gorm:"column:status"`
	DataStatus       int64        `gorm:"column:datastatus"`
}
type _RemoteScheduleDao struct {
	Db *gorm.DB
}

func (RemoteSchedule) TableName() string {
	return "RemoteSchedule"
}

func newRemoteScheduleDao(db *gorm.DB) *_RemoteScheduleDao {
	db.AutoMigrate(&RemoteSchedule{})

	return &_RemoteScheduleDao{Db: db.Model(&RemoteSchedule{})}
}

func (this *_RemoteScheduleDao) ListByRemoteId(repairId int64) ([]*RemoteSchedule, error) {
	objList := make([]*RemoteSchedule, 0)

	err := this.Db.Where("repairid = ? and datastatus = ? ", repairId, 0).Order("createtime").Find(&objList).Error

	return objList, err
}
