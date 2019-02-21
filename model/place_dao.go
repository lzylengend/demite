package model

import "github.com/jinzhu/gorm"

type Place struct {
	PlaceId    int64  `gorm:"column:placeid;primary_key;"`
	PlaceName  string `gorm:"column:placename;index:placename"`
	UpPlaceId  int64  `gorm:"column:upplaceid;index:upplaceid"` //0为根目录
	IsShow     int64  `gorm:"column:isshow;index:isshow"`
	DataStatus int64  `gorm:"column:datastatus"`
	CreateTime int64  `gorm:"column:createtime"`
	UpdateTime int64  `gorm:"column:updatetime"`
}

type _PlaceDao struct {
	Db *gorm.DB
}

func (Place) TableName() string {
	return "place"
}

func newPlaceDao(db *gorm.DB) *_PlaceDao {
	db.AutoMigrate(&Place{})

	return &_PlaceDao{Db: db.Model(&Place{})}
}

func (this *_PlaceDao) InsertBatch(objList []*Place) error {
	tx := this.Db.Begin()

	for _, v := range objList {
		err := tx.Create(v).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (this *_PlaceDao) ListByUpId(upPlaceId int64) ([]*Place, error) {
	objList := make([]*Place, 0)
	err := this.Db.Where("upplaceid = ? and datastatus = ?", upPlaceId, 0).Find(&objList).Error
	return objList, err
}
