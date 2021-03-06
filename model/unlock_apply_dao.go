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
	Creater    int64              `gorm:"column:creater"`
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

func (this *_UnlockApplyDao) Get(id int64) (*UnlockApply, error) {
	obj := &UnlockApply{}
	err := this.Db.Where("datastatus  = ? and id = ?", 0, id).First(obj).Error
	return obj, err
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

func (this *_UnlockApplyDao) DealApply(id int64, agree bool, userId int64) error {
	obj, err := UnlockApplyDao.Get(id)
	if err != nil {
		return err
	}

	_, gwObj, err := GoodsWXUserDao.GetAndExist(obj.GoodsUUID, obj.WXUserId)
	if err != nil {
		return err
	}

	status := GOODSWXUSERREFUSE
	if agree {
		status = GOODSWXUSERUNLOCK
	}

	obj.Status = status
	obj.Creater = userId
	obj.UpdateTime = time.Now().Unix()

	gwObj.UpdateTime = time.Now().Unix()
	gwObj.Status = status

	tx := this.Db.Begin()

	err = tx.Save(obj).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Save(gwObj).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (this *_UnlockApplyDao) ListByGoodUUIdWxUserIdStatus(goodName string, wxUserName string, limit int64, offsert int64, status string) ([]*UnlockApply, error) {
	sql := `datastatus  = ?`
	args := make([]interface{}, 0)
	args = append(args, 0)

	if status != "" {
		sql = sql + ` and status = ? `
		args = append(args, status)
	}

	goodIdList := make([]string, 0)
	if goodName != "" {
		goodList, err := GoodsDao.ListByQRCode(goodName, 99999, 0, "", "")
		if err != nil {
			return nil, err
		}

		for _, v := range goodList {
			goodIdList = append(goodIdList, v.GoodsUUID)
		}

		if len(goodIdList) != 0 {
			sql = sql + " and goodsuuid in (?)"
			args = append(args, goodIdList)
		}
	}

	wxUserIdList := make([]int64, 0)
	if wxUserName != "" {
		wxUserList, err := WxUserDao.List(wxUserName, 99999, 0)
		if err != nil {
			return nil, err
		}

		for _, v := range wxUserList {
			wxUserIdList = append(wxUserIdList, v.WxUserId)
		}

		if len(wxUserIdList) != 0 {
			sql = sql + " and wxuserid in (?)"
			args = append(args, wxUserIdList)
		}
	}

	objList := make([]*UnlockApply, 0)
	err := this.Db.Where(sql, args...).Order("createtime").Find(&objList).Limit(limit).Offset(offsert).Error

	return objList, err
}

func (this *_UnlockApplyDao) CountByGoodUUIdWxUserIdStatus(goodName string, wxUserName string, status string) (int64, error) {
	sql := `datastatus  = ?`
	args := make([]interface{}, 0)
	args = append(args, 0)

	if status != "" {
		sql = sql + ` and status = ? `
		args = append(args, status)
	}

	goodIdList := make([]string, 0)
	if goodName != "" {
		goodList, err := GoodsDao.ListByQRCode(goodName, 99999, 0, "", "")
		if err != nil {
			return 0, err
		}

		for _, v := range goodList {
			goodIdList = append(goodIdList, v.GoodsUUID)
		}

		if len(goodIdList) != 0 {
			sql = sql + " and goodsuuid in (?)"
			args = append(args, goodIdList)
		}
	}

	wxUserIdList := make([]int64, 0)
	if wxUserName != "" {
		wxUserList, err := WxUserDao.List(wxUserName, 99999, 0)
		if err != nil {
			return 0, err
		}

		for _, v := range wxUserList {
			wxUserIdList = append(wxUserIdList, v.WxUserId)
		}

		if len(wxUserIdList) != 0 {
			sql = sql + " and wxuserid in (?)"
			args = append(args, wxUserIdList)
		}
	}

	var n int64
	err := this.Db.Where(sql, args...).Count(&n).Error

	return n, err
}
