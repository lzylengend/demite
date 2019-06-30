package model

import (
	"github.com/jinzhu/gorm"
)

type Soft struct {
	Id         int64  `gorm:"column:id;primary_key"`
	Content    string `gorm:"column:content;type:text"`
	ClassId    int64  `gorm:"column:classid"`
	Title      string `gorm:"column:title"`
	Desc       string `gorm:"column:desc"`
	DataStatus int64  `gorm:"column:datastatus"`
	CreateTime int64  `gorm:"column:createtime"`
	UpdateTime int64  `gorm:"column:updatetime"`
}

type _SoftDao struct {
	Db *gorm.DB
}

func (Soft) TableName() string {
	return "Soft"
}

func newSoftDao(db *gorm.DB) *_SoftDao {
	db.AutoMigrate(&Soft{})

	return &_SoftDao{Db: db.Model(&Soft{})}
}

func (this *_SoftDao) Add(c *Soft) error {
	return this.Db.Create(c).Error
}

func (this *_SoftDao) List(classId int64, limit int64, offset int64) ([]*Soft, error) {
	res := []*Soft{}
	var err error

	if classId == 0 {
		err = this.Db.Where("datastatus = ? ", 0).Limit(limit).Offset(offset).Order("updatetime desc").Find(&res).Error
	} else {
		err = this.Db.Where("datastatus = ? and classid = ?", 0, classId).Limit(limit).Offset(offset).Order("updatetime desc").Find(&res).Error
	}

	return res, err
}

func (this *_SoftDao) Count(classId int64) (int64, error) {
	var n int64
	var err error

	if classId == 0 {
		err = this.Db.Where("datastatus = ? ", 0).Count(&n).Error
	} else {
		err = this.Db.Where("datastatus = ? and classid = ?", 0, classId).Count(&n).Error
	}

	return n, err
}

func (this *_SoftDao) Update(c *Soft) error {
	return this.Db.Save(c).Error
}

func (this *_SoftDao) Get(id int64) (*Soft, error) {
	res := &Soft{}
	err := this.Db.Where("id = ?", id).First(&res).Error
	return res, err
}
