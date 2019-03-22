package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type DrugClass struct {
	ClassId    int64  `gorm:"column:classid;primary_key;AUTO_INCREMENT"`
	ClassName  string `gorm:"column:classname;index:classname"`
	DataStatus int64  `gorm:"column:datastatus"`
	CreateTime int64  `gorm:"column:createtime"`
	UpdateTime int64  `gorm:"column:updatetime"`
}
type _DrugClassDao struct {
	Db *gorm.DB
}

func (DrugClass) TableName() string {
	return "drugclass"
}

func newDrugClassDao(db *gorm.DB) *_DrugClassDao {
	db.AutoMigrate(&DrugClass{})

	return &_DrugClassDao{Db: db.Model(&DrugClass{})}
}

func (this *_DrugClassDao) AddDrugClass(className string) (*DrugClass, error) {
	obj := &DrugClass{
		ClassName:  className,
		DataStatus: 0,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}

	err := this.Db.Create(obj).Error
	return obj, err
}

func (this *_DrugClassDao) Set(obj *DrugClass) error {
	obj.UpdateTime = time.Now().Unix()
	return this.Db.Save(obj).Error
}

func (this *_DrugClassDao) Get(id int64) (*DrugClass, error) {
	obj := &DrugClass{}
	err := this.Db.Where("classid = ? and datastatus = ?", id, 0).First(obj).Error
	return obj, err
}
