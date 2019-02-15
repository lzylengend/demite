package class_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListClassRequest struct {
	UpClassId int64 `json:"upclassid"`
}

type ListClassResponse struct {
	Data   []*classData          `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type classData struct {
	ClassId   int64  `json:"classid"`
	ClassName string `json:"classname"`
	Show      bool   `json:"show"`
}

func ListClass(c *gin.Context) {
	req := &ListClassRequest{}
	rsp := &ListClassResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	classList, err := model.ClassDao.ListClassByUp(req.UpClassId)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Data = make([]*classData, 0)
	for _, v := range classList {
		b := true
		if v.Show != 0 {
			b = false
		}

		rsp.Data = append(rsp.Data, &classData{
			ClassId:   v.ClassId,
			ClassName: v.ClassName,
			Show:      b,
		})
	}
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
