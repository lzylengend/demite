package model

const defaltUserGroupName = "admin"

type UserGroup struct {
	UserGroupId   int64  `gorm:"column:usergroupid;primary_key;AUTO_INCREMENT"`
	UserGroupName string `gorm:"column:usergroupname;index:groupname"`

	DataStatus int64 `gorm:"column:datastatus"`
	CreateTime int64 `gorm:"column:createtime"`
	UpdateTime int64 `gorm:"column:updatetime"`
}
