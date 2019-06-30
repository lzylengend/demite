package model

import (
	"github.com/jinzhu/gorm"
)

type Video struct {
	Id         int64  `gorm:"column:id;primary_key"`
	FileId     string `gorm:"column:fileid"`
	ClassId    int64  `gorm:"column:classid"`
	Hot        bool   `gorm:"column:hot"`
	Carousel   bool   `gorm:"column:carousel"`
	Title      string `gorm:"column:title"`
	Desc       string `gorm:"column:desc"`
	DataStatus int64  `gorm:"column:datastatus"`
	CreateTime int64  `gorm:"column:createtime"`
	UpdateTime int64  `gorm:"column:updatetime"`
}

type _VideoDao struct {
	Db *gorm.DB
}

func (Video) TableName() string {
	return "Video"
}

func newVideoDao(db *gorm.DB) *_VideoDao {
	db.AutoMigrate(&Video{})

	return &_VideoDao{Db: db.Model(&Video{})}
}

func (this *_VideoDao) Add(c *Video) error {
	return this.Db.Create(c).Error
}

func (this *_VideoDao) List(classId int64, limit int64, offset int64) ([]*Video, error) {
	res := []*Video{}
	var err error

	if classId == 0 {
		err = this.Db.Where("datastatus = ?", 0).Limit(limit).Offset(offset).Find(&res).Error
	} else {
		err = this.Db.Where("datastatus = ? and classid = ?", 0, classId).Limit(limit).Offset(offset).Order("updatetime desc").Find(&res).Error
	}

	return res, err
}

func (this *_VideoDao) ListHot() ([]*Video, error) {
	res := []*Video{}
	err := this.Db.Where("datastatus = ? and hot = ?", 0, true).Find(&res).Error
	return res, err
}

func (this *_VideoDao) ListCarousel() ([]*Video, error) {
	res := []*Video{}
	err := this.Db.Where("datastatus = ? and carousel = ?", 0, true).Find(&res).Error
	return res, err
}

func (this *_VideoDao) Update(c *Video) error {
	return this.Db.Save(c).Error
}

func (this *_VideoDao) Get(id int64) (*Video, error) {
	res := &Video{}
	err := this.Db.Where("id = ?", id).First(&res).Error
	return res, err
}
