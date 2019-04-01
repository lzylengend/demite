package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type repairStatus string

const (
	REPAIRSTATUSINIT = "init"
)

type Repair struct {
	RepairId   int64        `gorm:"column:repairid;primary_key;AUTO_INCREMENT"`
	GoodUUID   string       `gorm:"column:gooduuid;index:gooduuid"`
	WXUserId   int64        `gorm:"column:wxuserid;index:wxuserid"`
	GoodModel  string       `gorm:"column:goodmodel"`
	Phone      string       `gorm:"column:phone"`
	Name       string       `gorm:"column:name"`
	FaultDesc  string       `gorm:"column:faultdesc"`
	FaultType  string       `gorm:"column:faulttype"`
	CreateTime int64        `gorm:"column:createtime"`
	FileId     string       `gorm:"column:fileid"`
	Status     repairStatus `gorm:"column:status"`
	DataStatus int64        `gorm:"column:datastatus"`
}
type _RepairDao struct {
	Db *gorm.DB
}

func (Repair) TableName() string {
	return "Repair"
}

func newRepairDao(db *gorm.DB) *_RepairDao {
	db.AutoMigrate(&Repair{})

	return &_RepairDao{Db: db.Model(&Repair{})}
}

func (this *_RepairDao) Add(gooduuid string, goodmodel string, phone string, name string, faultdesc string, faulttype string, fileid string, wxuserid int64) (*Repair, error) {
	obj := &Repair{
		GoodUUID:   gooduuid,
		WXUserId:   wxuserid,
		GoodModel:  goodmodel,
		Phone:      phone,
		Name:       name,
		FaultDesc:  faultdesc,
		FaultType:  faulttype,
		CreateTime: time.Now().Unix(),
		FileId:     fileid,
		Status:     REPAIRSTATUSINIT,
		DataStatus: 0,
	}

	err := this.Db.Create(obj).Error
	return obj, err
}
