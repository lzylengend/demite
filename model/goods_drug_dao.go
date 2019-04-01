package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type GoodDrugs struct {
	Id         int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	GoodUUId   string `gorm:"column:goodsuuid;index:goodsuuid"`
	DrugId     int64  `gorm:"column:drugid;index:drugid"`
	CreateTime int64  `gorm:"column:createtime"`
	UpdateTime int64  `gorm:"column:updatetime"`
	DataStatus int64  `gorm:"column:datastatus"`
}
type _GoodDrugsDao struct {
	Db *gorm.DB
}

func (GoodDrugs) TableName() string {
	return "gooddrugs"
}

func newGoodDrugsDao(db *gorm.DB) *_GoodDrugsDao {
	db.AutoMigrate(&GoodDrugs{})

	return &_GoodDrugsDao{Db: db.Model(&GoodDrugs{})}
}

func (this *_GoodDrugsDao) Add(id []int64, uuid string) error {
	t := this.Db.Begin()
	for i, _ := range id {
		err := t.Create(&GoodDrugs{
			GoodUUId:   uuid,
			DrugId:     id[i],
			CreateTime: time.Now().Unix(),
			UpdateTime: time.Now().Unix(),
			DataStatus: 0,
		}).Error
		if err != nil {
			t.Rollback()
			return err
		}
	}

	t.Commit()

	return nil
}

func (this *_GoodDrugsDao) Update(id []int64, uuid string) error {
	objList, err := this.GetByUUID(uuid)
	if err != nil {
		return err
	}

	delList := make([]*GoodDrugs, 0)
	addList := make([]*GoodDrugs, 0)
	for i, v := range objList {
		existFlag := false
		for _, v2 := range id {
			if v.DrugId == v2 {
				existFlag = true
				break
			}
		}

		if !existFlag {
			objList[i].DataStatus = 1
			objList[i].UpdateTime = time.Now().Unix()
			delList = append(delList, objList[i])
		}
	}

	fmt.Println(delList)

	for _, v := range id {
		existFlag := false
		for _, v2 := range objList {
			if v2.DrugId == v {
				existFlag = true
				break
			}
		}

		if !existFlag {
			addList = append(addList, &GoodDrugs{
				GoodUUId:   uuid,
				DrugId:     v,
				CreateTime: time.Now().Unix(),
				UpdateTime: time.Now().Unix(),
				DataStatus: 0,
			})
		}
	}
	fmt.Println(addList)

	t := this.Db.Begin()

	for i, _ := range delList {
		err = t.Save(delList[i]).Error
		if err != nil {
			t.Rollback()
			return err
		}
	}

	for i, _ := range addList {
		err = t.Create(addList[i]).Error
		if err != nil {
			t.Rollback()
			return err
		}
	}

	t.Commit()

	return nil
}

func (this *_GoodDrugsDao) GetByUUID(uuid string) ([]*GoodDrugs, error) {
	objList := make([]*GoodDrugs, 0)
	err := this.Db.Where("goodsuuid = ? and datastatus = ?", uuid, 0).Find(&objList).Error
	return objList, err
}
