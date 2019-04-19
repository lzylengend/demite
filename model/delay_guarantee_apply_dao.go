package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type delayGuaranteeSatatus string

const (
	DELAYGUARANTEEAPPLYNG = "applying"
	DELAYGUARANTEECOMFIRM = "comfirm"
	DELAYGUARANTEEREFUSE  = "refuse"
)

type DelayGuaranteeApply struct {
	Id         int64              `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	GoodsUUID  string             `gorm:"column:goodsuuid;index:goodsuuid"`
	WXUserId   int64              `gorm:"column:wxuserid;index:wxuserid"`
	Status     goodsWXUserSatatus `gorm:"column:status"`
	SourceTime int64              `gorm:"column:sourcetime"`
	DelayTime  int64              `gorm:"column:delaytime"`
	Creater    int64              `gorm:"column:creater"`
	DataStatus int64              `gorm:"column:datastatus"`
	CreateTime int64              `gorm:"column:createtime"`
	UpdateTime int64              `gorm:"column:updatetime"`
}
type _DelayGuaranteeApplyDao struct {
	Db *gorm.DB
}

func (DelayGuaranteeApply) TableName() string {
	return "delpayguranteeapply"
}

func newDelayGuaranteeApplyDao(db *gorm.DB) *_DelayGuaranteeApplyDao {
	db.AutoMigrate(&DelayGuaranteeApply{})

	return &_DelayGuaranteeApplyDao{Db: db.Model(&DelayGuaranteeApply{})}
}

func (this *_DelayGuaranteeApplyDao) Add(goodUUID string, wxUserId int64) (*DelayGuaranteeApply, error) {
	g, err := GoodsDao.GetByUUID(goodUUID)
	if err != nil {
		return nil, err
	}

	obj := &DelayGuaranteeApply{
		GoodsUUID:  goodUUID,
		WXUserId:   wxUserId,
		Status:     DELAYGUARANTEEAPPLYNG,
		SourceTime: g.GuaranteeTime,
		DataStatus: 0,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}

	err = this.Db.Create(obj).Error
	return obj, err
}

func (this *_DelayGuaranteeApplyDao) CountByGoodUUIDWXUserIdStatus(goodUUID string, wxUserId int64, status delayGuaranteeSatatus) (int64, error) {
	var n int64
	err := this.Db.Where("goodsuuid  = ? and wxuserid = ? and datastatus = ? and status = ?", goodUUID, wxUserId, 0, status).Count(&n).Error
	return n, err
}

func (this *_DelayGuaranteeApplyDao) ListByGoodUUIdWxUserIdStatus(goodName string, wxUserName string, limit int64, offsert int64, status string) ([]*DelayGuaranteeApply, error) {
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

		if len(goodList) == 0 {
			return nil, nil
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

		if len(wxUserList) == 0 {
			return nil, nil
		}

		for _, v := range wxUserList {
			wxUserIdList = append(wxUserIdList, v.WxUserId)
		}

		if len(wxUserIdList) != 0 {
			sql = sql + " and wxuserid in (?)"
			args = append(args, wxUserIdList)
		}
	}

	objList := make([]*DelayGuaranteeApply, 0)
	err := this.Db.Where(sql, args...).Order("createtime").Find(&objList).Limit(limit).Offset(offsert).Error

	return objList, err
}

func (this *_DelayGuaranteeApplyDao) CountByGoodUUIdWxUserIdStatus(goodName string, wxUserName string, status string) (int64, error) {
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

		if len(goodList) == 0 {
			return 0, nil
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

		if len(wxUserList) == 0 {
			return 0, nil
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

func (this *_DelayGuaranteeApplyDao) DealApply(id int64, agree bool, userId int64, delayTime int64) error {
	obj, err := this.Get(id)
	if err != nil {
		return err
	}

	obj.Creater = userId
	obj.UpdateTime = time.Now().Unix()

	if delayTime != 0 {
		obj.DelayTime = delayTime
	}

	if !agree {
		obj.Status = DELAYGUARANTEEREFUSE
		err = this.Set(obj)
		if err != nil {
			return err
		}
		return nil
	}

	good, err := GoodsDao.GetByUUID(obj.GoodsUUID)
	if err != nil {
		return err
	}
	good.GuaranteeTime = delayTime
	good.UpdateTime = time.Now().Unix()
	obj.Status = DELAYGUARANTEECOMFIRM

	tx := this.Db.Begin()

	err = tx.Save(good).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Save(obj).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (this *_DelayGuaranteeApplyDao) Get(id int64) (*DelayGuaranteeApply, error) {
	obj := &DelayGuaranteeApply{}
	err := this.Db.Where("datastatus  = ? and id = ?", 0, id).First(obj).Error
	return obj, err
}

func (this *_DelayGuaranteeApplyDao) Set(obj *DelayGuaranteeApply) error {
	err := this.Db.Save(obj).Error
	return err
}
