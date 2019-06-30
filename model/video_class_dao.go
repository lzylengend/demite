package model

import (
	"github.com/jinzhu/gorm"
)

type VideoClass struct {
	Id         int64  `gorm:"column:id;primary_key"`
	Name       string `gorm:"column:name"`
	DataStatus int64  `gorm:"column:datastatus"`
	CreateTime int64  `gorm:"column:createtime"`
	UpdateTime int64  `gorm:"column:updatetime"`
}

type _VideoClassDao struct {
	Db *gorm.DB
}

func (VideoClass) TableName() string {
	return "VideoClass"
}

func newVideoClassDao(db *gorm.DB) *_VideoClassDao {
	db.AutoMigrate(&VideoClass{})

	return &_VideoClassDao{Db: db.Model(&VideoClass{})}
}

func (this *_VideoClassDao) Init() error {
	err := this.Db.AddUniqueIndex("idx_user_name_datastatus", "name", "datastatus").Error
	if err != nil {
		return err
	}

	return err
}

func (this *_VideoClassDao) Add(c *VideoClass) error {
	return this.Db.Create(c).Error
}

func (this *_VideoClassDao) List() ([]*VideoClass, error) {
	res := []*VideoClass{}
	err := this.Db.Where("datastatus = ?", 0).Find(&res).Error
	return res, err
}

func (this *_VideoClassDao) Update(c *VideoClass) error {
	return this.Db.Save(c).Error
}

func (this *_VideoClassDao) Get(id int64) (*VideoClass, error) {
	res := &VideoClass{}
	err := this.Db.Where("id = ?", id).First(&res).Error
	return res, err
}
