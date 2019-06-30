package model

import (
	"time"

	//	"errors"

	"github.com/jinzhu/gorm"
)

type repairStatus string

const (
	REPAIRSTATUSAPPLY   = "applying"
	REPAIRSTATUSCOMFIRM = "comfirm"
	REPAIRSTATUSFINISH  = "finish"
	REPAIRSTATUSREFUSE  = "refuse"
)

type Repair struct {
	RepairId   int64        `gorm:"column:repairid;primary_key;AUTO_INCREMENT"`
	GoodUUID   string       `gorm:"column:gooduuid;index:gooduuid"`
	WXUserId   int64        `gorm:"column:wxuserid;index:wxuserid"`
	GoodModel  string       `gorm:"column:goodmodel"`
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
	Status     repairStatus `gorm:"column:status"`
	DataStatus int64        `gorm:"column:datastatus"`
}
type _RepairDao struct {
	Db *gorm.DB
}

func (Repair) TableName() string {
	return "Repair"
}

func newRepairDao(db *gorm.DB) *_RepairDao {
	db.AutoMigrate(&Repair{})

	return &_RepairDao{Db: db.Model(&Repair{})}
}

func (this *_RepairDao) Apply(gooduuid string, goodmodel string, phone string, name string, hospital string, office string,
	faultdesc string, faulttype string, fileid1 string, fileid2 string, wxuserid int64) (*Repair, error) {
	obj := &Repair{
		GoodUUID:   gooduuid,
		WXUserId:   wxuserid,
		GoodModel:  goodmodel,
		Phone:      phone,
		Name:       name,
		Hospital:   hospital,
		Office:     office,
		FaultDesc:  faultdesc,
		FaultType:  faulttype,
		CreateTime: time.Now().Unix(),
		FileId1:    fileid1,
		FileId2:    fileid2,
		Status:     REPAIRSTATUSAPPLY,
		DataStatus: 0,
	}

	tx := this.Db.Begin()
	err := tx.Create(obj).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	objSchedule := &RepairSchedule{
		RepairId:   obj.RepairId,
		CreateId:   0,
		WxUserId:   wxuserid,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		Status:     REPAIRSTATUSAPPLY,
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

func (this *_RepairDao) ListByStatus(name string, limit int64, offset int64, status string) ([]*Repair, error) {
	objList := make([]*Repair, 0)
	sql := `datastatus  = ?`
	args := make([]interface{}, 0)
	args = append(args, 0)

	if status != "" {
		sql = sql + ` and status = ? `
		args = append(args, status)
	}

	if name != "" {
		sql = sql + ` and name like ?`
		name = "%" + name + "%"
		args = append(args, name)
	}

	err := this.Db.Where(sql, args...).Offset(offset).Limit(limit).Order("createtime desc").Find(&objList).Error

	return objList, err
}

func (this *_RepairDao) CountByStatus(name string, status string) (int64, error) {
	var n int64
	sql := `datastatus  = ?`
	args := make([]interface{}, 0)
	args = append(args, 0)

	if status != "" {
		sql = sql + ` and status = ? `
		args = append(args, status)
	}

	if name != "" {
		sql = sql + ` and name like ?`
		name = "%" + name + "%"
		args = append(args, name)
	}

	err := this.Db.Where(sql, args...).Count(&n).Error

	return n, err
}

func (this *_RepairDao) ListByWxUserIdAndGoodUUID(wxUserId int64, limit int64, offset int64, gooduuid string) ([]*Repair, error) {
	objList := make([]*Repair, 0)

	err := this.Db.Where("wxuserid = ? and datastatus = ? and gooduuid = ?", wxUserId, 0, gooduuid).Offset(offset).Limit(limit).Order("createtime  desc").Find(&objList).Error

	return objList, err
}

func (this *_RepairDao) CountByWxUserIdAndGoodUUID(wxUserId int64, gooduuid string) (int64, error) {
	var n int64
	err := this.Db.Where("wxuserid = ? and datastatus = ? and gooduuid = ?", wxUserId, 0, gooduuid).Count(&n).Error

	return n, err
}

func (this *_RepairDao) Get(id int64) (*Repair, error) {
	obj := &Repair{}
	err := this.Db.Where("repairid = ? and datastatus = ? ", id, 0).First(obj).Error
	return obj, err
}

func (this *_RepairDao) GetByWIdAndxUserId(id, wxUserId int64) (*Repair, error) {
	obj := &Repair{}
	err := this.Db.Where("wxuserid = ? and repairid = ? and datastatus = ? ", wxUserId, id, 0).First(obj).Error
	return obj, err
}

func (this *_RepairDao) Refuse(id int64, userId int64, reason string) error {
	obj, err := this.Get(id)
	if err != nil {
		return err
	}

	obj.Status = REPAIRSTATUSREFUSE
	obj.UpdateTime = time.Now().Unix()

	tx := this.Db.Begin()
	err = tx.Save(obj).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	objSchedule := &RepairSchedule{
		RepairId:   obj.RepairId,
		CreateId:   userId,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		Status:     REPAIRSTATUSREFUSE,
		Reason:     reason,
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

func (this *_RepairDao) Deal(id int64, userId int64, staffId int64, repairTime int64) error {
	obj, err := this.Get(id)
	if err != nil {
		return err
	}

	obj.Status = REPAIRSTATUSCOMFIRM
	obj.UpdateTime = time.Now().Unix()

	// if repairTime < time.Now().Unix() {
	// 	return errors.New("time error")
	// }

	tx := this.Db.Begin()
	err = tx.Save(obj).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	objSchedule := &RepairSchedule{
		RepairId:   obj.RepairId,
		CreateId:   userId,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		Status:     REPAIRSTATUSCOMFIRM,
		StaffId:    staffId,
		RepairTime: repairTime,
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

func (this *_RepairDao) Finish(id int64, wxUserId int64) error {
	obj, err := this.Get(id)
	if err != nil {
		return err
	}

	obj.Status = REPAIRSTATUSFINISH
	obj.UpdateTime = time.Now().Unix()

	tx := this.Db.Begin()
	err = tx.Save(obj).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	objSchedule := &RepairSchedule{
		RepairId:   obj.RepairId,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		Status:     REPAIRSTATUSFINISH,
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
