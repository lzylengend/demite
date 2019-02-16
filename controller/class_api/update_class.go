package class_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type UpdateClassRequest struct {
	ClassId int64  `json:"classid"`
	Name    string `json:"name"`
	Show    bool   `json:"show"`
}

type UpdateClassResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

func UpdateClass(c *gin.Context) {
	req := &UpdateClassRequest{}
	rsp := &UpdateClassResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	if req.Name == "" {
		rsp.Status = my_error.NotNilError("分组名")
		c.JSON(200, rsp)
		return
	}

	class, err := model.ClassDao.GetClassById(req.ClassId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	class.ClassName = req.Name
	if req.Show {
		class.IsShow = 0
	} else {
		class.IsShow = time.Now().Unix()
	}

	err = model.ClassDao.Set(class)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
