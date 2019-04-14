package model

import "github.com/jinzhu/gorm"

type RepairSchedule struct {
	RepairScheduleId int64        `gorm:"column:RepairScheduleid;primary_key;AUTO_INCREMENT"`
	RepairId         int64        `gorm:"column:repairid;index:repairid"`
	CreateId         int64        `gorm:"column:createid;"`
	WxUserId         int64        `gorm:"column:wxuserid;"`
	StaffId          int64        `gorm:"column:staffid"`
	RepairTime       int64        `gorm:"column:repairtime"`
	CreateTime       int64        `gorm:"column:createtime"`
	UpdateTime       int64        `gorm:"column:updatetime"`
	Status           repairStatus `gorm:"column:status"`
	DataStatus       int64        `gorm:"column:datastatus"`
}
type _RepairScheduleDao struct {
	Db *gorm.DB
}

func (RepairSchedule) TableName() string {
	return "RepairSchedule"
}

func newRepairScheduleDao(db *gorm.DB) *_RepairScheduleDao {
	db.AutoMigrate(&RepairSchedule{})

	return &_RepairScheduleDao{Db: db.Model(&RepairSchedule{})}
}

func (this *_RepairScheduleDao) ListByRepairId(repairId int64) ([]*RepairSchedule, error) {
	objList := make([]*RepairSchedule, 0)

	err := this.Db.Where("repairid = ? and datastatus = ? ", repairId, 0).Order("createtime").Find(&objList).Error

	return objList, err
}
