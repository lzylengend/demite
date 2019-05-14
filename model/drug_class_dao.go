package model

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type DrugClass struct {
	ClassId    int64  `gorm:"column:classid;primary_key;AUTO_INCREMENT"`
	ClassName  string `gorm:"column:classname;index:classname"`
	UpClassId  int64  `gorm:"column:upclassid;index:upclassid"` //0为根目录
	Path       string `gorm:"column:path"`
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

func (this *_DrugClassDao) AddDrugClass(className string, upClassId int64, path string) (*DrugClass, error) {
	//obj := &DrugClass{
	//	ClassName:  className,
	//	DataStatus: 0,
	//	CreateTime: time.Now().Unix(),
	//	UpdateTime: time.Now().Unix(),
	//}
	//
	//err := this.Db.Create(obj).Error
	//return obj, err

	obj := &DrugClass{
		ClassName:  className,
		UpClassId:  upClassId,
		DataStatus: 0,
		Path:       path,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}

	s := this.Db.Begin()
	err := s.Create(obj).Error
	if err != nil {
		s.Rollback()
		return nil, err
	}

	obj.Path = this.AddPath(obj.Path, obj.ClassId)
	err = s.Save(obj).Error
	if err != nil {
		s.Rollback()
		return nil, err
	}

	s.Commit()

	return obj, err
}

func (this *_DrugClassDao) AddPath(path string, classId int64) string {
	return path + "/" + strconv.Itoa(int(classId))
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

//func (this *_DrugClassDao) Exist(id int64) (bool, error) {
//	_, err := this.Get(id)
//	if err != nil {
//		return true, err
//	}
//	return false, err
//}

func (this *_DrugClassDao) List() ([]*DrugClass, error) {
	objList := make([]*DrugClass, 0)
	err := this.Db.Where("datastatus = ?", 0).Find(&objList).Error
	return objList, err
}

func (this *_DrugClassDao) ListClassByUp(upClassId int64) ([]*Class, error) {
	objList := make([]*Class, 0)
	err := this.Db.Where("upclassid = ? and datastatus = ?", upClassId, 0).Find(&objList).Error
	return objList, err
}
