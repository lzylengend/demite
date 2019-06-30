package model

import (
	"github.com/jinzhu/gorm"
)

type SoftClass struct {
	Id         int64  `gorm:"column:id;primary_key"`
	Name       string `gorm:"column:name"`
	DataStatus int64  `gorm:"column:datastatus"`
	CreateTime int64  `gorm:"column:createtime"`
	UpdateTime int64  `gorm:"column:updatetime"`
}

type _SoftClassDao struct {
	Db *gorm.DB
}

func (SoftClass) TableName() string {
	return "SoftClass"
}

func newSoftClassDao(db *gorm.DB) *_SoftClassDao {
	db.AutoMigrate(&SoftClass{})

	return &_SoftClassDao{Db: db.Model(&SoftClass{})}
}

func (this *_SoftClassDao) Init() error {
	err := this.Db.AddUniqueIndex("idx_user_name_datastatus", "name", "datastatus").Error
	if err != nil {
		return err
	}

	return err
}

func (this *_SoftClassDao) Add(c *SoftClass) error {
	return this.Db.Create(c).Error
}

func (this *_SoftClassDao) List() ([]*SoftClass, error) {
	res := []*SoftClass{}
	err := this.Db.Where("datastatus = ?", 0).Find(&res).Error
	return res, err
}

func (this *_SoftClassDao) Update(c *SoftClass) error {
	return this.Db.Save(c).Error
}

func (this *_SoftClassDao) Get(id int64) (*SoftClass, error) {
	res := &SoftClass{}
	err := this.Db.Where("id = ?", id).First(&res).Error
	return res, err
}
