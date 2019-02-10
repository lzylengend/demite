package model

import (
	"demite/util"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type UserPassword struct {
	UserPasswordId int64  `gorm:"column:UserPasswordid;primary_key"`
	Pwd            string `gorm:"column:Pwd"`
	Salt           string `gorm:"column:Salt"`
}

type _UserPasswordDao struct {
	Db *gorm.DB
}

func (UserPassword) TableName() string {
	return "UserPassword"
}

func newUserPasswordDao(db *gorm.DB) *_UserPasswordDao {
	db.AutoMigrate(&UserPassword{})

	return &_UserPasswordDao{Db: db.Model(&UserPassword{})}
}

func (this *_UserPasswordDao) Insert(obj *UserPassword) error {
	return this.Db.Create(obj).Error
}

func (this *_UserPasswordDao) New(userPasswordId int64, userPassword string) *UserPassword {
	salt := uuid.New().String()
	return &UserPassword{
		UserPasswordId: userPasswordId,
		Pwd:            util.Md5(this.PwdCombine(userPassword, salt)),
		Salt:           salt,
	}
}

func (this *_UserPasswordDao) PwdCombine(pwd string, salt string) string {
	return pwd + salt
}

func (this *_UserPasswordDao) GetById(id int64) (*UserPassword, error) {
	obj := &UserPassword{}
	err := this.Db.Where("UserPasswordid = ?", id).First(obj).Error
	return obj, err
}

func (this *_UserPasswordDao) Set(obj *UserPassword) error {
	err := this.Db.Save(obj).Error
	return err
}
