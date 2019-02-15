package model

import "github.com/jinzhu/gorm"

type Staff struct {
	StaffId    int64  `gorm:"column:staffid;primary_key;AUTO_INCREMENT"`
	StaffName  string `gorm:"column:staffname;index:staffname"`
	StaffDecs  string `gorm:"column:staffdesc;type:text"`
	StaffPhone string `gorm:"column:staffphone;"`
	CreatorId  int64  `gorm:"column:creatorid"`
	DataStatus int64  `gorm:"column:datastatus"`
	CreateTime int64  `gorm:"column:createtime"`
	UpdateTime int64  `gorm:"column:updatetime"`
}

type _StaffDao struct {
	Db *gorm.DB
}

func (Staff) TableName() string {
	return "staff"
}

func newStaffDao(db *gorm.DB) *_StaffDao {
	db.AutoMigrate(&Staff{})

	return &_StaffDao{Db: db.Model(&Staff{})}
}
