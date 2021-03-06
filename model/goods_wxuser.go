package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type goodsWXUserSatatus string

const (
	GOODSWXUSERLOCK     goodsWXUserSatatus = "lock"
	GOODSWXUSERAPPLYING goodsWXUserSatatus = "applying"
	GOODSWXUSERUNLOCK   goodsWXUserSatatus = "unlock"
	GOODSWXUSERREFUSE   goodsWXUserSatatus = "refuse"
)

type GoodsWXUser struct {
	Id         int64              `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	GoodsUUID  string             `gorm:"column:goodsuuid;index:goodsuuid"`
	WXUserId   int64              `gorm:"column:wxuserid;index:wxuserid"`
	Status     goodsWXUserSatatus `gorm:"column:status"`
	DataStatus int64              `gorm:"column:datastatus"`
	CreateTime int64              `gorm:"column:createtime"`
	UpdateTime int64              `gorm:"column:updatetime"`
}
type _GoodsWXUserDao struct {
	Db *gorm.DB
}

func (GoodsWXUser) TableName() string {
	return "goodswxuSser"
}

func newGoodsWXUserDao(db *gorm.DB) *_GoodsWXUserDao {
	db.AutoMigrate(&GoodsWXUser{})

	return &_GoodsWXUserDao{Db: db.Model(&GoodsWXUser{})}
}

func (this *_GoodsWXUserDao) Add(goodUUID string, wxUserId int64) error {
	obj := &GoodsWXUser{
		GoodsUUID:  goodUUID,
		WXUserId:   wxUserId,
		Status:     GOODSWXUSERLOCK,
		DataStatus: 0,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	return this.Db.Create(obj).Error
}

func (this *_GoodsWXUserDao) Set() {

}

func (this *_GoodsWXUserDao) GetAndExist(goodUUID string, wxUserId int64) (bool, *GoodsWXUser, error) {
	obj := &GoodsWXUser{
		Status: GOODSWXUSERLOCK,
	}
	err := this.Db.Where("datastatus  = ? and goodsuuid = ? and wxuserid = ?", 0, goodUUID, wxUserId).First(obj).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, obj, nil
		}
		return false, obj, err
	}
	return true, obj, nil
}

func (this *_GoodsWXUserDao) ListByUUID(uuid string) ([]*GoodsWXUser, error) {
	objList := make([]*GoodsWXUser, 0)
	err := this.Db.Where("datastatus  = ? and goodsuuid = ?", 0, uuid).Find(&objList).Error

	return objList, err
}

func (this *_GoodsWXUserDao) ListByWXId(wxId int64, limit int64, offset int64) ([]*GoodsWXUser, error) {
	objList := make([]*GoodsWXUser, 0)
	err := this.Db.Where("datastatus  = ? and wxuserid = ?", 0, wxId).Order("createtime desc").Offset(offset).Limit(limit).Find(&objList).Error

	return objList, err
}

func (this *_GoodsWXUserDao) CountByWXId(wxId int64) (int64, error) {
	var n int64
	err := this.Db.Where("datastatus  = ? and wxuserid = ?", 0, wxId).Count(&n).Error

	return n, err
}
