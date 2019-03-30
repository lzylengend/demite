package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Drug struct {
	DrugId                int64  `gorm:"column:drugid;primary_key;AUTO_INCREMENT"`
	DrugName              string `gorm:"column:drugname;index:drugname"`
	DrugClassId           int64  `gorm:"column:drugclassid"`
	Reagent               string `gorm:"column:reagent;type:text"`               //试剂
	ChromatographicColumn string `gorm:"column:chromatographiccolumn;type:text"` //色谱柱
	Controls              string `gorm:"column:controls;type:text"`              //质控品
	TestMethod            string `gorm:"column:testmethod;type:text"`            //检测方法
	DataStatus            int64  `gorm:"column:datastatus"`
	CreateTime            int64  `gorm:"column:createtime"`
	UpdateTime            int64  `gorm:"column:updatetime"`
}
type _DrugDao struct {
	Db *gorm.DB
}

func (Drug) TableName() string {
	return "drug"
}

func newDrugDao(db *gorm.DB) *_DrugDao {
	db.AutoMigrate(&Drug{})

	return &_DrugDao{Db: db.Model(&Drug{})}
}

func (this *_DrugDao) AddDrug(name string, classId int64, reagent string, chromatographiccolumn string, controls string, testmethod string) (*Drug, error) {
	obj := &Drug{
		DrugName:              name,
		DrugClassId:           classId,
		Reagent:               reagent,
		ChromatographicColumn: chromatographiccolumn,
		Controls:              controls,
		TestMethod:            testmethod,
		DataStatus:            0,
		CreateTime:            time.Now().Unix(),
		UpdateTime:            time.Now().Unix(),
	}

	err := this.Db.Create(obj).Error
	return obj, err
}

func (this *_DrugDao) Set(obj *Drug) error {
	obj.UpdateTime = time.Now().Unix()
	return this.Db.Save(obj).Error
}

func (this *_DrugDao) Get(id int64) (*Drug, error) {
	obj := &Drug{}
	err := this.Db.Where("drugid = ? and datastatus = ?", id, 0).First(obj).Error
	return obj, err
}

func (this *_DrugDao) ListByCreateTime(key string, limit, offset int64) ([]*Drug, error) {
	objList := make([]*Drug, 0)

	key = "%" + key + "%"
	err := this.Db.Where("drugname like ? and datastatus = ?", key, 0).Offset(offset).Limit(limit).Order("createtime").Find(&objList).Error
	return objList, err
}

func (this *_DrugDao) CountByKey(key string) (int64, error) {
	n := 0

	key = "%" + key + "%"
	err := this.Db.Where("drugname like ? and datastatus = ?", key, 0).Count(&n).Error
	return int64(n), err
}
