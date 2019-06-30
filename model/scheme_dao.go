package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

const (
	SchemeTypeIntelligence = "intelligence"
	SchemeTypeItem         = "item"
)

type Scheme struct {
	Id         int64  `gorm:"column:id;primary_key"`
	Content    string `gorm:"column:content;type:text"`
	Title      string `gorm:"column:title"`
	Desc       string `gorm:"column:desc"`
	FileId     string `json:"fileid"`
	SchemeType string `gorm:"column:schemetype"`
	DataStatus int64  `gorm:"column:datastatus"`
	CreateTime int64  `gorm:"column:createtime"`
	UpdateTime int64  `gorm:"column:updatetime"`
}

type _SchemeDao struct {
	Db *gorm.DB
}

func (Scheme) TableName() string {
	return "Scheme"
}

func newSchemeDao(db *gorm.DB) *_SchemeDao {
	db.AutoMigrate(&Scheme{})

	return &_SchemeDao{Db: db.Model(&Scheme{})}
}

func (this *_SchemeDao) init() error {
	_, err := this.GetIntelligence()
	if err == gorm.ErrRecordNotFound {
		err = this.Db.Save(&Scheme{
			SchemeType: SchemeTypeIntelligence,
			DataStatus: 0,
			CreateTime: time.Now().Unix(),
			UpdateTime: time.Now().Unix(),
		}).Error
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	_, err = this.GetItem()
	if err == gorm.ErrRecordNotFound {
		err = this.Db.Save(&Scheme{
			SchemeType: SchemeTypeItem,
			DataStatus: 0,
			CreateTime: time.Now().Unix(),
			UpdateTime: time.Now().Unix(),
		}).Error
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}

func (this *_SchemeDao) Add(c *Scheme) error {
	return this.Db.Create(c).Error
}

func (this *_SchemeDao) List(limit int64, offset int64) ([]*Scheme, error) {
	res := []*Scheme{}
	var err error

	err = this.Db.Where("datastatus = ? and schemetype = ?", 0, "").Limit(limit).Offset(offset).Order("updatetime desc").Find(&res).Error

	return res, err
}

func (this *_SchemeDao) Count() (int64, error) {
	var n int64
	var err error

	err = this.Db.Where("datastatus = ? and schemetype = ?", 0, "").Count(&n).Error

	return n, err
}

func (this *_SchemeDao) Update(c *Scheme) error {
	return this.Db.Save(c).Error
}

func (this *_SchemeDao) Get(id int64) (*Scheme, error) {
	res := &Scheme{}
	err := this.Db.Where("id = ?", id).First(&res).Error
	return res, err
}

func (this *_SchemeDao) GetIntelligence() (*Scheme, error) {
	res := &Scheme{}
	err := this.Db.Where("schemetype = ?", SchemeTypeIntelligence).First(&res).Error
	return res, err
}

func (this *_SchemeDao) GetItem() (*Scheme, error) {
	res := &Scheme{}
	err := this.Db.Where("schemetype = ?", SchemeTypeItem).First(&res).Error
	return res, err
}
