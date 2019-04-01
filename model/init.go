package model

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var UserDao *_UserDao
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

func Init() error {
	//db, err := gorm.Open("mysql", "debian-sys-maint:fYzuFNK68VdZTWJ0@/demite?charset=utf8&parseTime=True&loc=Local")
	//db, err := gorm.Open("mysql", "debian-sys-maint:P7Mo08KJ9qIYEZ9b@/demite?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", "root:612345@/demite?charset=utf8&parseTime=True&loc=Local")
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

	//init
	err = UserDao.initUserDao()
	if err != nil {
		return err
	}

	return nil
}
