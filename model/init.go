package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

func Init() error {
	db, err := gorm.Open("mysql", "lzy:612345@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}

	db.LogMode(true)
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))

}
