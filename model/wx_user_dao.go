package model

type WxUser struct {
	DataStatus int64 `gorm:"column:datastatus"`
	CreateTime int64 `gorm:"column:createtime"`
	UpdateTime int64 `gorm:"column:updatetime"`
}
