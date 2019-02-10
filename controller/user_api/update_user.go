package user_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

type UpdateUserRequest struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type UpdateUserResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

func UpdateUser(c *gin.Context) {
	req := &UpdateUserRequest{}
	rsp := &UpdateUserResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	u, err := model.UserDao.GetById(req.Id)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	if u.UserName != req.Name {
		u2, err := model.UserDao.GetByName(req.Name)
		if err != nil && err != gorm.ErrRecordNotFound {
			rsp.Status = my_error.DbError(err.Error())
			c.JSON(200, rsp)
			return
		}

		if u2 != nil {
			rsp.Status = my_error.UserNameExistError()
			c.JSON(200, rsp)
			return
		}

		u.UserName = req.Name
		u.UpdateTime = time.Now().Unix()
		err = model.UserDao.Set(u)
		if err != nil {
			rsp.Status = my_error.DbError(err.Error())
			c.JSON(200, rsp)
			return
		}
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
