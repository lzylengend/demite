package main

import (
	"demite/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

func main() {
	b, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	m := make(map[string]map[string]string)
	err = json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = model.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	objList := make([]*model.Place, 0)
	for k, v := range m {
		up := 0
		if k != "86" {
			up, err = strconv.Atoi(k)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		for k2, v2 := range v {
			id, err := strconv.Atoi(k2)
			if err != nil {
				fmt.Println(err)
				return
			}

			objList = append(objList, &model.Place{
				PlaceId:    int64(id),
				UpPlaceId:  int64(up),
				PlaceName:  v2,
				IsShow:     0,
				DataStatus: 0,
				CreateTime: time.Now().Unix(),
				UpdateTime: time.Now().Unix(),
			})
		}
	}

	err = model.PlaceDao.InsertBatch(objList)
	if err != nil {
		fmt.Println(err)
		return
	}
}
