package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type WxUser struct {
	WxUserId   int64  `gorm:"column:wxuserid;primary_key;AUTO_INCREMENT"`
	OpenId     string `gorm:"column:openid;index:openid;unique_index"`
	SessionKey string `gorm:"column:sessionkey"`
	NickName   string `gorm:"column:nickname"`
	Gender     string `gorm:"column:gender"`
	City       string `gorm:"column:city"`
	Province   string `gorm:"column:province"`
	AvatarUrl  string `gorm:"column:avatarUrl;type:varchar(300)"`
	Country    string `gorm:"column:country"`
	DataStatus int64  `gorm:"column:datastatus"`
	CreateTime int64  `gorm:"column:createtime"`
	UpdateTime int64  `gorm:"column:updatetime"`
}
type _WxUserDao struct {
	Db *gorm.DB
}

func (WxUser) TableName() string {
	return "WxUser"
}

func newWxUserDao(db *gorm.DB) *_WxUserDao {
	db.AutoMigrate(&WxUser{})

	return &_WxUserDao{Db: db.Model(&WxUser{})}
}

func (this *_WxUserDao) GetByOpenid(openId string) (*WxUser, error) {
	obj := &WxUser{}
	err := this.Db.Where("openid = ?", openId).First(obj).Error
	return obj, err
}

func (this *_WxUserDao) GetById(id int64) (*WxUser, error) {
	obj := &WxUser{}
	err := this.Db.Where("wxuserid = ?", id).First(obj).Error
	return obj, err
}

func (this *_WxUserDao) ExistOpenid(openId string) (*WxUser, bool, error) {
	obj, err := this.GetByOpenid(openId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return obj, false, nil
		}
		return obj, true, err
	}
	return obj, true, nil
}

func (this *_WxUserDao) AddWxUser(openId, sessionKey, nickName, gender, city, province, avatarUrl, country string) (*WxUser, error) {
	obj := this.NewWxUser(openId, sessionKey, nickName, gender, city, province, avatarUrl, country)

	err := this.Db.Create(obj).Error
	return obj, err
}

func (this *_WxUserDao) Set(obj *WxUser) error {
	err := this.Db.Save(obj).Error
	return err
}

func (this *_WxUserDao) NewWxUser(openId, sessionKey, nickName, gender, city, province, avatarUrl, country string) *WxUser {
	return &WxUser{
		OpenId:     openId,
		SessionKey: sessionKey,
		NickName:   nickName,
		Gender:     gender,
		City:       city,
		Province:   province,
		AvatarUrl:  avatarUrl,
		Country:    country,
		DataStatus: 0,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
}

func (this *_WxUserDao) List(key string, limit int64, offset int64) ([]*WxUser, error) {
	objList := make([]*WxUser, 0)
	var err error
	key = "%" + key + "%"

	err = this.Db.Where("nickname like ? and datastatus = ? ", key, 0).Offset(offset).Limit(limit).Order("createtime").Find(&objList).Error

	return objList, err
}

func (this *_WxUserDao) Count(key string) (int64, error) {
	var n int64
	key = "%" + key + "%"

	err := this.Db.Where("nickname like ? and datastatus = ? ", key, 0).Count(&n).Error
	return n, err
}
