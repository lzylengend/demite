package model

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type Class struct {
	ClassId    int64  `gorm:"column:classid;primary_key;AUTO_INCREMENT"`
	ClassName  string `gorm:"column:classname;index:classname"`
	UpClassId  int64  `gorm:"column:upclassid;index:upclassid"` //0为根目录
	IsShow     int64  `gorm:"column:isshow;index:isshow"`
	Path       string `gorm:"column:path"`
	DataStatus int64  `gorm:"column:datastatus"`
	CreateTime int64  `gorm:"column:createtime"`
	UpdateTime int64  `gorm:"column:updatetime"`
}

type _ClassDao struct {
	Db *gorm.DB
}

func (Class) TableName() string {
	return "class"
}

func newClassDao(db *gorm.DB) *_ClassDao {
	db.AutoMigrate(&Class{})

	return &_ClassDao{Db: db.Model(&Class{})}
}

func (this *_ClassDao) AddClass(className string, upClassId int64, path string) (*Class, error) {
	obj := &Class{
		ClassName:  className,
		UpClassId:  upClassId,
		DataStatus: 0,
		IsShow:     time.Now().Unix(),
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

func (this *_ClassDao) Insert(obj *Class) error {
	return this.Db.Create(obj).Error
}

func (this *_ClassDao) Set(obj *Class) error {
	obj.UpdateTime = time.Now().Unix()
	return this.Db.Save(obj).Error
}

func (this *_ClassDao) GetClassById(classId int64) (*Class, error) {
	obj := &Class{}
	err := this.Db.Where("classid = ? and datastatus = ?", classId, 0).First(obj).Error
	return obj, err
}

func (this *_ClassDao) ListClassByUp(upClassId int64) ([]*Class, error) {
	objList := make([]*Class, 0)
	err := this.Db.Where("upclassid = ? and datastatus = ?", upClassId, 0).Find(&objList).Error
	return objList, err
}

func (this *_ClassDao) AddPath(path string, classId int64) string {
	return path + "/" + strconv.Itoa(int(classId))
}

func (this *_ClassDao) ExistId(Id int64) (*Class, bool, error) {
	obj, err := this.GetClassById(Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return obj, false, nil
		}
		return obj, true, err
	}
	return obj, true, nil
}
