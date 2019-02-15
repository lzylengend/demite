package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var UserDao *_UserDao
var UserPasswordDao *_UserPasswordDao
var WxUserDao *_WxUserDao
var ClassDao *_ClassDao

func Init() error {
	//db, err := gorm.Open("mysql", "debian-sys-maint:fYzuFNK68VdZTWJ0@/demite?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", "debian-sys-maint:P7Mo08KJ9qIYEZ9b@/demite?charset=utf8&parseTime=True&loc=Local")
	//db, err := gorm.Open("mysql", "root:612345@/demite?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}

	db.LogMode(true)
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))

	UserDao = newUserDao(db)
	UserPasswordDao = newUserPasswordDao(db)
	WxUserDao = newWxUserDao(db)
	ClassDao = newClassDao(db)

	//init
	err = UserDao.initUserDao()
	if err != nil {
		return err
	}

	return nil
}
