package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

const defaltUserGroupName = "管理员"
const defaltNormalUserGroupName = "普通用户"

type UserGroup struct {
	UserGroupId   int64  `gorm:"column:usergroupid;primary_key;AUTO_INCREMENT"`
	UserGroupName string `gorm:"column:usergroupname;index:groupname"`

	AuthDelGoods     bool `gorm:"column:authdelgoods"`
	AuthShieldWxUser bool `gorm:"column:authshieldwxuser"`
	AuthUserManage   bool `gorm:"column:authusermanage"`
	AuthDelStaff     bool `gorm:"column:authdelstaff"`

	DataStatus int64 `gorm:"column:datastatus"`
	CreateTime int64 `gorm:"column:createtime"`
	UpdateTime int64 `gorm:"column:updatetime"`
}

type _UserGroupDao struct {
	Db *gorm.DB
}

func (UserGroup) TableName() string {
	return "user"
}

func newUserGroupDao(db *gorm.DB) *_UserGroupDao {
	db.AutoMigrate(&UserGroup{})

	return &_UserGroupDao{Db: db.Model(&UserGroup{})}
}

func (this *_UserGroupDao) initUserGroupDao() error {
	c, err := this.Count()
	if err != nil {
		return err
	}

	if c >= 1 {
		return nil
	}

	tx := this.Db.Begin()
	obj := &UserGroup{
		UserGroupName: defaltUserGroupName,

		AuthDelGoods:     true,
		AuthShieldWxUser: true,
		AuthUserManage:   true,
		AuthDelStaff:     true,

		DataStatus: 0,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	err = tx.Create(obj).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	obj2 := &UserGroup{
		UserGroupName: defaltNormalUserGroupName,

		AuthDelGoods:     false,
		AuthShieldWxUser: false,
		AuthUserManage:   false,
		AuthDelStaff:     false,

		DataStatus: 0,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}

	err = tx.Create(obj2).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (this *_UserGroupDao) Count() (int64, error) {
	var n int64 = 0
	err := this.Db.Count(&n).Error
	return n, err
}

func (this *_UserGroupDao) Get(id int64) (*UserGroup, error) {
	obj := &UserGroup{}
	err := this.Db.Where("usergroupid = ?", id).First(obj).Error
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (this *_UserGroupDao) List() ([]*UserGroup, error) {
	objList := make([]*UserGroup, 0)

	err := this.Db.Where("datastatus = ?", 0).Find(&objList).Error
	return objList, err
}
