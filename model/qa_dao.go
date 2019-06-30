package model

import (
	"github.com/jinzhu/gorm"
)

type QA struct {
	Id         int64  `gorm:"column:id;primary_key"`
	Content    string `gorm:"column:content;type:text"`
	Title      string `gorm:"column:title"`
	Desc       string `gorm:"column:desc"`
	DataStatus int64  `gorm:"column:datastatus"`
	CreateTime int64  `gorm:"column:createtime"`
	UpdateTime int64  `gorm:"column:updatetime"`
}

type _QADao struct {
	Db *gorm.DB
}

func (QA) TableName() string {
	return "QA"
}

func newQADao(db *gorm.DB) *_QADao {
	db.AutoMigrate(&QA{})

	return &_QADao{Db: db.Model(&QA{})}
}

func (this *_QADao) Add(c *QA) error {
	return this.Db.Create(c).Error
}

func (this *_QADao) List(classId int64, limit int64, offset int64) ([]*QA, error) {
	res := []*QA{}
	var err error

	if classId == 0 {
		err = this.Db.Where("datastatus = ? ", 0).Limit(limit).Offset(offset).Order("updatetime desc").Find(&res).Error
	} else {
		err = this.Db.Where("datastatus = ? and classid = ?", 0, classId).Limit(limit).Offset(offset).Order("updatetime desc").Find(&res).Error
	}

	return res, err
}

func (this *_QADao) Count(classId int64) (int64, error) {
	var n int64
	var err error

	if classId == 0 {
		err = this.Db.Where("datastatus = ? ", 0).Count(&n).Error
	} else {
		err = this.Db.Where("datastatus = ? and classid = ?", 0, classId).Count(&n).Error
	}

	return n, err
}

func (this *_QADao) Update(c *QA) error {
	return this.Db.Save(c).Error
}

func (this *_QADao) Get(id int64) (*QA, error) {
	res := &QA{}
	err := this.Db.Where("id = ?", id).First(&res).Error
	return res, err
}
