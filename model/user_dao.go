package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

const defaltUserName = "admin"
const DefaltUserPwd = "123456"

type User struct {
	UserId     int64  `gorm:"column:userid;primary_key;AUTO_INCREMENT"`
	UserName   string `gorm:"column:username;index:name"`
	DataStatus int64  `gorm:"column:datastatus"`
	CreateTime int64  `gorm:"column:createtime"`
	UpdateTime int64  `gorm:"column:updatetime"`
}

type _UserDao struct {
	Db *gorm.DB
}

func (User) TableName() string {
	return "user"
}

func newUserDao(db *gorm.DB) *_UserDao {
	db.AutoMigrate(&User{})

	return &_UserDao{Db: db.Model(&User{})}
}

func (this *_UserDao) initUserDao() error {
	err := this.Db.AddUniqueIndex("idx_user_name_datastatus", "username", "datastatus").Error
	if err != nil {
		return err
	}

	c, err := this.Count()
	if err != nil {
		return err
	}

	if c >= 1 {
		return nil
	}

	err = this.AddUser(defaltUserName, DefaltUserPwd)
	if err != nil {
		return err
	}
	return nil
}

func (this *_UserDao) AddUser(username string, pwd string) error {
	objUser := this.NewUser(username)

	tx := this.Db.Begin()

	err := tx.Create(objUser).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	objPwd := UserPasswordDao.New(objUser.UserId, pwd)

	err = tx.Create(objPwd).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (this *_UserDao) NewUser(username string) *User {
	return &User{
		UserName:   username,
		DataStatus: 0,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
}

func (this *_UserDao) Insert(obj *User) error {
	return this.Db.Create(obj).Error
}

func (this *_UserDao) Count() (int, error) {
	n := 0
	err := this.Db.Count(&n).Error
	return n, err
}

func (this *_UserDao) GetByName(name string) (*User, error) {
	obj := &User{}

	err := this.Db.Where("username = ?", name).First(obj).Error
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (this *_UserDao) GetById(id int64) (*User, error) {
	obj := &User{}
	err := this.Db.Where("userid = ?", id).First(obj).Error
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (this *_UserDao) ListByKey(limit, offset int64, key string) ([]*User, error) {
	objList := make([]*User, 0)

	key = "%" + key + "%"
	err := this.Db.Where("username like ? and datastatus = ?", key, 0).Offset(offset).Limit(limit).Order("createtime").Find(&objList).Error
	return objList, err
}

func (this *_UserDao) CountByKey(key string) (int64, error) {
	n := 0

	key = "%" + key + "%"
	err := this.Db.Where("username like ? and datastatus = ?", key, 0).Count(&n).Error
	return int64(n), err
}

func (this *_UserDao) Set(obj *User) error {
	return this.Db.Save(obj).Error
}
