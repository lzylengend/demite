package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Staff struct {
	StaffId      int64  `gorm:"column:staffid;primary_key;AUTO_INCREMENT"`
	StaffNO      string `gorm:"column:staffno;index:staffno;unique"`
	StaffName    string `gorm:"column:staffname;index:staffname"`
	StaffDecs    string `gorm:"column:staffdesc;type:text"`
	StaffPhone   string `gorm:"column:staffphone;"`
	StaffGroupId int64  `gorm:"column:staffgroupid;index:staffgroupid"`
	DataStatus   int64  `gorm:"column:datastatus"`
	CreateTime   int64  `gorm:"column:createtime"`
	UpdateTime   int64  `gorm:"column:updatetime"`
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

func (this *_StaffDao) Add(name string, no string, staffDesc string, phone string) (*Staff, error) {
	obj := &Staff{
		StaffNO:    no,
		StaffName:  name,
		StaffDecs:  staffDesc,
		StaffPhone: phone,
		DataStatus: 0,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}

	err := this.Db.Create(obj).Error
	return obj, err
}

func (this *_StaffDao) GetAndExistByNO(no string) (bool, *Staff, error) {
	obj := &Staff{}
	err := this.Db.Where("staffno = ? and datastatus = ?", no, 0).First(obj).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, obj, nil
		}
		return true, obj, err
	}

	return true, obj, nil
}

func (this *_StaffDao) List(key string, limit, offset int64) ([]*Staff, error) {
	objList := make([]*Staff, 0)
	var err error
	key = "%" + key + "%"

	err = this.Db.Where("staffname like ? and datastatus = ? ", key, 0).Offset(offset).Limit(limit).Order("createtime").Find(&objList).Error

	return objList, err
}

func (this *_StaffDao) Count(key string) (int64, error) {
	var err error
	key = "%" + key + "%"

	var n int64
	err = this.Db.Where("staffname like ? and datastatus = ? ", key, 0).Count(&n).Error

	return n, err
}

func (this *_StaffDao) Get(id int64) (*Staff, error) {
	obj := &Staff{}
	err := this.Db.Where("staffid = ? and datastatus = ?", id, 0).First(obj).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (this *_StaffDao) Set(obj *Staff) error {
	return this.Db.Save(obj).Error
}
