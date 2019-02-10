package user_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type DeleteUserRequest struct {
	Id int64 `json:"id"`
}

type DeleteUserResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

func DeleteUser(c *gin.Context) {
	req := &DeleteUserRequest{}
	rsp := &DeleteUserResponse{}
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

	u.DataStatus = time.Now().Unix()
	err = model.UserDao.Set(u)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
