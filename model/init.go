package model

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var UserDao *_UserDao
var UserGroupDao *_UserGroupDao
var UserPasswordDao *_UserPasswordDao
var WxUserDao *_WxUserDao
var ClassDao *_ClassDao
var DrugClassDao *_DrugClassDao
var PlaceDao *_PlaceDao
var ProduceDao *_ProductDao
var OrdertDao *_OrdertDao
var OrderLogDao *_OrdertLogDao
var GoodsDao *_GoodsDao
var DrugDao *_DrugDao
var GoodDrugsDao *_GoodDrugsDao
var GoodsWXUserDao *_GoodsWXUserDao
var UnlockApplyDao *_UnlockApplyDao
var DelayGuaranteeApplyDao *_DelayGuaranteeApplyDao
var StaffDao *_StaffDao
var RepairScheduleDao *_RepairScheduleDao
var RepairDao *_RepairDao
var RemoteDao *_RemoteDao
var RemoteScheduleDao *_RemoteScheduleDao
var VideoClassDao *_VideoClassDao
var VideoDao *_VideoDao
var MaterialDao *_MaterialDao
var MaterialClassDao *_MaterialClassDao
var SchemeDao *_SchemeDao
var SoftDao *_SoftDao
var SoftClassDao *_SoftClassDao
var QADao *_QADao

func Init(dbPath string) error {
	//db, err := gorm.Open("mysql", "debian-sys-maint:fYzuFNK68VdZTWJ0@/demite?charset=utf8&parseTime=True&loc=Local")
	//db, err := gorm.Open("mysql", "debian-sys-maint:P7Mo08KJ9qIYEZ9b@/demite?charset=utf8&parseTime=True&loc=Local")
	//db, err := gorm.Open("mysql", "root:612345@/demite?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", dbPath)
	if err != nil {
		return err
	}

	db.LogMode(true)
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))

	UserDao = newUserDao(db)
	UserPasswordDao = newUserPasswordDao(db)
	WxUserDao = newWxUserDao(db)
	ClassDao = newClassDao(db)
	PlaceDao = newPlaceDao(db)
	ProduceDao = newProductDao(db)
	OrdertDao = newOrderDao(db)
	OrderLogDao = newOrderLogDao(db)
	GoodsDao = newGoodsDao(db)
	DrugClassDao = newDrugClassDao(db)
	DrugDao = newDrugDao(db)
	GoodDrugsDao = newGoodDrugsDao(db)
	GoodsWXUserDao = newGoodsWXUserDao(db)
	UnlockApplyDao = newUnlockApplyDao(db)
	DelayGuaranteeApplyDao = newDelayGuaranteeApplyDao(db)
	StaffDao = newStaffDao(db)
	RepairDao = newRepairDao(db)
	RepairScheduleDao = newRepairScheduleDao(db)
	RemoteDao = newRemoteDao(db)
	RemoteScheduleDao = newRemoteScheduleDao(db)
	UserGroupDao = newUserGroupDao(db)
	VideoClassDao = newVideoClassDao(db)
	VideoDao = newVideoDao(db)
	MaterialDao = newMaterialDao(db)
	MaterialClassDao = newMaterialClassDao(db)
	SchemeDao = newSchemeDao(db)
	SoftDao = newSoftDao(db)
	SoftClassDao = newSoftClassDao(db)
	QADao = newQADao(db)

	//init
	err = UserDao.initUserDao()
	if err != nil {
		return err
	}

	err = UserGroupDao.initUserGroupDao()
	if err != nil {
		return err
	}

	err = VideoClassDao.Init()
	if err != nil {
		return err
	}

	err = SchemeDao.init()
	if err != nil {
		return err
	}

	return nil
}
