package model

import (
	"github.com/jinzhu/gorm"
)

type Material struct {
	Id         int64  `gorm:"column:id;primary_key"`
	FileId     string `gorm:"column:fileid"`
	ClassId    int64  `gorm:"column:classid"`
	Title      string `gorm:"column:title"`
	Desc       string `gorm:"column:desc"`
	DataStatus int64  `gorm:"column:datastatus"`
	CreateTime int64  `gorm:"column:createtime"`
	UpdateTime int64  `gorm:"column:updatetime"`
}

type _MaterialDao struct {
	Db *gorm.DB
}

func (Material) TableName() string {
	return "Material"
}

func newMaterialDao(db *gorm.DB) *_MaterialDao {
	db.AutoMigrate(&Material{})

	return &_MaterialDao{Db: db.Model(&Material{})}
}

func (this *_MaterialDao) Add(c *Material) error {
	return this.Db.Create(c).Error
}

func (this *_MaterialDao) List(classId int64, limit int64, offset int64) ([]*Material, error) {
	res := []*Material{}
	var err error

	if classId == 0 {
		err = this.Db.Where("datastatus = ?", 0).Limit(limit).Offset(offset).Order("updatetime desc").Find(&res).Error
	} else {
		err = this.Db.Where("datastatus = ? and classid = ?", 0, classId).Limit(limit).Offset(offset).Order("updatetime desc").Find(&res).Error
	}

	return res, err
}

func (this *_MaterialDao) Update(c *Material) error {
	return this.Db.Save(c).Error
}

func (this *_MaterialDao) Get(id int64) (*Material, error) {
	res := &Material{}
	err := this.Db.Where("id = ?", id).First(&res).Error
	return res, err
}
