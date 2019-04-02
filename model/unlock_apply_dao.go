package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type UnlockApply struct {
	Id         int64              `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	GoodsUUID  string             `gorm:"column:goodsuuid;index:goodsuuid"`
	WXUserId   int64              `gorm:"column:wxuserid;index:wxuserid"`
	Status     goodsWXUserSatatus `gorm:"column:status"`
	DataStatus int64              `gorm:"column:datastatus"`
	CreateTime int64              `gorm:"column:createtime"`
	UpdateTime int64              `gorm:"column:updatetime"`
}
type _UnlockApplyDao struct {
	Db *gorm.DB
}

func (UnlockApply) TableName() string {
	return "unlockapply"
}

func newUnlockApplyDao(db *gorm.DB) *_UnlockApplyDao {
	db.AutoMigrate(&UnlockApply{})

	return &_UnlockApplyDao{Db: db.Model(&UnlockApply{})}
}

func (this *_UnlockApplyDao) GetByStatusAndExit(goodUUID string, wxUser int64, status goodsWXUserSatatus) (bool, *UnlockApply, error) {
	obj := &UnlockApply{}
	err := this.Db.Where("datastatus  = ? and goodsuuid = ? and wxuserid = ? and status = ?", 0, goodUUID, wxUser, status).First(obj).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, obj, nil
		}
		return false, obj, err
	}
	return true, obj, nil
}

func (this *_UnlockApplyDao) CountByNotStatus(goodUUID string, wxUser int64, status goodsWXUserSatatus) (int64, error) {
	var n int64
	err := this.Db.Where("datastatus  = ? and goodsuuid = ? and wxuserid = ? and status <> ?", 0, goodUUID, wxUser, status).Count(&n).Error
	return n, err
}

func (this *_UnlockApplyDao) Apply(goodUUID string, wxid int64, gw *GoodsWXUser) error {
	obj := &UnlockApply{
		GoodsUUID:  goodUUID,
		WXUserId:   wxid,
		Status:     GOODSWXUSERAPPLYING,
		DataStatus: 0,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}

	gw.Status = GOODSWXUSERAPPLYING
	gw.UpdateTime = time.Now().Unix()

	tx := this.Db.Begin()

	err := tx.Create(obj).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Save(gw).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
