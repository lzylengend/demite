package model

import (
	"github.com/jinzhu/gorm"
)

type MaterialClass struct {
	Id         int64  `gorm:"column:id;primary_key"`
	Name       string `gorm:"column:name"`
	DataStatus int64  `gorm:"column:datastatus"`
	CreateTime int64  `gorm:"column:createtime"`
	UpdateTime int64  `gorm:"column:updatetime"`
}

type _MaterialClassDao struct {
	Db *gorm.DB
}

func (MaterialClass) TableName() string {
	return "MaterialClass"
}

func newMaterialClassDao(db *gorm.DB) *_MaterialClassDao {
	db.AutoMigrate(&MaterialClass{})

	return &_MaterialClassDao{Db: db.Model(&MaterialClass{})}
}

func (this *_MaterialClassDao) Init() error {
	err := this.Db.AddUniqueIndex("idx_user_name_datastatus", "name", "datastatus").Error
	if err != nil {
		return err
	}

	return err
}

func (this *_MaterialClassDao) Add(c *MaterialClass) error {
	return this.Db.Create(c).Error
}

func (this *_MaterialClassDao) List() ([]*MaterialClass, error) {
	res := []*MaterialClass{}
	err := this.Db.Where("datastatus = ?", 0).Find(&res).Error
	return res, err
}

func (this *_MaterialClassDao) Update(c *MaterialClass) error {
	return this.Db.Save(c).Error
}

func (this *_MaterialClassDao) Get(id int64) (*MaterialClass, error) {
	res := &MaterialClass{}
	err := this.Db.Where("id = ?", id).First(&res).Error
	return res, err
}
