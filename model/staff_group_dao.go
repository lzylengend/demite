package model

import "github.com/jinzhu/gorm"

type StaffGroup struct {
	StaffGroupId   int64  `gorm:"column:staffgroupid;primary_key;AUTO_INCREMENT"`
	StaffGroupName string `gorm:"column:staffgroupname;index:staffgroupname"`
	StaffGroupDecs string `gorm:"column:staffgroupdesc;type:text"`
	CreatorId      int64  `gorm:"column:creatorid"`
	DataStatus     int64  `gorm:"column:datastatus"`
	CreateTime     int64  `gorm:"column:createtime"`
	UpdateTime     int64  `gorm:"column:updatetime"`
}

type _StaffGroupDao struct {
	Db *gorm.DB
}

func (StaffGroup) TableName() string {
	return "staffgroup"
}

func newStaffGroupDao(db *gorm.DB) *_StaffGroupDao {
	db.AutoMigrate(&StaffGroup{})

	return &_StaffGroupDao{Db: db.Model(&StaffGroup{})}
}
