package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type remoteStatus string

const (
	REMOTESTATUSAPPLY   = "applying"
	REMOTESTATUSCOMFIRM = "comfirm"
	REMOTESTATUSFINISH  = "finish"
)

type Remote struct {
	RemoteId   int64        `gorm:"column:remoteid;primary_key;AUTO_INCREMENT"`
	WXUserId   int64        `gorm:"column:wxuserid;index:wxuserid"`
	Hospital   string       `gorm:"column:hospital"`
	Office     string       `gorm:"column:office"`
	Phone      string       `gorm:"column:phone"`
	Name       string       `gorm:"column:name"`
	FaultDesc  string       `gorm:"column:faultdesc"`
	FaultType  string       `gorm:"column:faulttype"`
	CreateTime int64        `gorm:"column:createtime"`
	UpdateTime int64        `gorm:"column:updatetime"`
	FileId1    string       `gorm:"column:fileid1"`
	FileId2    string       `gorm:"column:fileid2"`
	Status     remoteStatus `gorm:"column:status"`
	DataStatus int64        `gorm:"column:datastatus"`
}
type _RemoteDao struct {
	Db *gorm.DB
}

func (Remote) TableName() string {
	return "Remote"
}

func newRemoteDao(db *gorm.DB) *_RemoteDao {
	db.AutoMigrate(&Remote{})

	return &_RemoteDao{Db: db.Model(&Remote{})}
}

func (this *_RemoteDao) Apply(phone string, name string, hospital string, office string,
	faultdesc string, faulttype string, fileid1 string, fileid2 string, wxuserid int64) (*Remote, error) {
	obj := &Remote{
		WXUserId:   wxuserid,
		Phone:      phone,
		Name:       name,
		Hospital:   hospital,
		Office:     office,
		FaultDesc:  faultdesc,
		FaultType:  faulttype,
		CreateTime: time.Now().Unix(),
		FileId1:    fileid1,
		FileId2:    fileid2,
		Status:     REMOTESTATUSAPPLY,
		DataStatus: 0,
	}

	tx := this.Db.Begin()
	err := tx.Create(obj).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	objSchedule := &RemoteSchedule{
		RemoteId:   obj.RemoteId,
		CreateId:   0,
		WxUserId:   wxuserid,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		Status:     REMOTESTATUSAPPLY,
		DataStatus: 0,
	}
	err = tx.Create(objSchedule).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return obj, nil
}

func (this *_RemoteDao) List(name string, limit int64, offset int64) ([]*Remote, error) {
	objList := make([]*Remote, 0)
	var err error
	name = "%" + name + "%"

	err = this.Db.Where("name like ? and datastatus = ? ", name, 0).Offset(offset).Limit(limit).Order("createtime").Find(&objList).Error

	return objList, err
}

func (this *_RemoteDao) Count(name string) (int64, error) {
	var n int64
	var err error
	name = "%" + name + "%"

	err = this.Db.Where("name like ? and datastatus = ? ", name, 0).Count(&n).Error

	return n, err
}

func (this *_RemoteDao) ListByWxUserIdAndGoodUUID(wxUserId int64, limit int64, offset int64, gooduuid string) ([]*Remote, error) {
	objList := make([]*Remote, 0)

	err := this.Db.Where("wxuserid = ? and datastatus = ? and gooduuid = ?", wxUserId, 0, gooduuid).Offset(offset).Limit(limit).Order("createtime").Find(&objList).Error

	return objList, err
}

func (this *_RemoteDao) CountByWxUserIdAndGoodUUID(wxUserId int64, gooduuid string) (int64, error) {
	var n int64
	err := this.Db.Where("wxuserid = ? and datastatus = ? and gooduuid = ?", wxUserId, 0, gooduuid).Count(&n).Error

	return n, err
}

func (this *_RemoteDao) Get(id int64) (*Remote, error) {
	obj := &Remote{}
	err := this.Db.Where("remoteid = ? and datastatus = ? ", id, 0).First(obj).Error
	return obj, err
}

func (this *_RemoteDao) GetByWIdAndxUserId(id, wxUserId int64) (*Remote, error) {
	obj := &Remote{}
	err := this.Db.Where("wxuserid = ? and remoteid = ? and datastatus = ? ", wxUserId, id, 0).First(obj).Error
	return obj, err
}

func (this *_RemoteDao) Deal(id int64, userId int64, staffId int64, remoteTime int64) error {
	obj, err := this.Get(id)
	if err != nil {
		return err
	}

	obj.Status = REMOTESTATUSCOMFIRM
	obj.UpdateTime = time.Now().Unix()

	tx := this.Db.Begin()
	err = tx.Save(obj).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	objSchedule := &RemoteSchedule{
		RemoteId:   obj.RemoteId,
		CreateId:   userId,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		Status:     REMOTESTATUSCOMFIRM,
		StaffId:    staffId,
		RemoteTime: remoteTime,
		DataStatus: 0,
	}
	err = tx.Create(objSchedule).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (this *_RemoteDao) Finish(id int64, wxUserId int64) error {
	obj, err := this.Get(id)
	if err != nil {
		return err
	}

	obj.Status = REMOTESTATUSFINISH
	obj.UpdateTime = time.Now().Unix()

	tx := this.Db.Begin()
	err = tx.Save(obj).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	objSchedule := &RemoteSchedule{
		RemoteId:   obj.RemoteId,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		Status:     REMOTESTATUSFINISH,
		WxUserId:   wxUserId,
		DataStatus: 0,
	}
	err = tx.Create(objSchedule).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
