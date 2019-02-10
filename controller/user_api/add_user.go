package user_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AddUserRequest struct {
	Name string `json:"name"`
}

type AddUserResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

func AddUser(c *gin.Context) {
	req := &AddUserRequest{}
	rsp := &AddUserResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	u, err := model.UserDao.GetByName(req.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	if u != nil {
		rsp.Status = my_error.UserNameExistError()
		c.JSON(200, rsp)
		return
	}

	err = model.UserDao.AddUser(req.Name, model.DefaltUserPwd)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
