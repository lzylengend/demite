package model

import "github.com/jinzhu/gorm"

type UserDao struct {
	UserId   int64  `gorm:"column:userid;primary_key"`
	UserName string `gorm:"column:username"`
}

func (UserDao) TableName() string {
	return "user"
}

func newUserDao(db *gorm.DB) {
	db.AutoMigrate(&UserDao{})
}
